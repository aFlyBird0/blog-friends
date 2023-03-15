package main

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
	"time"

	"github.com/parnurzeal/gorequest"
	"go.uber.org/multierr"
)

// Generate 从 issue 和模板文件中生成友链
func Generate(issueContent, tplFile string) (group, friendShortcode string, err error) {
	res, err := Extract(issueContent)
	if err != nil {
		return "", "", err
	}

	name := res[0]
	url, err := handleURLProtocol(res[1])
	if err != nil {
		return "", "", err
	}
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

func handleURLProtocol(url string) (string, error) {
	req := gorequest.New().Timeout(10 * time.Second)

	// if url has protocol, try it
	if strings.HasPrefix(url, "http") {
		fmt.Println("has protocol")
		_, _, errs := req.Get(url).End()
		fmt.Println(errs)
		return url, multierr.Combine(errs...)
	}

	// no protocol, try https first
	const HTTPS = "https://"
	_, _, err := req.Get(HTTPS + url).End()
	if err == nil {
		return HTTPS + url, nil
	}

	// try http
	const HTTP = "http://"
	_, _, errs := req.Get(HTTP + url).End()
	return HTTP + url, multierr.Combine(errs...)
}
