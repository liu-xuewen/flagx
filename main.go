package main

import (

	"fmt"
	"github.com/liu-xuewen/flagx/fg"
)

func main() {

	// fg.String主要工作
	//1. 以value初始化了一个string指针.为什么要生成这个string指针
	//2. 主要调用了fg.Var, 生成一个新的Flag,放入FlagSet的formal,  map[string]*Flag
	stringPRet := fg.String("test", "testValue", "testUsage")

	fg.Parse()
	fmt.Println(*stringPRet) //在NewStringValue这里对p进行了赋值，所以返回指向3这个value的指针
}
