## Learning


After installing Go from the official website, the basic first step is to initialize a module with `go mod init`.

Example:
`go mod init github.com/username/repo`

This command creates a Go module. A module is a collection of one or more packages.

Then I created a new file `main.go` and wrote the following code:
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

At the top of every Go file, you declare a package name.

- For an executable program, the package is usually `package main`.
- For reusable code, the package name is usually a short identifier like `utils`, `models`, or `mathops`.
- The module path like `github.com/username/repo` is not used as the `package` name inside source files.

To run a Go program, the entry point is:

```go
func main() {}
```

So the important difference is:

- `go mod init github.com/username/repo` defines the module path
- `package main` defines the package for that file
