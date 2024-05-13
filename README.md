# 程序运行参数打印机

许多软件在运行过程中，会调用一些第三方的程序来实现它的业务逻辑。通常我们可以直接在软件的安装目录找到对应的一些第三方程序，但是我们却不知道主程序是如何调用的他们。

## 使用方法
- 复制配置文件`printer.ini.template` 为`printer.ini`，并将和执行文件(`command_printer_go.exe`)放入想要知道运行参数的程序所在的目录中
- 将原程序改名，并把新的名字配置到配置文件的Target
- 将`command_printer_go.exe` 修改为原程序名
- 启动软件，然后查看`command_log.txt`

