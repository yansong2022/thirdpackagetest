package main

import (
	"github.com/yansong2022/thirdpackagetest/mymethod"
	"github.com/yansong2022/thirdpackagetest/mystruct"
)

func main() {
	myservice := mystruct.NewService("test", "ys", 4)
	myservice.PrintService()
	mymethod.TestMethod()
}
