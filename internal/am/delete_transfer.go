package am

import (
	"context"
	"fmt"

	temporal_tools "go.artefactual.dev/tools/temporal"
)

const DeleteTransferActivityName = "DeleteTransferActivity"

type DeleteTransferActivityParams struct {
	Destination string
}

type DeleteTransferActivity struct {
	client TransferClient
}

type DeleteTransferActivityResult struct{}

func NewDeleteTransferActivity(client TransferClient) *DeleteTransferActivity {
	return &DeleteTransferActivity{client: client}
}

func (a *DeleteTransferActivity) Execute(
	ctx context.Context,
	params *DeleteTransferActivityParams,
) (*DeleteTransferActivityResult, error) {
	logger := temporal_tools.GetLogger(ctx)
	logger.V(1).Info("Execute DeleteTransferActivity",
		"destination", params.Destination,
	)

	err := a.client.Delete(ctx, params.Destination)
	if err != nil {
		return &DeleteTransferActivityResult{}, fmt.Errorf("delete transfer: path: %q: %v", params.Destination, err)
	}

	return &DeleteTransferActivityResult{}, nil
}
