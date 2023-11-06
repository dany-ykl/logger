# Golang zap logger

The logger is written based on the package: https://github.com/uber-go/zap.

## Using
```go
// your project
// file main.go
package main

import (
	"github.com/dany-ykl/logger"
	"log"
)

func main() {
	err := logger.InitLogger(logger.Config{
		Namespace:   "test",
		Development: false,
		Filepath:    "", // <path/to/file/log.txt>
		Level:       logger.InfoLevel,
	})
	if err != nil {
		log.Fatalln(err)
	}

	Display()
}
```

```go
// your project
// file log.go
package main

import (
	"github.com/dany-ykl/logger"
	"go.uber.org/zap"
)

func Display() {
	logger.Info("display", zap.String("stringFiels", "string"))
}
```