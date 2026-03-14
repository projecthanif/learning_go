## Learning


So basically after installing golang from the official website, a straight forward process 
learned about `go mod init ${anything but most of the time its like this github.com/username/repo` a go command use to initialize a new module. a module is a collection of packages.
Then created a new file `main.go` and wrote the following code:
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```


At the start of every go file, you need to declare the package name using `package main`. not  `main` most of the time its like this `github.com/username/repo` for the main file and `github.com/username/repo/pkg` for package files.
just like other programming languages go starts with ```go func main(){} ```
