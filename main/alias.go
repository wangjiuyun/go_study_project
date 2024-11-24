package main

//如果导入一个包却没有使用它，则会在构建程序时发生错误
import fm "fmt" //alias3

const hardEight = (1 << 100) >> 97

func main() {
	fm.Println("hello,world")
	fm.Println(hardEight)
}
