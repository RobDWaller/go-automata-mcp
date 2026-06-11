# Go Automata MCP

[![Go Version](https://img.shields.io/badge/Go-1.26-blue?logo=go)](go.mod)

A [Model Context Protocol (MCP)](https://modelcontextprotocol.io) server that generates [Wolfram elementary cellular automata](https://en.wikipedia.org/wiki/Elementary_cellular_automaton). Built in Go using the [`go-automata`](https://github.com/robdwaller/go-automata) library.

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
    "automata": {
      "command": "/path/to/automata-mcp"
    }
  }
}
```

## Tools

| Tool       | Description                     | Inputs                                                                  | Output                                       |
|------------|---------------------------------|-------------------------------------------------------------------------|----------------------------------------------|
| `generate` | Generate cellular automata      | `steps` (int), `rule` (int 0-255), `seed` (string of `0`/`1`)          | `automata` (string) – generations as rows    |

The `generate` tool runs a one-dimensional cellular automaton with periodic (wrap-around) boundaries. The `rule` parameter selects one of Wolfram's 256 rules (e.g., Rule 30, Rule 90, Rule 110). The `seed` is a string of `0` and `1` characters representing the initial row. The output is returned as lines of `0`/`1` characters, one per generation.

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

This opens a web UI at `http://localhost:6274` where you can list tools, call the `generate` tool, and inspect requests and responses.

## License

MIT
