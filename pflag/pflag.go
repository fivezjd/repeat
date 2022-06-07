/**
 * @Author: realpeanut
 * @Date: 2022/6/7 22:58
 */
package main

import (
	"fmt"
	"github.com/spf13/pflag"
)

func main() {
	demo8()
}

func demo1() {
	//支持长选项、默认值和使用文本，并将标志的值存储在指针中。

	var name = pflag.String("name", "tom", "input your name")
	fmt.Println(*name)
}

func demo2() {
	//支持长选项、短选项、默认值和使用文本，并将标志的值存储在指针中。
	var name = pflag.StringP("name", "n", "tom", "input your name")
	fmt.Println(name)
}

func demo3() {
	//支持长选项、默认值和使用文本，并将标志的值绑定到变量
	var name string
	pflag.StringVar(&name, "name", "tom", "input your name")
	fmt.Println(name)
}

func demo4() {
	//支持长选项、短选项、默认值和使用文本，并将标志的值绑定到变量。

	var name string
	pflag.StringVarP(&name, "name", "n", "tom", "Input Your Name")
}

func demo5() {
	// go run pflag.go arg1 arg2
	//解析命令行后面的参数
	pflag.Parse()
	//获取参数总数
	fmt.Printf("argument number is: %v\n", pflag.NArg())
	//获取参数值
	fmt.Printf("argument list is: %v\n", pflag.Args())
	//获取指定序列的参数
	fmt.Printf("the first argument is: %v\n", pflag.Arg(0))
}

func demo6() {
	//var ip = pflag.IntP("flagname", "f", 1234, "help message")
	pflag.Lookup("flagname").NoOptDefVal = "4567"
}

func demo7() {
	//Pflag 可以弃用标志或者标志的简写。弃用的标志或标志简写在帮助文本中会被隐藏，并在使用不推荐的标志或简写时打印正确的用法提示。例如，弃用名为 logmode 的标志，并告知用户应该使用哪个标志代替：
	err := pflag.CommandLine.MarkDeprecated("logmod", "please use --log-mod instead")
	if err != nil {
		return
	}
}

func demo8() {
	//保留名为 port 的标志，但是弃用它的简写形式。
	var port int
	pflag.IntVarP(&port, "port", "P", 3306, "MySQL service host port.")

	// deprecate a flag shorthand by specifying its flag name and a usage message
	err := pflag.CommandLine.MarkShorthandDeprecated("port", "please use --port only")
	if err != nil {
		return
	}
}

func demo9() {
	//可以将 Flag 标记为隐藏的，这意味着它仍将正常运行，但不会显示在 usage/help 文本中。例如：隐藏名为 secretFlag 的标志，只在内部使用，并且不希望它显示在帮助文本或者使用文本中。代码如下：

	// hide a flag by specifying its name
	err := pflag.CommandLine.MarkHidden("secretFlag")
	if err != nil {
		return
	}
}
