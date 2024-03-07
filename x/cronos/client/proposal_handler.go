package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/swag-eag/swa/v2/x/swa/client/cli"
)

// ProposalHandler is the token mapping change proposal handler.
var ProposalHandler = govclient.NewProposalHandler(cli.NewSubmitTokenMappingChangeProposalTxCmd)
