package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
	"text/template"
)

func main() {
	var inputFilename, outputFilename string

	flag.StringVar(&inputFilename, "in", "input.md", "inputFilename")
	flag.StringVar(&outputFilename, "out", "output.md", "outputFilename")
	flag.Parse()

	content, err := os.ReadFile(inputFilename)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))

	res := extract(string(content))

	for i, v := range res {
		// 去空白
		res[i] = strings.TrimSpace(v)
	}

	name := res[0]
	url := res[1]
	word := res[2]
	logo := res[3]
	group := res[4]
	primaryColor := res[5]
	imgAnimation := res[6]
	borderAnimation := res[7]

	if group == "None" || group == "_No response_" {
		group = "同行的伙伴们"
	}
	if primaryColor == "None" || primaryColor == "_No response_" {
		primaryColor = "default"
	}
	if imgAnimation == "None" || imgAnimation == "_No response_" {
		imgAnimation = "rotate"
	}
	if borderAnimation == "None" || borderAnimation == "_No response_" {
		borderAnimation = "shadow"
	}

	type Friend struct {
		Name            string
		Url             string
		Word            string
		Logo            string
		Group           string
		PrimaryColor    string
		ImgAnimation    string
		BorderAnimation string
	}

	f := Friend{
		Name:            name,
		Url:             url,
		Word:            word,
		Logo:            logo,
		Group:           group,
		PrimaryColor:    primaryColor,
		ImgAnimation:    imgAnimation,
		BorderAnimation: borderAnimation,
	}

	// 根据 friend.tpl 渲染
	tpl, err := template.New("friend.tpl").Delims("[[", "]]").ParseFiles("friend.tpl")
	if err != nil {
		panic(err)
	}

	out, _ := os.OpenFile(outputFilename, os.O_CREATE|os.O_WRONLY, 0666)
	defer out.Close()

	err = tpl.Execute(out, f)
	if err != nil {
		panic(err)
	}

}

// 提取出 msg 内的 ### 中间的内容
func extract(s string) []string {
	reg := regexp.MustCompile(`### 站点名字 (.*) ### 站点地址 (.*) ### 站点描述 (.*) ### 站点图标或个人头像 (.*) ### 友链分组（可自定义）(.*) ### 主题色 (.*) ### 头像动画 (.*) ### 边框动画 (.*) ### 还有什么其他想说的吗❤️ (.*)`)
	matches := reg.FindAllStringSubmatch(s, -1)
	if len(matches) == 0 {
		panic("no matches")
	}
	return matches[0][1:]
}
