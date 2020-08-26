# go-logger

go-logger is a logger tool

Download:

```
go get github.com/Danile71/go-logger
```

### Usage

set DEBUG_LEVEL (0-7)

set DEBUG_URL (url to rest api, post method)

Example:

```go
package main

import (
	"github.com/Danile71/logger"
)

func main() {
	logger.Emergf("test %d %s", 123, "help")
	logger.Emerg("test ", 123, " help")

	logger.Alertf("test %d %s", 123, "help")
	logger.Alert("test ", 123, " help")

	logger.Critf("test %d %s", 123, "help")
	logger.Crit("test ", 123, " help")

	logger.Errorf("test %d %s", 123, "help")
	logger.Error("test ", 123, " help")

	logger.Warnf("test %d %s", 123, "help")
	logger.Warn("test ", 123, " help")

	logger.Noticef("test %d %s", 123, "help")
	logger.Notice("test ", 123, " help")

	logger.Infof("test %d %s", 123, "help")
	logger.Info("test ", 123, " help")

	logger.Debugf("test %d %s", 123, "help")
	logger.Debug("test ", 123, " help")
}
```