# Uninitialized Map Access Panic in Go

This repository demonstrates a common error in Go: accessing a map before it has been initialized. This results in a runtime panic due to a nil pointer dereference.

## The Bug

The `bug.go` file contains code that attempts to assign a value to a map without first initializing the map. This causes a panic because the map variable is `nil` and cannot be used until it is properly initialized.

## The Solution

The `bugSolution.go` file provides a corrected version that initializes the map before assigning a value.  This prevents the panic and allows the code to execute successfully.

## How to Reproduce

1. Clone this repository.
2. Navigate to the directory.
3. Run `go run bug.go`.  This will result in a panic.
4. Run `go run bugSolution.go`. This will execute without error.