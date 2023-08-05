package client

import (
	govclient "github.com/blockgenx/blockgen-sdk/x/gov/client"
	"github.com/blockgenx/blockgen-sdk/x/params/client/cli"
)

// ProposalHandler is the param change proposal handler.
var ProposalHandler = govclient.NewProposalHandler(cli.NewSubmitParamChangeProposalTxCmd)
