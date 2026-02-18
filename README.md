# hello-agent

A simple web application that displays "Hello Atlassian from Github Copilot" in large font.

## Requirements

- Go 1.16 or later

## Building

```bash
go build -o hello-agent main.go
```

## Running

```bash
./hello-agent
```

The application will start a web server on `http://localhost:8080`.

Open your browser and navigate to `http://localhost:8080` to see the greeting message.

## Implementation Details

This application uses only the Go standard library:
- `net/http` for the web server
- `html/template` for HTML templating
- `log` for logging
