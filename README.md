# Go Automata MCP

[![Go Version](https://img.shields.io/badge/Go-1.26-blue?logo=go)](go.mod)

A simple [Model Context Protocol (MCP)](https://modelcontextprotocol.io) server built in Go. It exposes a `greet` tool that returns a personalised greeting.

## Quick Start

```bash
go build -o automata-mcp .
```

## Usage

This server communicates over **stdio transport**, so it's compatible with any MCP client (Claude Desktop, opencode, etc.).

Add it to your MCP client configuration:

```json
{
  "mcpServers": {
    "greeter": {
      "command": "/path/to/automata-mcp"
    }
  }
}
```

## Tools

| Tool    | Description | Input                                    | Output                                     |
| ------- | ----------- | ---------------------------------------- | ------------------------------------------ |
| `greet` | Say hi      | `name` (string) – the person to greet    | `greeting` (string) – the greeting message |

## Development

```bash
go build -o automata-mcp .
```

## Debugging

Use the [MCP Inspector](https://github.com/modelcontextprotocol/inspector) to interactively test and debug the server.

Build the binary first, then launch the inspector:

```bash
go build -o automata-mcp .
npx @modelcontextprotocol/inspector ./automata-mcp
```

This opens a web UI at `http://localhost:6274` where you can list tools, call the `greet` tool, and inspect requests and responses.

## License

MIT
