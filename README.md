# hello-agent

A simple Go web application that displays "Hello Atlassian from Github Copilot".

## Requirements

- Go 1.24 or later

## Running the Application

1. Build the application:
```bash
go build -o hello-agent .
```

2. Run the application:
```bash
./hello-agent
```

3. Access the application at `http://localhost:8080`

The application will display: **Hello Atlassian from Github Copilot**

## Implementation

This application uses only the Go standard library (`net/http`, `fmt`, and `log` packages) to create a simple HTTP server that serves the required message on port 8080.
