# Golang Learning Roadmap

This roadmap starts from Hitesh Choudhary's playlist, ["Let's go with golang"](https://www.youtube.com/playlist?list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa), and extends it into a more complete path for modern Go development.

The playlist is a strong beginner series, but it is older, so this roadmap adds the topics you should learn after the course, especially modules-first workflows, testing, interfaces, `context`, and generics.

## What this roadmap is built around

- Playlist: `Let's go with golang`
- Creator: Hitesh Choudhary
- Size: 57 videos
- Playlist page notes: last updated on 13 Sept 2021

## How to use this roadmap

- Watch 2 to 4 videos at a time.
- Rebuild the code from memory after each session.
- Keep one folder for experiments and one Markdown file for notes.
- Do the checkpoint project before moving to the next phase.
- Spend more time coding than watching.

## Important note before you start

One playlist lesson covers `GOPATH`. Learn it for background, but for real projects today you should mainly use Go modules with `go mod init` and `go mod tidy`.

## Phase 0 - Setup and learning habits

### Goal

Get your environment ready and make sure you learn by building, not only by watching.

### Do this first

- Install Go from the official site.
- Install VS Code or GoLand.
- Create a folder called `go-playground` for small exercises.
- Create a GitHub repo for your learning projects.
- Start a notes file: `golang-notes.md`

### Practice

- Write a tiny `hello world`
- Run a program from the terminal
- Create your first module with `go mod init`

### Outcome

You can run Go code locally and you have a place to save exercises and notes.

## Phase 1 - Core syntax and basic mental model

### Playlist section

Videos 1 to 12

- Welcome to series on GO programming language
- Before you start with golang
- Golang installation and hello world
- GOPATH and reading go docs
- Lexer in golang and Types
- Variables, types and constants
- Comma ok syntax and packages in golang
- Conversions in golang
- Handling time in golang
- Build for windows, linux and mac
- Memory management in golang
- Pointers in golang

### Learn these ideas

- How Go code is organized into packages
- Basic types, zero values, constants, and variables
- Short declaration vs explicit declaration
- Type conversion
- Basic memory model and pointers
- Reading docs instead of guessing

### Mini project

Build a small CLI profile app that:

- asks for a name and age
- prints the current time
- shows a formatted summary
- uses constants, functions, and type conversion

### Exit criteria

- You understand variables, types, constants, and functions
- You are not confused by pointers anymore
- You can read package docs on your own

## Phase 2 - Data structures and control flow

### Playlist section

Videos 13 to 24

- Array in golang
- Slices in golang
- How to remove a value from slice based on index in golang
- Maps in golang
- Structs in golang
- If else in golang
- Switch case in golang and online playground
- Loop break continue and goto in golang
- Functions in golang
- Methods in golang
- Defer in golang
- Working with files in golang

### Learn these ideas

- When to use arrays, slices, maps, and structs
- Value types vs reference-like behavior with slices and maps
- Loops and branching
- Functions vs methods
- `defer` for cleanup
- Basic file I/O

### Mini project

Build a CLI task manager that:

- stores tasks in a slice of structs
- groups tasks by status with a map
- reads and writes tasks to a local file
- supports add, list, complete, and delete

### Exit criteria

- You can model data with structs
- You understand slices well enough to modify them safely
- You can read from and write to files

## Phase 3 - HTTP, JSON, and modern modules

### Playlist section

Videos 25 to 33

- Handling web request in golang
- Handling URL in golang
- Creating server for golang frontend
- How to make GET request in golang
- How to make POST request with JSON data in golang
- How to send form data in golang
- How to create JSON data in golang
- How to consume JSON data in golang
- A long video on MOD in golang

### Learn these ideas

- Building basic HTTP servers with `net/http`
- Working with routes, URLs, query params, and request bodies
- Encoding and decoding JSON
- Calling external APIs
- Managing dependencies with Go modules

### Mini project

Build a weather or quote API client that:

- calls a public API
- decodes JSON into structs
- handles errors cleanly
- saves the last response to a file

### Exit criteria

- You can build and call HTTP endpoints
- You understand request and response flow
- You can use modules confidently

## Phase 4 - CRUD APIs and database basics

### Playlist section

Videos 34 to 50

- Building API in golang - Models
- Sending a API json response for all courses in golang
- Get one course based on request id in golang
- Add a course controller in golang
- Update a course controller in golang
- Delete a course controller in golang
- Handling routes and testing routes in golang
- MongoDB setup for API in golang
- Defining models for netflix in golang
- Making a connection to database in golang
- Insert data in mongodb in golang
- Update a record in mongodb in golang
- Delete one and delete many in mongodb in golang
- Get all collection in mongodb in golang
- Get all movies from DB in golang
- Mark movie as watched in golang
- Delete 1 and all movie in golang

