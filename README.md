# 🛠️terminal.mcp

> A secure and pluggable MCP server to run terminal commands on your local machine or cloud server — remotely, safely, and with LLMs or agentic clients.

---

## ✨ Features

- 🖥️ Execute terminal commands securely via the [Model Context Protocol (MCP)](https://modelcontextprotocol.org)
- 🌐 Works locally or on cloud servers (run anywhere!)
- 🔒 Designed for safety with parameterized tools and isolated execution
- 🤖 Plug into LLMs, autonomous agents, or even your own apps
- 🧩 Simple Go SDK-based implementation with full MCP Inspector support

---

## 🚀 Getting Started

### 1. Clone the repo

```bash
git clone https://github.com/your-username/terminal.mcp.git
cd terminal.mcp
````

### 2. Run the MCP server

```bash
go run main.go
```

This starts the MCP server over `stdio`, waiting for clients (like ChatGPT, Claude, or custom clients) to connect.

---

## 🛠️ Tool: `run`

```json
{
  "command": "ls -la"
}
```

The server executes the command and returns the output. Supports safe execution via schema-validated tools.

---

## 🧪 Test with MCP Inspector

Use [MCP Inspector](https://github.com/modelcontextprotocol/inspector):

```bash
npx @modelcontextprotocol/inspector -- go run main.go
```

Then open [http://localhost:6274](http://localhost:6274) and test the `run` tool interactively.

---

## ☁️ Deploy on Cloud (optional)

You can run this server on any VPS, container, or cloud function. Just expose stdin/stdout or use a TCP transport (WIP).

---

## 📦 Built With

* [Go](https://golang.org)
* [MCP Go SDK](https://github.com/modelcontextprotocol/go-sdk)
* [MCP Inspector](https://github.com/modelcontextprotocol/inspector)

---

## 🔐 Security Notice

This server executes terminal commands. Use only with authenticated and trusted clients. Avoid deploying it publicly without a secure transport + auth layer.
