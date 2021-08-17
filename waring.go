package Printlog

import (
	"fmt"
	"os"
)

// FmtWaring 向终端输出 虽然接收interface{}参数，但是参数必须是int 或 string，或是fmt.Sprintf("")
func(n *NewLog) FmtWaring(data ...interface{}) {
	str := timeStr("Waring", data)

	_, err := fmt.Fprintf(os.Stdout, colorStr(36, str))
	if err != nil {
		return
	}
}

// FileWaring 向log 文件输出 虽然接收interface{}参数，但是参数必须是int 或 string，或是fmt.Sprintf("")
func(n *NewLog) FileWaring(data ...interface{}) {
	str := timeStr("Waring", data)

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
