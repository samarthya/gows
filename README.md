# Command line arguments
Add command line parameters to display a simple banner utility

```
cmd -m="My message" -n 2

```

# Steps to execute
- `cmd.go` package name can be changed to main and execute the file.
- Added `main.go` to invoke the Fileread.

# Start Process
- Using the `os` packages [`StartProcess`](https://golang.org/pkg/os/#StartProcess) method.
- StartProcess is a low-level interface
- If there is any error, it will of type *PathError