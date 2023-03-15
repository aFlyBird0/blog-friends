package main

import (
	"fmt"
	"regexp"
	"strings"
)

// Extract 提取出 msg 内的 ### 中间的内容，并去空白
func Extract(s string) ([]string, error) {
	const tpl = `### 站点名字 (.*) ### 站点地址 (.*) ### 站点描述 (.*) ### 站点图标或个人头像 (.*) ### 友链分组（可自定义）(.*) ### 主题色 (.*) ### 头像动画 (.*) ### 边框动画 (.*) ### 还有什么其他想说的吗❤️ (.*)`
	reg := regexp.MustCompile(tpl)
	matches := reg.FindAllStringSubmatch(s, -1)
	if len(matches) == 0 {
		return nil, fmt.Errorf("no matches found")
	}

	res := matches[0][1:]
	// 去空白
	for i, v := range res {
		res[i] = strings.TrimSpace(v)
	}
	return res, nil
}
