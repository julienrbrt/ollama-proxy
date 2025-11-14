# Ollama Proxy

[![Test](https://github.com/julienrbrt/ollama-proxy/actions/workflows/test.yml/badge.svg)](https://github.com/julienrbrt/ollama-proxy/actions/workflows/test.yml)

A lightweight Go reverse proxy for Ollama with Bearer token authentication.

## Features

- Simple reverse proxy to Ollama API
- Bearer token authentication
- Easy configuration via environment variables

## Installation

```bash
go install github.com/julienrbrt/ollama-proxy@latest
```

```bash
go build -o ollama-proxy
```

## Configuration

The proxy is configured via environment variables:

- `OP_AUTH_TOKEN` (required): Bearer token for API authentication
- `OP_OLLAMA_URL` (optional): Ollama server URL (default: `http://localhost:11434`)
- `OP_PORT` (optional): Proxy server port (default: `11433`)

## Usage

### Start the proxy

```bash
export OP_AUTH_TOKEN="your-secret-token"
export OP_OLLAMA_URL="http://localhost:11434"
export OP_PORT="11433"
ollama-proxy
```
