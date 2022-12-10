# Utils

Utils is a package created for handling input in the terminal. It's similar to the `input()` function in Python.

### Installation

You can install the package using the command below

```bash
go get github.com/ixxiv/utils
```

### Import the package

```go
import "github.com/ixxiv/utils"
```

### Example

```go
package main

import "fmt"

func main() {
    // print out my github username

    ghUsername := utils.Input("GH Username")

    fmt.Printf("Welcome %s", ghUsername)

    // Prints: Welcome {the ghUsername you inputed}
}
```

### Contributions

PRs are accepted. You can also create a new issue if necessary.
