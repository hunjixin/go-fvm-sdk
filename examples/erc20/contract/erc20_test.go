package contract

import (
	"testing"

	"github.com/filecoin-project/go-address"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin/v9/migration"

	"github.com/filecoin-project/go-state-types/big"
	"github.com/ipfs-force-community/go-fvm-sdk/sdk"
	"github.com/ipfs-force-community/go-fvm-sdk/sdk/adt"
	"github.com/ipfs-force-community/go-fvm-sdk/sdk/sys/simulated"
	"github.com/ipfs-force-community/go-fvm-sdk/sdk/types"
	"github.com/stretchr/testify/assert"
)

func TestErc20TokenFakeSetBalance(t *testing.T) {
	simulator, ctx := simulated.CreateSimulateEnv(&types.InvocationContext{}, big.NewInt(1), big.NewInt(1))
	addr, err := simulated.NewF1Address()
	assert.NoError(t, err)
	simulator.SetActor(abi.ActorID(1), addr, migration.Actor{})

	empMap, err := adt.MakeEmptyMap(adt.AdtStore(ctx), adt.BalanceTableBitwidth)
	assert.Nil(t, err)
	emptyBalance, err := empMap.Root()
	assert.Nil(t, err)

	erc20State := &Erc20Token{Name: "name", Symbol: "symbol", Decimals: 8, TotalSupply: abi.NewTokenAmount(100000), Balances: emptyBalance, Allowed: emptyBalance}
	err = erc20State.FakeSetBalance(ctx, &FakeSetBalance{Addr: addr, Balance: abi.NewTokenAmount(1)})
	assert.NoError(t, err)
}

func TestErc20TokenGetter(t *testing.T) {
	simulator, ctx := simulated.CreateSimulateEnv(&types.InvocationContext{}, abi.NewTokenAmount(1), abi.NewTokenAmount(1))
	addr, err := simulated.NewF1Address()
	assert.NoError(t, err)
	simulator.SetActor(abi.ActorID(1), addr, migration.Actor{})

	empMap, err := adt.MakeEmptyMap(adt.AdtStore(ctx), adt.BalanceTableBitwidth)
	assert.Nil(t, err)
	emptyBalance, err := empMap.Root()
	assert.Nil(t, err)

	erc20State := &Erc20Token{Name: "EP Coin", Symbol: "EP", Decimals: 8, TotalSupply: abi.NewTokenAmount(100000), Balances: emptyBalance, Allowed: emptyBalance}

	t.Run("get name", func(t *testing.T) {
		assert.Equal(t, erc20State.Name, "EP Coin")
	})

	t.Run("get symbol", func(t *testing.T) {
		assert.Equal(t, erc20State.Symbol, "EP")
	})

	t.Run("get decimals", func(t *testing.T) {
		assert.Equal(t, erc20State.Decimals, uint8(8))
	})

	t.Run("get supply", func(t *testing.T) {
		assert.Equal(t, erc20State.TotalSupply.Uint64(), uint64(100000))
	})
}

func TestErc20TokenGetBalanceOf(t *testing.T) {
	simulator, ctx := simulated.CreateSimulateEnv(&types.InvocationContext{}, abi.NewTokenAmount(1), abi.NewTokenAmount(1))
	actor := abi.ActorID(1)
	addr, err := simulated.NewF1Address()
	assert.NoError(t, err)
	simulator.SetActor(actor, addr, migration.Actor{})

	balanceMap, err := adt.MakeEmptyMap(adt.AdtStore(ctx), adt.BalanceTableBitwidth)
	assert.Nil(t, err)
	emptyRoot, err := balanceMap.Root()

	assert.Nil(t, balanceMap.Put(types.ActorKey(actor), simulated.NewPtrTokenAmount(100)))
	balanceRoot, err := balanceMap.Root()
	assert.Nil(t, err)

	erc20State := &Erc20Token{Name: "Ep Coin", Symbol: "EP", Decimals: 8, TotalSupply: abi.NewTokenAmount(100000), Balances: balanceRoot, Allowed: emptyRoot}
	sdk.SaveState(ctx, erc20State) //Save state

	got, err := erc20State.GetBalanceOf(ctx, &addr)
	assert.Nil(t, err)
	assert.Equal(t, got.Uint64(), uint64(100))
}

