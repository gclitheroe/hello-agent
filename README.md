# hello-agent

A simple web application that displays "Hello Atlassian from Github Copilot" in large font.

## Requirements

- Go 1.16 or later

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

You should see "Hello Atlassian from Github Copilot" displayed in large font.

## Development

This application uses only the Go standard library and serves a simple HTML page with embedded CSS for styling.
