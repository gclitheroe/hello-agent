# hello-agent

A simple Go web application that displays "Hello Atlassian from Github Copilot" in large font.

## Requirements

- Go 1.x or later

## Running the application

### Option 1: Run directly with go run

```bash
go run main.go
```

### Option 2: Build and run the binary

```bash
go build -o hello-app main.go
./hello-app
```

The server will start on port 8080. Visit http://localhost:8080 in your browser to see the greeting message.

## Implementation

This application uses only the Go standard library:
- `net/http` for the HTTP server
- `fmt` and `log` for output and logging

The web page displays "Hello Atlassian from Github Copilot" in large font (72px) with centered styling.
