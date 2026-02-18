# hello-agent

A simple web application that displays "Hello Atlassian from Github Copilot" in large font.

## Requirements

- Go 1.x or higher

## Building

```bash
go build -o hello-agent .
```

## Running

```bash
./hello-agent
```

The application will start a web server on `http://localhost:8080`.

## Implementation

This application uses only the Go standard library:
- `net/http` for the HTTP server
- `fmt` for formatting output
- `log` for logging

The web page displays the greeting message centered on the page with a large font size (72px).
