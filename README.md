# go-logger

go-logger is a logger tool

Download:

```
go get github.com/Danile71/go-logger
```

### Usage

set DEBUG_LEVEL (0-7)

set DEBUG_URL (url to rest api, post method)

set DEBUG_TIME_FORMAT ("Jan _2 15:04:05")

Example:

```go
package main

import (
	"errors"

	"github.com/Danile71/go-logger"
)

func main() {
	logger.SetTimeFormat("Jan _2 15:04:05")

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

	logger.Logf(logger.ERR, "test %d %s", 123, "help")
	logger.Log(logger.ERR, "test ", 123, " help")

	var err error
	if logger.OnError(err) {
		panic("err")
	}
	err = errors.New("test error")
	if !logger.OnError(err) {
		panic("err")
	}
}
```