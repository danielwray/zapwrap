# Zap Wrap || Golang Zap Logger Package

Zap Wrap or, Golang Zap Logger package, is a pre-configured Zap Logger package that can be imported into Go applications.

The package name is ```gzl``` to make referencing throughout code a little easier.

## Example

```golang
package main

import (
    "github.com/danielwray/zapwrap/gzl"
)

func main() {
    gzl.Info("Some useful message")
}
```