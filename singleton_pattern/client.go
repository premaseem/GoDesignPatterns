package main

import (
	"fmt"
	"github.com/premaseem/GoDesignPatterns/singleton_pattern/singleton"
)

func main()  {
	fmt.Print("Singleton Design Pattern")
         instance1 := singleton.GetInstance()
         instance2 := singleton.GetInstance()

	if instance1 == instance2 {
		fmt.Print(" Success : Only single instance was created and returned ")
	} else {
		fmt.Print(" Failue : Singleton pattern failed, multiple instances got created ")
	}
}