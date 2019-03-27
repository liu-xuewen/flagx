package main

import (

	"fmt"
	"github.com/liu-xuewen/flagx/fg"
)

func main() {

	// fg.String主要工作
	stringPRet := fg.String("test", "testValue", "testUsage")

	fmt.Println(*stringPRet) //在NewStringValue这里对p进行了赋值，所以返回指向3这个value的指针
}