### Learn these ideas

- REST API structure
- Models, handlers, and routes
- CRUD patterns
- Database connection flow
- JSON responses for clients
- Separating app logic from transport logic

### Mini project

Build a movie, book, or note API with:

- CRUD endpoints
- request validation
- JSON responses
- MongoDB or any database you are comfortable with

### Exit criteria

- You can build a basic backend API in Go
- You understand how handlers, models, and routes fit together
- You can persist data outside memory

## Phase 5 - Concurrency and safe shared state

### Playlist section

Videos 51 to 57

- Creating routes and testing API in golang
- Concurrency and goroutines in golang
- Wait groups in golang
- Mutex in golang
- Race Condition in golang
- Channels and Deadlock in golang
- Math, crypto and random number in golang

### Learn these ideas

- Goroutines for concurrent work
- Coordinating work with `sync.WaitGroup`
- Protecting shared state with `sync.Mutex`
- Race conditions and deadlocks
- Channels for communication
- Basic standard library utilities

### Mini project

Build a concurrent URL checker that:

- checks many URLs at once
- collects results safely
- uses wait groups and channels
- reports failures and response times

### Exit criteria

- You can explain when to use goroutines, channels, and mutexes
- You can recognize deadlock and race-condition risks

## Phase 6 - Modern Go topics the playlist does not cover enough

This is the most important extension after the playlist.

### 1. Error handling

Learn:

- returning errors properly
- wrapping errors with `fmt.Errorf`
- checking errors with `errors.Is` and `errors.As`

Practice:

- refactor your earlier projects so every error is handled clearly

### 2. Interfaces and composition

Learn:

- small interfaces
- implicit interface implementation
- composition over inheritance

Practice:

- define an interface for storage, then swap file storage for in-memory storage

### 3. Testing

Learn:

- unit tests with the `testing` package
- table-driven tests
- subtests
- benchmarks

Practice:

- add tests to your task manager and API handlers

### 4. `context` and timeouts

Learn:

- request-scoped cancellation
- deadlines and timeouts
- passing `context.Context` through layers

Practice:

- add timeouts to outgoing HTTP calls and database operations

### 5. Generics

Generics arrived after this playlist was made, so you should learn:

- type parameters
- reusable generic helpers
- when generics help and when they do not

### 6. Project structure

Learn:

- how to split packages sensibly
- keeping `main` thin
- separating handlers, services, repositories, and models

### 7. Databases beyond MongoDB

The playlist leans MongoDB-heavy. Also learn:

- SQL basics
- `database/sql`
- migrations
- basic query design

### 8. Logging, configuration, and deployment

Learn:

- environment variables
- structured logging
- Docker basics
- how to run a Go app in production

## Suggested 8-week pace

| Week | Focus | Output |
| --- | --- | --- |
| 1 | Phase 0 + Phase 1 | Small CLI profile app |
| 2 | Finish Phase 2 | CLI task manager |
| 3 | Finish Phase 3 | JSON API client |
| 4 | Start Phase 4 | Simple CRUD API |
| 5 | Finish Phase 4 | CRUD API with database |
| 6 | Phase 5 | Concurrent URL checker |
| 7 | Testing + interfaces + errors | Refactor and test previous projects |
| 8 | Context + generics + project structure | Clean final backend project |

## Project ladder

Build these in order:

1. `hello-cli`
2. `profile-summary-cli`
3. `task-manager-cli`
4. `json-api-client`
5. `courses-api`
6. `movies-api-with-db`
7. `concurrent-url-checker`
8. `final-go-backend`

## Final target

By the end of this roadmap, you should be able to:

- build command-line apps in Go
- read and write files
- model data with structs and methods
- build HTTP APIs
- work with JSON and databases
- use goroutines, channels, mutexes, and wait groups safely
- write tests for your Go code
- structure a small to medium backend project

## Best learning loop

For every topic, do this:

1. Watch the lesson.
2. Rebuild the code without copying.
3. Change something in the example.
4. Write down what confused you.
5. Build a tiny project using that idea.

## Recommended references after the playlist

- [Official Go getting started tutorial](https://go.dev/doc/tutorial/getting-started)
- [A Tour of Go](https://go.dev/tour/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go package documentation](https://pkg.go.dev/std)
- [The Go language specification](https://go.dev/ref/spec)

## If you want the shortest possible path

Use this order:

1. Finish the playlist.
2. Build one CLI app.
3. Build one REST API.
4. Learn testing.
5. Learn interfaces and `context`.
6. Learn concurrency properly.
7. Build one final project from scratch without tutorial help.
