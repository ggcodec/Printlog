package Printlog

import (
	"fmt"
	"os"
)

// FmtError 向终端输出 虽然接收interface{}参数，但是参数必须是int 或 string，或是fmt.Sprintf("")
func(n *NewLog) FmtError(data ...interface{}) {
	str := timeStr("Error", data)

	_, err := fmt.Fprintf(os.Stdout, colorStr(31, str))
	if err != nil {
		return
	}
}

// FileError 向log 文件输出 虽然接收interface{}参数，但是参数必须是int 或 string，或是fmt.Sprintf("")
func(n *NewLog) FileError(data ...interface{}) {
	str := timeStr("Error", data)

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
