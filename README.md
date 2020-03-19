# daemon

Go demonization library. It's sevlyar/go-daemon wrapper.

# Summary

* Easy to use
* MIT license
* sevlyar/go-daemon wrapper

# Install

```
go get github.com/wmentor/daemon
```

# Usage

```go
package main

import (
  "github.com/wmentor/daemon"
)

func main() {

  config := "pid=/var/run/main.pid dir=/dir log=/var/log/myapp.log"
  if err := daemon.Run(config); err != nil {
    panic("daemon didn't run")
  }

  // your code 
}

```

You must call daemon.Run before exclusive usage of resources (bind ports, open files for write).

Any of the parameters (pid, log, dir) can be deleted from dsn string.

```go
package main

import (
  "github.com/wmentor/daemon"
)

func main() {

  // no pid and no log
  if err := daemon.Run(""); err != nil {
    panic("daemon didn't run")
  }

  // your code 
}
```
