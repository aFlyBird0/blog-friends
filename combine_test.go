package main

import (
	"bytes"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const shortCodeForTest = `
{{< friend
name="测试站点"
url="http://test.com"
logo="https://example.com/avatar.png"
word="这是站点描述"
primary-color="default"
img-animation="rotate"
border-animation="shadow"
>}}
`

var _ = Describe("Combine", func() {
	const originFriendContent = `
---
title: "友链墙"
url: friends
hiddenFromHomePage: true
---

{{< friend-css >}}

> 提交友链方法：先将本站加入友链，然后在前往[此表单](https://github.com/aFlyBird0/blog-friends/issues/new?assignees=aFlyBird0&labels=&template=apply-friend-link.yml&title=%E7%94%B3%E8%AF%B7%E5%8F%8B%E9%93%BE:+)填写对应信息，提交即可。

# 博客相关
感谢 Hugo 博客框架，LoveIt 主题，以及 雨临Lewis 的友链教程。

{{< friend
name="Hugo"
url="https://gohugo.io/"
logo="https://d33wubrfki0l68.cloudfront.net/c38c7334cc3f23585738e40334284fddcaf03d5e/2e17c/images/hugo-logo-wide.svg"
word="The world’s fastest framework for building websites"
>}}

{{< friend
name="LoveIt"
url="https://github.com/dillonzq/LoveIt"
logo="https://hugoloveit.com/images/avatar.png"
word="❤️A clean, elegant but advanced blog theme for Hugo"
>}}


{{< friend
name="雨临Lewis的博客"
url="https://lewky.cn"
logo="https://lewky.cn/images/avatar.jpg"
word="不想当写手的码农不是好咸鱼_(xз」∠)_"
>}}

# 同行的伙伴们

{{< friend
name="iyear"
url="https://blog.ljyngup.com/"
logo="https://blog.ljyngup.com/usr/uploads/2021/09/2354485274.jpg"
word="记录本是反抗"
>}}

{{< friend
name="Forever Young"
url="http://erdengk.top/"
logo="https://avatars.githubusercontent.com/u/37730787?v=4"
word="生活 开源 技术"
>}}
`
	const (
		groupExistMid = `博客相关`
		groupNotExist = `不存在的分组`
		groupExistEnd = `同行的伙伴们`
	)

	var (
		group         string
		finalContent  string
		expectContent string
		buffer        bytes.Buffer
	)

	BeforeEach(func() {
		buffer.WriteString(originFriendContent)
	})

	JustBeforeEach(func() {
		err := Combine(group, shortCodeForTest, &buffer)
		Expect(err).ShouldNot(HaveOccurred())
		finalContent = buffer.String()
		fmt.Println(finalContent)
		fmt.Println(expectContent)
		Expect(finalContent).To(Equal(expectContent))
	})

	JustAfterEach(func() {
		buffer.Reset()
	})

	When("group is not exist", func() {
		BeforeEach(func() {
			group = groupNotExist
			expectContent = originFriendContent + "# " + groupNotExist + "\n" + shortCodeForTest
		})

		It("should combine successfully", func() {})

	})

	When("group is exist and in the middle", func() {
		BeforeEach(func() {
			group = groupExistMid
			expectContent = `
---
title: "友链墙"
url: friends
hiddenFromHomePage: true
---

{{< friend-css >}}

> 提交友链方法：先将本站加入友链，然后在前往[此表单](https://github.com/aFlyBird0/blog-friends/issues/new?assignees=aFlyBird0&labels=&template=apply-friend-link.yml&title=%E7%94%B3%E8%AF%B7%E5%8F%8B%E9%93%BE:+)填写对应信息，提交即可。

# 博客相关
感谢 Hugo 博客框架，LoveIt 主题，以及 雨临Lewis 的友链教程。

{{< friend
name="Hugo"
url="https://gohugo.io/"
logo="https://d33wubrfki0l68.cloudfront.net/c38c7334cc3f23585738e40334284fddcaf03d5e/2e17c/images/hugo-logo-wide.svg"
word="The world’s fastest framework for building websites"
>}}

{{< friend
name="LoveIt"
url="https://github.com/dillonzq/LoveIt"
logo="https://hugoloveit.com/images/avatar.png"
word="❤️A clean, elegant but advanced blog theme for Hugo"
>}}


{{< friend
name="雨临Lewis的博客"
url="https://lewky.cn"
logo="https://lewky.cn/images/avatar.jpg"
word="不想当写手的码农不是好咸鱼_(xз」∠)_"
>}}

{{< friend
name="测试站点"
url="http://test.com"
logo="https://example.com/avatar.png"
word="这是站点描述"
primary-color="default"
img-animation="rotate"
border-animation="shadow"
>}}

# 同行的伙伴们

{{< friend
name="iyear"
url="https://blog.ljyngup.com/"
logo="https://blog.ljyngup.com/usr/uploads/2021/09/2354485274.jpg"
word="记录本是反抗"
>}}

{{< friend
name="Forever Young"
url="http://erdengk.top/"
logo="https://avatars.githubusercontent.com/u/37730787?v=4"
word="生活 开源 技术"
>}}
`
		})
		It("should combine successfully", func() {})

	})

	When("group is exist and in the end", func() {
		BeforeEach(func() {
			group = groupExistEnd
			expectContent = `
---
title: "友链墙"
url: friends
hiddenFromHomePage: true
---

{{< friend-css >}}

> 提交友链方法：先将本站加入友链，然后在前往[此表单](https://github.com/aFlyBird0/blog-friends/issues/new?assignees=aFlyBird0&labels=&template=apply-friend-link.yml&title=%E7%94%B3%E8%AF%B7%E5%8F%8B%E9%93%BE:+)填写对应信息，提交即可。

# 博客相关
感谢 Hugo 博客框架，LoveIt 主题，以及 雨临Lewis 的友链教程。

{{< friend
name="Hugo"
url="https://gohugo.io/"
logo="https://d33wubrfki0l68.cloudfront.net/c38c7334cc3f23585738e40334284fddcaf03d5e/2e17c/images/hugo-logo-wide.svg"
word="The world’s fastest framework for building websites"
>}}

{{< friend
name="LoveIt"
url="https://github.com/dillonzq/LoveIt"
logo="https://hugoloveit.com/images/avatar.png"
word="❤️A clean, elegant but advanced blog theme for Hugo"
>}}


{{< friend
name="雨临Lewis的博客"
url="https://lewky.cn"
logo="https://lewky.cn/images/avatar.jpg"
word="不想当写手的码农不是好咸鱼_(xз」∠)_"
>}}

# 同行的伙伴们

{{< friend
name="iyear"
url="https://blog.ljyngup.com/"
logo="https://blog.ljyngup.com/usr/uploads/2021/09/2354485274.jpg"
word="记录本是反抗"
>}}

{{< friend
name="Forever Young"
url="http://erdengk.top/"
logo="https://avatars.githubusercontent.com/u/37730787?v=4"
word="生活 开源 技术"
>}}

{{< friend
name="测试站点"
url="http://test.com"
logo="https://example.com/avatar.png"
word="这是站点描述"
primary-color="default"
img-animation="rotate"
border-animation="shadow"
>}}
`

		})

		It("should combine successfully", func() {})

	})
})
