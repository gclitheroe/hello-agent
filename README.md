# hello-agent

A simple web application that displays "Hello Atlassian from GitHub Copilot" in large font.

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

3. Open your browser and navigate to:
   ```
   http://localhost:8080
   ```

You should see "Hello Atlassian from GitHub Copilot" displayed in large font.

## Development

This application uses only the Go standard library and serves a simple HTML page with embedded CSS for styling.
