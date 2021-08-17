package Printlog

import (
	"fmt"
	"os"
)

// FmtDebug 向终端输出 虽然接收interface{}参数，但是参数必须是int 或 string，或是fmt.Sprintf("")
func(n *NewLog) FmtDebug(data ...interface{}) {
	if !n.IsDebug {
		return
	}
	str := timeStr("Debug", data)

	_, err := fmt.Fprintln(os.Stdout, colorStr(33, str))
	if err != nil {
		return
	}
}

// FileDebug 向log 文件输出 虽然接收interface{}参数，但是参数必须是int 或 string，或是fmt.Sprintf("")
func(n *NewLog) FileDebug(data ...interface{}) {
	if !n.IsDebug {
		return
	}
	str := timeStr("Debug", data)

	str += "\n"
	f, err := os.OpenFile(n.OutFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	defer f.Close()
	if err != nil {
		return
	}
	_, err = f.Write([]byte(str))
	if err != nil {
		return
	}

}
