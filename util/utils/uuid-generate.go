package utils

import (
	"fmt"
	"github.com/go-basic/uuid"
	"github.com/shirou/gopsutil/host"
	"os"
)

var FILE_PATH = "/uuid.txt"

func UuidGenerate(ids []string) string {
	path, _ := os.Getwd()
	FILE_PATH = path + FILE_PATH
	//判断文件是否存在，如果存在，读取文件id
	fileExist, _ := PathExists(FILE_PATH)
	fmt.Println("fileExist ====== ", fileExist)
	if fileExist {
		id, err := ReadFile(FILE_PATH)
		if err == nil {
			return id
		}
	}
	hostNo, err := host.HostID()
	if err != nil {
		hostNo = uuid.New()
	}
	if len(ids) > 0 {
		//判断是否存在ids里面
		isExist := false
		for _, val := range ids {
			if val == hostNo {
				isExist = true
			}
			if isExist {
				hostNo = uuid.New()
			}
		}
		fmt.Println("isExist ====== ", fileExist)
	}
	fmt.Println("hostNo ====== ", hostNo)
	//写入文件
	WriteFile(FILE_PATH, []byte(hostNo), 0666)
	return hostNo

}
