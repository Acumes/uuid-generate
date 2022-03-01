package utils

import (
	"github.com/go-basic/uuid"
	"github.com/shirou/gopsutil/host"
	"os"
	"strings"
)

func UuidGenerate(ids []string) string {
	path, _ := os.Getwd()
	var file_path = path + "/uuid.txt"
	//判断文件是否存在，如果存在，读取文件id
	fileExist, _ := PathExists(file_path)
	if fileExist {
		id, err := ReadFile(file_path)
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
		for _, val := range ids {
			val = strings.Replace(val, " ", "", -1)
			val = strings.Replace(val, "\n", "", -1)
			if val == hostNo {
				hostNo = uuid.New()
			}
		}
	}
	//写入文件
	WriteFile(file_path, []byte(hostNo), 0666)
	return hostNo

}
