# TaskX

TaskX is a simple CLI task manager that lets you create, view and update tasks. This project was implemented as part of a case study
to illustrate my approach to solving a problem.

## Technology

- Golang
- Cobra CLI
- UberFX
- gRPC
- MySQL
- Docker

## Structure

This project consists of 2 major parts: the CLI application and the gRPC server environment.

Everything is structured around the CLI application as it controls the starting and stopping of the server. The gRPC server
exists as it's own Go module but in standard use it is run in a Docker container alongside a MySQL database using docker compose.


## Usage

1. Clone this repository
2. Ensure that your go version is at least `1.22.0` and that you have docker installed locally
3. Run `go build` to build the CLI application
4. The resulting binary can be run with the following commands:

   - `./taskx server` starts the docker containers
   - `./taskx server shutdown` stops the docker containers
   - `./taskx client get` returns a list of your tasks
   - `./taskx client create -d "Some Task"` creates a new task with the description "Some Task"
   - `./taskx client complete 3` marks the task with ID 3 as complete



 
