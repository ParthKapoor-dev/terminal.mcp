package tools

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type ListParams struct {
}

func Ls(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParamsFor[ListParams]) (*mcp.CallToolResultFor[any], error) {

	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: "Hello again",
			},
		},
	}, nil

}
