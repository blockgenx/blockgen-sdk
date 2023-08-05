package main

import (
	"context"
	"testing"

	"github.com/blockgenx/blockgen-sdk/cosmovisor"
	"github.com/blockgenx/blockgen-sdk/testutil"
	"github.com/stretchr/testify/require"
)

func TestVersionCommand_Error(t *testing.T) {
	logger := cosmovisor.NewLogger()

	rootCmd.SetArgs([]string{"version"})
	_, out := testutil.ApplyMockIO(rootCmd)
	ctx := context.WithValue(context.Background(), cosmovisor.LoggerKey, logger)

	require.Error(t, rootCmd.ExecuteContext(ctx))
	require.Contains(t, out.String(), "DAEMON_NAME is not set")
}
