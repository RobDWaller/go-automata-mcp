package main

import (
	"context"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/robdwaller/go-automata/automata"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type Input struct {
	Steps uint `json:"steps" jsonschema:"the number of steps to run"`
	Rule uint8 `json:"rule" jsonschema:"the automata rule to use"`
	Seed string `json:"seed" jsonschema:"the initial automata seed"`
}

type Output struct {
	Automata string `json:"automata" jsonschema:"the cellular automata to return to the user"`
}

func GenerateAutomata(ctx context.Context, req *mcp.CallToolRequest, input Input) (
	*mcp.CallToolResult,
	Output,
	error,
) {
	generated := automata.Generate(input.Steps, input.Rule, input.Seed)

	var keys []uint
	for k := range generated {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	var sb strings.Builder
	for i, k := range keys {
		if i > 0 {
			sb.WriteString("\r\n")
		}
		for _, v := range generated[k] {
			sb.WriteString(strconv.Itoa(int(v)))
		}
	}

	return nil, Output{Automata: sb.String()}, nil
}

func main() {
	// Create a server with a single tool.
	server := mcp.NewServer(&mcp.Implementation{Name: "automata", Version: "v0.0.1"}, nil)
	mcp.AddTool(server, &mcp.Tool{Name: "generate", Description: "generate cellular automata"}, GenerateAutomata)
	// Run the server over stdin/stdout, until the client disconnects.
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}