func TestErc20TokenTransfer(t *testing.T) {
	setup := func(t *testing.T, fromInitBalance abi.TokenAmount) (*simulated.FvmSimulator, address.Address, address.Address) {
		simulator, ctx := simulated.CreateSimulateEnv(&types.InvocationContext{}, abi.NewTokenAmount(1), abi.NewTokenAmount(1))
		fromActor := abi.ActorID(1)
		fromAddr, err := simulated.NewF1Address()
		assert.NoError(t, err)
		simulator.SetActor(fromActor, fromAddr, migration.Actor{})

		toActor := abi.ActorID(2)
		toAddr, err := simulated.NewF1Address()
		assert.NoError(t, err)
		simulator.SetActor(toActor, toAddr, migration.Actor{})

		balanceMap, err := adt.MakeEmptyMap(adt.AdtStore(ctx), adt.BalanceTableBitwidth)
		assert.NoError(t, err)
		emptyRoot, err := balanceMap.Root()
		assert.NoError(t, err)
		assert.NoError(t, balanceMap.Put(types.ActorKey(fromActor), &fromInitBalance))
		balanceRoot, err := balanceMap.Root()
		assert.Nil(t, err)

		erc20State := &Erc20Token{Name: "Ep Coin", Symbol: "EP", Decimals: 8, TotalSupply: abi.NewTokenAmount(100000), Balances: balanceRoot, Allowed: emptyRoot}
		_ = sdk.SaveState(ctx, erc20State) //Save state

		// set info of context
		simulator.SetCallContext(&types.InvocationContext{
			Caller: fromActor,
		})
		return simulator, fromAddr, toAddr
	}

	t.Run("successful", func(t *testing.T) {
		simulator, fromAddr, toAddr := setup(t, abi.NewTokenAmount(1000))

		var newState Erc20Token
		sdk.LoadState(simulator.Context, &newState)

		assert.NoError(t, newState.Transfer(simulator.Context, &TransferReq{
			ReceiverAddr:   toAddr,
			TransferAmount: abi.NewTokenAmount(100),
		}))

		fromBalance, err := newState.GetBalanceOf(simulator.Context, &fromAddr)
		assert.NoError(t, err)
		assert.Equal(t, uint64(900), fromBalance.Uint64())

		toBalance, err := newState.GetBalanceOf(simulator.Context, &toAddr)
		assert.NoError(t, err)
		assert.Equal(t, uint64(100), toBalance.Uint64())
	})

	t.Run("fail transfer zero", func(t *testing.T) {
		simulator, _, toAddr := setup(t, abi.NewTokenAmount(1000))

		var newState Erc20Token
		sdk.LoadState(simulator.Context, &newState)

		assert.EqualError(t, newState.Transfer(simulator.Context, &TransferReq{
			ReceiverAddr:   toAddr,
			TransferAmount: abi.NewTokenAmount(0),
		}), "transfer value must bigger than zero")
	})

	t.Run("fail balance not enough", func(t *testing.T) {
		simulator, _, toAddr := setup(t, abi.NewTokenAmount(1000))

		var newState Erc20Token
		sdk.LoadState(simulator.Context, &newState)

		assert.EqualError(t, newState.Transfer(simulator.Context, &TransferReq{
			ReceiverAddr:   toAddr,
			TransferAmount: abi.NewTokenAmount(10000),
		}), "transfer amount should be less than balance of sender (1): 10000 should be <= to 1000")
	})
}

