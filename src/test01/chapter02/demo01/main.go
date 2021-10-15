package main
import "fmt"

func main() {
	//定义变量/声明变量
	var i int
	//给i赋值
	i = 10
	//使用变量
	fmt.Println("i=", i)

	//golang变量的使用方式
	//指定变量类型，声明后若不赋值，使用默认值
	//int的默认值为0
	var j int
	fmt.Println("j=", j)

	//根据值自行判定变量类型
	var num = "liiia"
	fmt.Println("num=", num)
}
