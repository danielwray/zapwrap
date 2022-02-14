# Zap Wrap (Golang Zap Logger Package)

Zap Wrap or, Golang Zap Logger package, is a pre-configured Zap Logger package that can be imported into Go applications.

## Example

```golang
package main

import (
    "github.com/danielwray/zapwrap"
)

func main() {
    zapwrap.Info("Some useful message")
}
```