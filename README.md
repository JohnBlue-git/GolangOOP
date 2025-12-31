# OOP Style Design and Package Design in Go

Go uses structs and interfaces to achieve object-oriented patterns, and packages to organize code into reusable modules.

## Example: OOP Style Design in Go

This example demonstrates how to use structs, interfaces, and methods in Go to achieve object-oriented programming patterns.

### Files

- [`main.go`](./main.go): Contains the Go source code example.

### How to Run

1. Navigate to the `OOP` directory:
    ```sh
    cd ...
    ```
2. Run the program:
    ```sh
    go run main.go
    ```

### Source Code Preview

```go
package main

import "fmt"

// Interface
type Shape interface {
     Area() float64
}

// Struct
type Rectangle struct {
     Width, Height float64
}

// Method (func with receiver)
func (r Rectangle) Area() float64 {
     return r.Width * r.Height
}

type Circle struct {
     Radius float64
}

func (c Circle) Area() float64 {
     return 3.14 * c.Radius * c.Radius
}

// Function using interface
func printArea(s Shape) {
     fmt.Println("Area:", s.Area())
}

func main() {
     r := Rectangle{Width: 3, Height: 4}
     c := Circle{Radius: 5}

     printArea(r)
     printArea(c)
}
```

- `struct` defines data structure.
- `interface` defines behavior.
- `func` with receiver implements methods.

The full code is available in [`main.go`](./main.go).

---

## Package Design in Go

Go packages are a fundamental way to organize and structure code. They provide modularity, reusability, and encapsulation.

### What is a Package?

A package is a collection of Go source files in the same directory that are compiled together. Packages serve several purposes:
- **Code Organization**: Group related functionality together
- **Encapsulation**: Control what's visible outside the package
- **Reusability**: Share code across different programs
- **Namespace Management**: Avoid naming conflicts

### Package Basics

1. **Package Declaration**: Every Go file starts with a `package` declaration
   ```go
   package main  // executable package
   package utils // library package
   ```

2. **Package Names**:
   - `main` package: Entry point for executable programs
   - Other names: Create library packages that can be imported

3. **Visibility Rules**:
   - **Exported** (public): Identifiers starting with uppercase letter (e.g., `Person`, `CalculateArea`)
   - **Unexported** (private): Identifiers starting with lowercase letter (e.g., `person`, `calculateArea`)

### Package Structure Example

```
OOP/                  # This directory is a self-contained Go module
├── go.mod            # Module definition (module name: oop-example)
├── main.go           # Main executable
├── models/           # Package for data models
│   └── shapes.go
├── utils/            # Package for utility functions
│   └── math.go
└── services/         # Package for business logic
    └── calculator.go
```

### Importing Packages

This project has its own `go.mod` file, making it a self-contained module called `oop-example`.

```go
import (
    "fmt"                      // Standard library package
    "oop-example/models"       // Local package (relative to module root)
    "oop-example/utils"        // Another local package
    "oop-example/services"     // Services package
)
```

**Important**: In Go, imports are always relative to the module root defined in `go.mod`, not relative to the current file location. The module name (`oop-example`) acts as the base path for all imports within this project.

### Package Design Best Practices

1. **Single Responsibility**: Each package should have a clear, focused purpose
2. **Minimal Dependencies**: Reduce coupling between packages
3. **Clear API**: Export only what's necessary
4. **Descriptive Names**: Use clear, concise package names (lowercase, no underscores)
5. **Avoid Circular Dependencies**: Package A shouldn't import package B if B imports A

### Package Files Structure

- [`models/shapes.go`](./models/shapes.go): Defines shape structures and interfaces
- [`utils/math.go`](./utils/math.go): Provides math utility functions
- [`services/calculator.go`](./services/calculator.go): Implements business logic for calculations

### Example: Using Multiple Packages

See the updated [`main.go`](./main.go) for a complete example of how packages work together.

### Key Concepts

#### 1. Package Initialization
```go
// init() function runs automatically when package is imported
func init() {
    fmt.Println("Package initialized")
}
```

#### 2. Internal Packages
- Packages in `internal/` directory are only importable by nearby code
- Example: `./internal/config` can only be imported by `OOP` and its sub-packages

#### 3. Module System (go.mod)
- Modern Go uses modules for dependency management
- Module path defines the import path prefix
- See the root `go.mod` file for module configuration

### How Packages Work Together

1. **Define**: Create packages with exported types and functions
2. **Import**: Import packages into other files
3. **Use**: Call exported functions and create exported types
4. **Build**: Go compiler links all packages together

### Running the Package Example

```sh
go run main.go
```
