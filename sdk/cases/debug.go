package main

import (
	"github.com/ipfs-force-community/go-fvm-sdk/sdk"
	"github.com/ipfs-force-community/go-fvm-sdk/sdk/testing"
	"github.com/stretchr/testify/assert"
)

func main() {}

//go:export invoke
func Invoke(_ uint32) uint32 {
	t := testing.NewTestingT()
	defer t.CheckResult()

	logger, err := sdk.NewLogger()
	assert.Nil(t, err, "create debug logger %v", err)

	enabled := logger.Enabled()
	assert.Equal(t, false, enabled)

	err = logger.Log("")
	assert.Nil(t, err)

	return 0
}
