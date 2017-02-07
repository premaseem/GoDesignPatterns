package singleton

import "fmt"

type singleton struct {

}

var instance *singleton

func GetInstance() *singleton {

	if instance == nil {
		instance = &singleton{}
		fmt.Println(" Created first instance ")
	}
	return instance
}