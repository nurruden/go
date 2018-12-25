package tailf

import (
	"fmt"
	"github.com/hpcloud/tail"
	"xlog"
)

var (
	tailObj *tail.Tail
)

func Init(filename string) (err error) {

	tailObj, err = tail.TailFile(filename, tail.Config{
		ReOpen: true,
		Follow: true,
		Location: &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll: true,
	})
	if err != nil {
		xlog.LogError("tail file err:", err)
		return
	}
	return
}

func ReadLine()(msg *tail.Line,err error) {
	var ok bool
	msg,ok = <- tailObj.Lines
	if !ok {
		err = fmt.Errorf("Read line failed")
		return
	}
	return
}