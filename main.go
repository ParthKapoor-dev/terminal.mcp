package main

import (
	"context"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/parthkapoor-dev/terminal-mcp/tools"
)

func main() {
	// Create a server with a single tool.
	server := mcp.NewServer(&mcp.Implementation{Name: "Terminal MCP Server", Version: "v1.0.0"}, nil)

	mcp.AddTool(server, &mcp.Tool{Name: "greet", Description: "This Tools is used to say hi"}, tools.SayHi)
	mcp.AddTool(server, &mcp.Tool{Name: "ls", Description: "List Files and Directories at given path"}, tools.Ls)

	// Run the server over stdin/stdout, until the client disconnects
	if err := server.Run(context.Background(), mcp.NewStdioTransport()); err != nil {
		log.Fatal(err)
	}
}
