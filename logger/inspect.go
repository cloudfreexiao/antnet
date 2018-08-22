package logger

import (
	"bytes"

	"github.com/davecgh/go-spew/spew"
)

func init() {
	SetDebug(true)
}

// dump string for debug
func Inspect(dump ...interface{}) string {
	return spew.Sdump(dump...)
}

func TraceBack(dump ...interface{})  {
	if x := recover(); x != nil {
		var buff byte.Buffer
		buff.WriteString(fmt.Sprintf("dump:%v\n", x))
		i := 0
		funcName, file, line, ok := runtime.Caller(i)
		for ok {
			buff.WriteString(fmt.Sprintf("frame %v:[func:%v,file:%v,line:%v]\n", i, runtime.FuncForPC(funcName).Name(), file, line))
			i++
			funcName, file, line, ok = runtime.Caller(i)
		}

		buff.WriteString(fmt.Sprintf("DUMP DATA:%v\n", spew.Sdump(dump)))
		ERROR(buff.String())
	} else {
		ERROR("dump recover is nil:", spew.Sdump(dump))
	}
}