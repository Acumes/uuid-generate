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
	fmt.Println("==========ids==========", ids)
	if len(ids) > 0 {
		//判断是否存在ids里面
		isExist := false
		for _, val := range ids {
			fmt.Println("==========val==========", val)
			if val == hostNo {
				fmt.Println("====================", true)
				isExist = true
			}
		}
		if isExist {
			hostNo = uuid.New()
		}
	}
	//写入文件
	WriteFile(FILE_PATH, []byte(hostNo), 0666)
	return hostNo

}
