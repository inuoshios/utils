# Utils

Utils is a package created for handling input in the terminal. It's similar to the `input()` function in Python.

### Installation

You can install the package using the command below

```bash
go get github.com/inuoshios/utils
```

### Import the package

```go
import "github.com/inuoshios/utils"
```

### Example

```go
package main

import "fmt"

func main() {
    // print out my github username

    ghUsername, err := utils.Input("GH Username")
    if err != nil {
      fmt.Printf("error reading github username %v\n", err)
    }

    fmt.Printf("Welcome %s", ghUsername)

    // Prints: Welcome {the ghUsername you inputed}
}
```

### Contributions

PRs are accepted. You can also create a new issue if necessary.
