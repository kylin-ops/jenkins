package main

import (
	"fmt"
	"github.com/kylin-ops/jenkins"
)

func main() {
	//	ctx := context.Background()
	jenkin, _ := jenkins.NewClient("http://192.168.31.220:9090", "admin", "admin")
	// fmt.Println(jenkin.JobBuild("test", map[string]string{"env":"test"}))
	fmt.Println(jenkin.JobGetConfig("test"))
}