func TestApprovalAndTransfer(t *testing.T) {
	setup := func(t *testing.T, fromInitBalance abi.TokenAmount) (*simulated.FvmSimulator, address.Address, address.Address, address.Address) {
		simulator, ctx := simulated.CreateSimulateEnv(&types.InvocationContext{}, abi.NewTokenAmount(1), abi.NewTokenAmount(1))
		fromActor := abi.ActorID(1)
		fromAddr, err := simulated.NewF1Address()
		assert.NoError(t, err)
		simulator.SetActor(fromActor, fromAddr, migration.Actor{})

		approvalActor := abi.ActorID(2)
		approvalAddr, err := simulated.NewF1Address()
		assert.NoError(t, err)
		simulator.SetActor(approvalActor, approvalAddr, migration.Actor{})

		toActor := abi.ActorID(3)
		toAddr, err := simulated.NewF1Address()
		assert.NoError(t, err)
		simulator.SetActor(toActor, toAddr, migration.Actor{})

		balanceMap, err := adt.MakeEmptyMap(adt.AdtStore(ctx), adt.BalanceTableBitwidth)
		assert.NoError(t, err)
		emptyRoot, err := balanceMap.Root()
		assert.NoError(t, err)
		assert.NoError(t, balanceMap.Put(types.ActorKey(fromActor), &fromInitBalance))
		balanceRoot, err := balanceMap.Root()
		assert.Nil(t, err)

		erc20State := &Erc20Token{Name: "Ep Coin", Symbol: "EP", Decimals: 8, TotalSupply: abi.NewTokenAmount(100000), Balances: balanceRoot, Allowed: emptyRoot}
		_ = sdk.SaveState(ctx, erc20State) //Save state

		// set info of context
		simulator.SetCallContext(&types.InvocationContext{
			Caller: fromActor,
		})
		return simulator, fromAddr, approvalAddr, toAddr
	}

	t.Run("success approval and transfer", func(t *testing.T) {

	})

	t.Run("fail approval zero balance", func(t *testing.T) {
		simulator, fromAddr, approvalAddr, _ := setup(t, abi.NewTokenAmount(1000))
		fromId, err := simulator.ResolveAddress(fromAddr)
		assert.NoError(t, err)
		var newState Erc20Token
		sdk.LoadState(simulator.Context, &newState)
		ctx := simulator.Context
		simulator.SetCallContext(&types.InvocationContext{
			ValueReceived:    abi.NewTokenAmount(0),
			Caller:           fromId,
			Receiver:         0,
			MethodNumber:     0,
			NetworkCurrEpoch: 0,
			NetworkVersion:   0,
		})
		assert.EqualError(t, newState.Approval(ctx, &ApprovalReq{
			SpenderAddr:  approvalAddr,
			NewAllowance: abi.NewTokenAmount(0),
		}), "allow value must bigger than zero")
	})

	t.Run("fail transferfrom zero balance ", func(t *testing.T) {
		simulator, fromAddr, approvalAddr, toAddr := setup(t, abi.NewTokenAmount(1000))
		fromId, err := simulator.ResolveAddress(fromAddr)
		assert.NoError(t, err)
		approvalId, err := simulator.ResolveAddress(approvalAddr)
		assert.NoError(t, err)
		var newState Erc20Token
		sdk.LoadState(simulator.Context, &newState)
		ctx := simulator.Context
		simulator.SetCallContext(&types.InvocationContext{
			ValueReceived:    abi.NewTokenAmount(0),
			Caller:           fromId,
			Receiver:         0,
			MethodNumber:     0,
			NetworkCurrEpoch: 0,
			NetworkVersion:   0,
		})
		assert.NoError(t, newState.Approval(ctx, &ApprovalReq{
			SpenderAddr:  approvalAddr,
			NewAllowance: abi.NewTokenAmount(100),
		}))

		simulator.SetCallContext(&types.InvocationContext{
			ValueReceived:    abi.NewTokenAmount(0),
			Caller:           approvalId,
			Receiver:         0,
			MethodNumber:     0,
			NetworkCurrEpoch: 0,
			NetworkVersion:   0,
		})
		assert.EqualError(t, newState.TransferFrom(ctx, &TransferFromReq{
			OwnerAddr:      fromAddr,
			ReceiverAddr:   toAddr,
			TransferAmount: abi.NewTokenAmount(0),
		}), "send value must bigger than zero")
	})

	t.Run("fail transferfrom zero balance ", func(t *testing.T) {
		simulator, fromAddr, approvalAddr, toAddr := setup(t, abi.NewTokenAmount(1000))
		approvalId, err := simulator.ResolveAddress(approvalAddr)
		assert.NoError(t, err)
		var newState Erc20Token
		sdk.LoadState(simulator.Context, &newState)
		ctx := simulator.Context

		simulator.SetCallContext(&types.InvocationContext{
			ValueReceived:    abi.NewTokenAmount(0),
			Caller:           approvalId,
			Receiver:         0,
			MethodNumber:     0,
			NetworkCurrEpoch: 0,
			NetworkVersion:   0,
		})
		assert.EqualError(t, newState.TransferFrom(ctx, &TransferFromReq{
			OwnerAddr:      fromAddr,
			ReceiverAddr:   toAddr,
			TransferAmount: abi.NewTokenAmount(1),
		}), "approved amount for 1-2 less than zero")
	})

	t.Run("fail transferfrom zero balance ", func(t *testing.T) {
		simulator, fromAddr, approvalAddr, toAddr := setup(t, abi.NewTokenAmount(1000))
		fromId, err := simulator.ResolveAddress(fromAddr)
		assert.NoError(t, err)
		approvalId, err := simulator.ResolveAddress(approvalAddr)
		assert.NoError(t, err)
		var newState Erc20Token
		sdk.LoadState(simulator.Context, &newState)
		ctx := simulator.Context
		simulator.SetCallContext(&types.InvocationContext{
			ValueReceived:    abi.NewTokenAmount(0),
			Caller:           fromId,
			Receiver:         0,
			MethodNumber:     0,
			NetworkCurrEpoch: 0,
			NetworkVersion:   0,
		})
		assert.NoError(t, newState.Approval(ctx, &ApprovalReq{
			SpenderAddr:  approvalAddr,
			NewAllowance: abi.NewTokenAmount(100),
		}))

		simulator.SetCallContext(&types.InvocationContext{
			ValueReceived:    abi.NewTokenAmount(0),
			Caller:           approvalId,
			Receiver:         0,
			MethodNumber:     0,
			NetworkCurrEpoch: 0,
			NetworkVersion:   0,
		})
		assert.EqualError(t, newState.TransferFrom(ctx, &TransferFromReq{
			OwnerAddr:      fromAddr,
			ReceiverAddr:   toAddr,
			TransferAmount: abi.NewTokenAmount(200),
		}), "transfer amount should be less than approved spending amount of 2: 200 should be <= to 100")
	})

	t.Run("fail transferfrom zero balance ", func(t *testing.T) {
		simulator, fromAddr, approvalAddr, toAddr := setup(t, abi.NewTokenAmount(60))
		fromId, err := simulator.ResolveAddress(fromAddr)
		assert.NoError(t, err)
		approvalId, err := simulator.ResolveAddress(approvalAddr)
		assert.NoError(t, err)
		var newState Erc20Token
		sdk.LoadState(simulator.Context, &newState)
		ctx := simulator.Context
		simulator.SetCallContext(&types.InvocationContext{
			ValueReceived:    abi.NewTokenAmount(0),
			Caller:           fromId,
			Receiver:         0,
			MethodNumber:     0,
			NetworkCurrEpoch: 0,
			NetworkVersion:   0,
		})
		assert.NoError(t, newState.Approval(ctx, &ApprovalReq{
			SpenderAddr:  approvalAddr,
			NewAllowance: abi.NewTokenAmount(100),
		}))

		simulator.SetCallContext(&types.InvocationContext{
			ValueReceived:    abi.NewTokenAmount(0),
			Caller:           approvalId,
			Receiver:         0,
			MethodNumber:     0,
			NetworkCurrEpoch: 0,
			NetworkVersion:   0,
		})
		assert.EqualError(t, newState.TransferFrom(ctx, &TransferFromReq{
			OwnerAddr:      fromAddr,
			ReceiverAddr:   toAddr,
			TransferAmount: abi.NewTokenAmount(80),
		}), "transfer amount should be less than balance of token owner (1): 80 should be <= to 60")
	})
}
