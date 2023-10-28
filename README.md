# hotelbookd
Hotel reservation backend written in go

This document provides an overview of the available Makefile commands and how to set up this project.

## Prerequisites

1. **Clone the Repository:** Start by cloning the repository to your local machine:
   ```bash
   git clone https://github.com/CosmoBean/hotelbookd.git
   ```

2. **Docker and Docker Compose:** Ensure Docker and Docker Compose are installed.

3. **Go Programming Language:** Ensure you have Go version 1.20 or higher installed.

4. **Environment Variables:** Create a `.env` file in the root directory. Copy the contents from `.env.sample` and replace the placeholder values with your actual configurations.

## Available Commands

### `start-basic-containers`

Boots up the basic containers required for the project.

**Usage:**
```bash
make start-basic-containers
```

### `stop-basic-containers`

Shuts down the basic containers and removes any orphaned containers.

**Usage:**
```bash
make stop-basic-containers
```

### `run-server`

Runs the main Go program.

**Usage:**
```bash
make run-server
```

### `build`

Builds the Go binary and places it in the `bin` directory with the name `api`.

**Usage:**
```bash
make build
```

### `run`

Builds the Go binary and then executes it.

**Usage:**
```bash
make run
```

### `test`

Runs all the tests available in the project.

**Usage:**
```bash
make test
```

## Notes

- Ensure Docker is running when using the `start-basic-containers` and `stop-basic-containers` commands.
- Ensure Go is correctly installed and set up for the `run-server`, `build`, `run`, and `test` commands.
- Always check the `.env` file and ensure it's correctly set up with your configurations.

For any issues or further queries, please refer to the project documentation or contact the project maintainers.
