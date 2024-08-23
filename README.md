# Go-HTTP

`go-http` is a simple HTTP server built in Go, specifically designed for testing purposes. It provides essential configurations and scripts to quickly set up and run the server. The project also includes automated tests and supports load testing with `k6`.

## Installation

### Preqrisites

- [Go](https://golang.org/doc/install) 1.XX or later
- [k6](https://k6.io/docs/getting-started/installation/) for load testing

### Clone the Repository

First, clone the repository to your local machine:

```bash
git clone thtps://github.com/your_username/go-http.git
cd go-http
```

### Set Up the Environment

```bash
# Configure Go Modules
go mod tidy
```

### Build the Project

```bash
make build
```

https://k6.io/docs/getting-started/installation/ to install `k6` on your system.

## Usage

### Running the Server

To run the server, use the generated binary:

```sh
./dist/app --help
```

Alternatively, you can use `make` to run the server:

```sh
make run-server
```

### Running Load Tests

To perform load testing with `k6`, use the included script:

```bash
k6 run k6.js
```

## Testing

To run unit tests for the project, use the following command:

```bash
make test
```

## Contributions

Contributions are welcome! If you find any issues or have examples you'd like to add, please create a pull request. See the [CONTRIBUTING](./CONTRIBUTING.md) file for more details.

## License

This project is licensed under the [MIT](./LICENSE) License. See the `LICENSE` file for more details.
