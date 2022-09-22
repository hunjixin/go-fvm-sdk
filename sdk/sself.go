package sdk

import (
	addr "github.com/filecoin-project/go-address"
	"github.com/ipfs-force-community/go-fvm-sdk/sdk/sys"
	"github.com/ipfs-force-community/go-fvm-sdk/sdk/types"
	"github.com/ipfs/go-cid"
)

// Root Get the IPLD root CID. Fails if the actor doesn't have state (before the first call to
// `set_root` and after actor deletion).
func Root() (cid.Cid, error) {
	return sys.SelfRoot()
}

// SetRoot set the actor's state-tree root.
//
// Fails if:
//
// - The new root is not in the actor's "reachable" set.
// - Fails if the actor has been deleted.
func SetRoot(c cid.Cid) error {
	return sys.SelfSetRoot(c)
}

// CurrentBalance gets the current balance for the calling actor.
func CurrentBalance() *types.TokenAmount {
	tok, err := sys.SelfCurrentBalance()
	if err != nil {
		panic(err.Error())
	}
	return tok
}

// SelfDestruct destroys the calling actor, sending its current balance
// to the supplied address, which cannot be itself.
//
// Fails if the beneficiary doesn't exist or is the actor being deleted.
func SelfDestruct(addr addr.Address) error {
	return sys.SelfDestruct(addr)
}
