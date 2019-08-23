package main

import (
	"Golang-lee/learngo/crwaler/engine"
	"Golang-lee/learngo/crwaler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
