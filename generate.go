package main

import (
	"bytes"
	"text/template"
)

// Generate 从 issue 和模板文件中生成友链
func Generate(issueContent, tplFile string) (group, friendShortcode string, err error) {
	res, err := Extract(issueContent)
	if err != nil {
		return "", "", err
	}

	name := res[0]
	url := res[1]
	word := res[2]
	logo := res[3]
	group = defaultIfEmpty(res[4], "同行的伙伴们")
	primaryColor := defaultIfEmpty(res[5], "default")
	imgAnimation := defaultIfEmpty(res[6], "rotate")
	borderAnimation := defaultIfEmpty(res[7], "shadow")

	friend := map[string]string{
		"name":            name,
		"url":             url,
		"word":            word,
		"logo":            logo,
		"group":           group,
		"primaryColor":    primaryColor,
		"imgAnimation":    imgAnimation,
		"borderAnimation": borderAnimation,
	}

	friendShortcode, err = render(tplFile, friend)

	return group, friendShortcode, err
}

// 从模板文件和数据中生成友链shortcode
func render(tplFile string, data map[string]string) (string, error) {
	tpl, err := template.New(tplFile).Delims("[[", "]]").ParseFiles(tplFile)
	if err != nil {
		return "", err
	}

	var out bytes.Buffer
	err = tpl.Execute(&out, data)
	if err != nil {
		return "", err
	}

	return out.String(), nil
}
