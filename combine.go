package main

import (
	"bytes"
	"fmt"
	"io"
)

// Combine 把 shortcode 和 已有的 friend.md 文件合并
func Combine(group, shortcode string, file io.ReadWriter) error {
	origin, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	groupWithHeader := fmt.Sprintf("# %s\n", group)
	groupIndexBegin := bytes.Index(origin, []byte(groupWithHeader))

	if groupIndexBegin == -1 {
		// 如果 group 不存在，就在文件末尾插入 group 和 shortcode
		origin = append(origin, []byte(groupWithHeader)...)
		origin = append(origin, []byte(shortcode)...)
	} else {
		// 找到含有 group 的行，并且在下一个 # 开头的行之前插入 group 和 shortcode

		// 先按这个 group 分成两段
		originFirstSection := origin[:groupIndexBegin]
		originSecondSection := origin[groupIndexBegin:]

		// 找到第二段中第一个 # 开头的行
		groupIndexEnd := bytes.Index(originSecondSection, []byte("\n# "))
		if groupIndexEnd == -1 {
			// 如果没有找到，就说明这个 group 是最后一个 group，直接在末尾插入 shortcode
			originSecondSection = append(originSecondSection, []byte(shortcode)...)
		} else {
			// 如果找到了，就在这个 group 之后插入 shortcode
			originSecondSection = append(originSecondSection[:groupIndexEnd], append([]byte(shortcode), originSecondSection[groupIndexEnd:]...)...)
		}

		// 把两段合并起来
		origin = append(originFirstSection, originSecondSection...)
	}

	_, err = file.Write(origin)

	return err
}
