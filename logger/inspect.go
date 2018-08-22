package logger

import (
	spew "github.com/davecgh/go-spew/spew"
)

func init() {
	SetDebug(true)
}

// dump for debug
func DUMP(v ...interface{})  {
	DEBUG(spew(v...))
}