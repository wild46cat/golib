package common

import (
	"fmt"
	"os"
	"time"
)

func init() {
}

type RollingFile struct {
	f          *os.File
	fileprefix string
	formatStr  string
	filesuffix string
}

func NewRollingFile(filePrefix, formatStr, fileSuffix string) (*RollingFile, error) {
	formatTimeStr := time.Now().Format(formatStr)
	file, err := os.OpenFile(filePrefix+formatTimeStr+fileSuffix, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	rollingfile := &RollingFile{
		f:          file,
		fileprefix: filePrefix,
		formatStr:  formatStr,
		filesuffix: fileSuffix,
	}
	return rollingfile, nil

}

func (rf *RollingFile) Write(p []byte) (n int, err error) {
	//判断文件是否需要更换
	formatTimeStr := time.Now().Format(rf.formatStr)
	if rf.f.Name() != formatTimeStr {
		rf.f.Close()
		file, err := os.OpenFile(rf.fileprefix+formatTimeStr+rf.filesuffix, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			fmt.Printf("error create file")
		}
		rf.f = file
	}
	return rf.f.Write(p)
}
