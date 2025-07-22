package tools

import (
	"context"
	"log"
	"os/exec"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type ListParams struct {
	Path string `json:"path" jsonschema:"path to list"`
}

func Ls(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParamsFor[ListParams]) (*mcp.CallToolResultFor[any], error) {

	cmd := exec.Command("ls", params.Arguments.Path)

	output, err := cmd.Output()
	if err != nil {
		log.Println("Err: ", error.Error(err))
		return nil, err
	}

	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: string(output),
			},
		},
	}, nil

}
