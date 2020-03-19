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

	if err := daemon.Run("pid=/var/run/main.pid dir=/home/user log=/var/log/myapp/myapp.log"); err != nil {
		panic("daemon didn't run")
	}

	// your code 
}

```

You must call daemon.Run before exclusive usage of resources (bind ports, open files for write).

Any of the parameters (pid, log, dir) can be deleted.
