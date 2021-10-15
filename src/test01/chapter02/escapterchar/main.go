package main

import "fmt" //fmt包中提供格式化，输出，输入函数

func main(){
	//转义字符
	fmt.Println("tom\tjack")

	fmt.Println("hello\nworld")

	fmt.Println("D:\\awesomeProject\\src\\test01")

	fmt.Println("tom\"I think I'd love you whatever however whenever\"")
	//\r 回车，从当前行的最前面开始输出，覆盖掉以前内容
	fmt.Println("天龙八部雪山飞狐\r张飞")

	fmt.Println("姓名 \t年龄\t籍贯\t住址\njohn\t12\t河北\t北京")

	num := 2 +4 *5
	fmt.Println("mun=", num)


}
