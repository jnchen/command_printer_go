package main


import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"io/ioutil"
	"gopkg.in/ini.v1"
)


func main(){
	// 读取配置文件
	cfg, _ := ini.Load("printer.ini")
	// 读取要执行的命令
	target_program := cfg.Section("Command").Key("Target").String()
	// 读取参数记录文件
	target_log := cfg.Section("Command").Key("Record").String()

	args := os.Args

	print_args := strings.Join(args, " ")
	_ = ioutil.WriteFile(target_log, []byte(print_args), 0644)

	cmd := exec.Command(target_program, args[1:]...)
	output, _ := cmd.Output()
	fmt.Println(string(output))
}
