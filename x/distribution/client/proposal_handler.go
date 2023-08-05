package client

import (
	"github.com/blockgenx/blockgen-sdk/x/distribution/client/cli"
	govclient "github.com/blockgenx/blockgen-sdk/x/gov/client"
)

// ProposalHandler is the community spend proposal handler.
var (
	ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal)
)
