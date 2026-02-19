# hello-agent

A simple Golang web application that displays "Hello Atlassian from Github Copilot".

## Requirements

- Go 1.24.13 or higher

## Running the Application

1. Build the application:
```bash
go build -o hello-agent main.go
```

2. Run the application:
```bash
./hello-agent
```

3. Open your browser and navigate to `http://localhost:8080`

You should see the message: "Hello Atlassian from Github Copilot"

## Development

The application uses only the Go standard library and creates a simple HTTP server on port 8080.
