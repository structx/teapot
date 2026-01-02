# 🍵 Teapot

Teapot is a high-performance, zero-allocation structured logging library for Go. Built with zero external dependencies, it utilizes an optimized sync.Pool and an explicit attribute system to provide a high-speed logging experience without the overhead of reflection or interface boxing.

## 📦 Installation

```bash
go get github.com/structx/teapot
```

## 🚀 Key Features

Zero Allocations: Optimized to hit 0 B/op and 0 allocs/op in standard hot paths.\
Explicit Type System: Uses a "Fat Attr" struct to avoid reflection and unnecessary heap escapes.\
Native JSON Support: High-speed JSON output designed for production log aggregators.\
Built-in Stack Traces: Automatic capture of runtime.Stack for ERROR and FATAL levels.

## ⚡ Performance

```bash
goos: linux
goarch: amd64
pkg: github.com/structx/teapot
cpu: Intel(R) Core(TM) i7-10850H CPU @ 2.70GHz
BenchmarkLogger_JSON-12         70972780               171.4 ns/op             0 B/op          0 allocs/op
PASS
ok      github.com/structx/teapot       12.340s
```

## 🛠 Usage
Strongly Typed Attributes

Teapot avoids the overhead of interface{} by providing explicit attribute constructors. This ensures your logs are fast and type-safe.
```go
package main

import (
    "github.com/structx/teapot"
)

func main() {
    log := teapot.New(
        teapot.WithLevel(teapot.INFO),
    )

    log.Infof("user login attempt", 
        teapot.String("user", "alice"),
        teapot.Int("attempts", 1),
        teapot.Bool("success", true),
    )
}
```

## Error Handling & Stack Traces

Teapot automatically captures a full stack trace when logging at the ERROR or FATAL level. Use the Err helper for consistent naming:
```go
if err := db.Connect(); err != nil {
    // This will include a "stacktrace" field in the JSON output automatically
    log.Error("database failure", teapot.Err(err))
}
```

## ⚙️ Configuration
Option	Description	Default\
WithLevel(lvl)	Sets minimum logging threshold	`DEBUG`\
WithWriter(w)	Destination (file, stdout, etc)	os.Stdout

## ⚖️ License
MIT License. See [License](LICENSE) for details.