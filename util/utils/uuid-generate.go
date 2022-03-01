package utils

import (
	"fmt"
	"github.com/go-basic/uuid"
	"github.com/shirou/gopsutil/host"
	"os"
	"strings"
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
		for _, val := range ids {
			val = strings.Replace(val, " ", "", -1)
			val = strings.Replace(val, "\n", "", -1)
			fmt.Println("==========val==========", val, "========", val == hostNo, "=========hostNo=========", hostNo)
			if val == hostNo {
				fmt.Println("====================", true)
				hostNo = uuid.New()
			}
		}
	}
	//写入文件
	err = WriteFile(FILE_PATH, []byte(hostNo), 0666)
	if err != nil {
		fmt.Println("errrr==============", err)
	}
	return hostNo

}
