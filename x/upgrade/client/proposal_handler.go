package client

import (
	govclient "github.com/blockgenx/blockgen-sdk/x/gov/client"
	"github.com/blockgenx/blockgen-sdk/x/upgrade/client/cli"
)

var (
	LegacyProposalHandler       = govclient.NewProposalHandler(cli.NewCmdSubmitLegacyUpgradeProposal)
	LegacyCancelProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitLegacyCancelUpgradeProposal)
)
