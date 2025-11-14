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

- `AUTH_TOKEN` (required): Bearer token for API authentication
- `OLLAMA_URL` (optional): Ollama server URL (default: `http://localhost:11434`)
- `PORT` (optional): Proxy server port (default: `8080`)

## Usage

### Start the proxy

```bash
export AUTH_TOKEN="your-secret-token"
export OLLAMA_URL="http://localhost:11434"
export PORT="8080"
ollama-proxy
```
