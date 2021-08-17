package Printlog

import (
	"fmt"
	"os"
)

// FmtPanic  向终端输出 虽然接收interface{}参数，但是参数必须是int 或 string，或是fmt.Sprintf("")
func(n *NewLog) FmtPanic(data ...interface{}) {
	str := timeStr("Panic", data)

	_, err := fmt.Fprintln(os.Stdout, colorStr(35, str))
	if err != nil {
		os.Exit(12)
	}
	os.Exit(12)

}

// FilePanic 向log 文件输出 虽然接收interface{}参数，但是参数必须是int 或 string，或是fmt.Sprintf("")
func(n *NewLog) FilePanic(data ...interface{}) {
	str := timeStr("Panic", data)

	str += "\n"
	f, err := os.OpenFile(n.OutFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	defer f.Close()
	if err != nil {
		n.FileError(str)
		os.Exit(12)
	}
	_, err = f.Write([]byte(str))
	if err != nil {
		os.Exit(12)
	}
	os.Exit(12)

}