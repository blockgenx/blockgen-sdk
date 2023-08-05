package testutil

import (
	"github.com/blockgenx/blockgen-sdk/testutil"
	clitestutil "github.com/blockgenx/blockgen-sdk/testutil/cli"
	"github.com/blockgenx/blockgen-sdk/testutil/network"
	"github.com/blockgenx/blockgen-sdk/x/authz/client/cli"
)

func CreateGrant(val *network.Validator, args []string) (testutil.BufferWriter, error) {
	cmd := cli.NewCmdGrantAuthorization()
	clientCtx := val.ClientCtx
	return clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
}
