# gologger

A simple and flexible logger interface abstraction for Go.

## Installation

```bash
go get github.com/kordar/gologger
```

## Usage

### Using the Standard Logger

`gologger` provides a default `StdLogger` which uses the standard library's `log` package.

```go
package main

import (
	"github.com/kordar/gologger"
)

func main() {
	// Use global functions directly
	logger.Info("This is an info message")
	logger.Infof("This is a formatted %s message", "info")

	logger.Warn("This is a warning message")
	logger.Error("This is an error message")
}
```

### Initializing a Custom Logger

You can initialize the global logger with any implementation of the `Logger` interface.

```go
import (
    "github.com/kordar/gologger"
    // import your logger adapter here
)

func main() {
    // myLogger := ... // Create your logger instance
    // logger.InitGlobal(myLogger)

    // logger.Info("Now using custom logger")
}
```

## Interface

The `Logger` interface defines the following methods:

- `WithField(key string, value interface{}) Logger`
- `WithFields(fields Fields) Logger`
- `Trace(args ...interface{})`
- `Tracef(format string, args ...interface{})`
- `Debug(args ...interface{})`
- `Debugf(format string, args ...interface{})`
- `Info(args ...interface{})`
- `Infof(format string, args ...interface{})`
- `Warn(args ...interface{})`
- `Warnf(format string, args ...interface{})`
- `Error(args ...interface{})`
- `Errorf(format string, args ...interface{})`
- `Panic(args ...interface{})`
- `Panicf(format string, args ...interface{})`
- `Fatal(args ...interface{})`
- `Fatalf(format string, args ...interface{})`
