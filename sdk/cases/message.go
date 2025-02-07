package main

import (
	"context"

	"github.com/ipfs-force-community/go-fvm-sdk/sdk"
	"github.com/ipfs-force-community/go-fvm-sdk/sdk/testing"
	"github.com/stretchr/testify/assert"
)

func main() {} //nolint

//go:export invoke
func Invoke(_ uint32) uint32 { //nolint
	t := testing.NewTestingT()
	defer t.CheckResult()
	ctx := context.Background()
	_, err := sdk.Caller(ctx)
	assert.Nil(t, err)
	//	assert.Equal(t, caller, 1) todo unable to verify caller, its random value in tester

	receiver, err := sdk.Receiver(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 10000, int(receiver))

	methodNum, err := sdk.MethodNumber(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 1, int(methodNum))

	valueRecieved, err := sdk.ValueReceived(ctx)
	assert.Nil(t, err)
	assert.Equal(t, "10000000000000000000", valueRecieved.String())
	return 0
}
