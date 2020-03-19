package daemon

import (
	"os"
	"time"

	"github.com/sevlyar/go-daemon"
	params "github.com/wmentor/dsn"
)

func Run(dsn string) error {

	ds, err := params.New(dsn)
	if err != nil {
		return err
	}

	pid := ds.GetString("pid", "")
	dir := ds.GetString("dir", ".")
	lg := ds.GetString("log", "")

	var cntxt *daemon.Context

	cntxt = &daemon.Context{
		PidFileName: pid,
		PidFilePerm: 0644,
		LogFileName: lg,
		LogFilePerm: 0640,
		WorkDir:     dir,
		Umask:       027,
		Args:        os.Args,
	}

	child, err := cntxt.Reborn()
	if err != nil {
		return err
	}

	if child != nil {
		time.Sleep(time.Second)
		os.Exit(0)
	}

	return nil
}
