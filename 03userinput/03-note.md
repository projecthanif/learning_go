## User Input in Go

In this lesson, I learned how to take user input from the terminal in Go.

The program uses these packages:

- `fmt` for printing messages
- `os` to access standard input
- `bufio` to read input from the terminal

Example from `main.go`:

```go
reader := bufio.NewReader(os.Stdin)

fmt.Println("Enter the rating for our pizza:")

input, err := reader.ReadString('\n')
```

## What This Does

```go
bufio.NewReader(os.Stdin)
```

This creates a reader that can read text entered by the user in the terminal.

- `os.Stdin` means standard input
- `bufio.NewReader(...)` wraps it in a buffered reader

## Reading Input

```go
input, err := reader.ReadString('\n')
```

This reads the user's input until a newline character is found.

- `input` stores the text entered by the user
- `err` stores any error that happens while reading
- `'\n'` means the program will keep reading until the user presses Enter

## Error Handling

Your code checks for an error like this:

```go
if err != nil {
	fmt.Printf("Error : %v", err)
} else {
	fmt.Printf("Thanks for the rating %s", input)
}
```

This means:

- if an error happens, print the error
- otherwise, print the user's input

## What I Learned Here

- Go can take user input from the terminal
- `os.Stdin` is used to read from standard input
- `bufio.NewReader` helps read text input
- `ReadString('\n')` reads until Enter is pressed
- error handling is important when reading input

## One Simple Flow

1. Print a message to the user
2. Wait for the user to type something
3. Read the input
4. Check for errors
5. Print the result

## Example Summary

This lesson builds a simple terminal interaction:

- the program welcomes the user
- asks for a pizza rating
- reads the typed response
- prints the result back to the user
