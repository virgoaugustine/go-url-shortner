# URL Shortener

This project is a URL shortener service written in Go. It allows users to shorten URLs and retrieve the original URLs using the shortened keys.

## Features

- Shorten URLs and generate unique keys.
- Retrieve original URLs using the shortened keys.
- Save and load URLs from a JSON file.
- Optionally enable an RPC server for remote procedure calls. (Yet to work on this)

## Requirements

- Go 1.16 or later

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/virgoaugustine/url-shortener.git
   cd url-shortener
   ```

2. Build the project:
   ```sh
   go build
   ```

## Usage

1. Run the server:
   ```sh
   ./url-shortener -port=4000 -file=urls.json -rpc=false
   ```

2. Add a URL:
    - Open a browser and go to `http://localhost:4000/add`
    - Enter the URL you want to shorten and submit the form.

3. Redirect to the original URL:
    - Use the shortened URL key in the browser, e.g., `http://localhost:4000/u0hleB`

## Command-Line Flags

- `-port`: Port to run the server (default: `4000`)
- `-file`: File to store the saved URLs (default: `urls.json`)
- `-rpc`: Enable RPC server (default: `false`)

## Project Structure

- `main.go`: Entry point of the application.
- `store.go`: Contains the `URLStore` struct and methods for managing URLs.
- `key.go`: Contains the method  for generating unique keys.

