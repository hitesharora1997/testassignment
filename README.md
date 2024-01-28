# Request Counter Server

Request Counter Server is a simple HTTP server written in Go that counts and tracks incoming requests within a 60-second window. It provides an API endpoint that returns the total number of requests made in the last 60 seconds. Additionally, it periodically saves the request data to a JSON file and restores data from that file when the server is restarted.

## Usage

### Build and Run

To build and run the server, follow these steps:

1. Ensure you have Go installed on your system. You can download it from [https://golang.org/dl/](https://golang.org/dl/).

2. Clone this repository:

   ```bash
   git clone https://github.com.git/hitesharora1997/testassignment
    ```

3. ```bash
   cd testassigment
    ```

4. ```bash
   go run cmd/main.go
    ```


### API Endpoint
To access the API endpoint that returns the total number of requests made in the last 60 seconds, make a GET request to http://localhost:8080/.