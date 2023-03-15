package main

import (
	"flag"
	"fmt"
	"os"
)

const TplFile = "friend.tpl"

func main() {
	var inputFilename, friendFilename string

	flag.StringVar(&inputFilename, "in", "input.md", "issueFilename")
	flag.StringVar(&friendFilename, "friend", "contents/friend.md", "origin friend.md filename")
	flag.Parse()

	content, err := os.ReadFile(inputFilename)
	if err != nil {
		panic(err)
	}
	fmt.Printf("原始issue内容: %s\n", string(content))

	group, friendShortcode, err := Generate(string(content), friendFilename)
	if err != nil {
		return
	}
	fmt.Printf("生成的group: %s\n", group)
	fmt.Printf("生成的shortcode: %s\n", friendShortcode)

	friendFile, err := os.OpenFile(friendFilename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	defer friendFile.Close()

	err = Combine(group, friendShortcode, friendFile)
	if err != nil {
		panic(err)
	}
	fmt.Printf("合并 %s 成功\n", friendFilename)

}

func isEmpty(s string) bool {
	return s == "None" || s == "_No response_"
}

func defaultIfEmpty(s, def string) string {
	if isEmpty(s) {
		return def
	}
	return s
}
