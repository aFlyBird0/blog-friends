package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const inputForTest = `### 站点名字 测试站点 ### 站点地址 blog.aflybird.cn ### 站点描述 这是站点描述 ### 站点图标或个人头像 https://example.com/avatar.png ### 友链分组（可自定义） _No response_ ### 主题色 None ### 头像动画 None ### 边框动画 None ### 还有什么其他想说的吗❤️ _No response_`
const inputForTestWithCustomGroup = `### 站点名字 测试站点 ### 站点地址 blog.aflybird.cn ### 站点描述 这是站点描述 ### 站点图标或个人头像 https://example.com/avatar.png ### 友链分组（可自定义） 这是自定义分组 ### 主题色 None ### 头像动画 None ### 边框动画 None ### 还有什么其他想说的吗❤️ _No response_`
const customGroupForTest = "这是自定义分组"

var _ = Describe("Extract", func() {

	When("input is valid and group is default", func() {
		It("should return valid result", func() {
			res, err := Extract(inputForTest)
			Expect(err).To(BeNil())
			Expect(res).To(Equal([]string{
				"测试站点",
				"blog.aflybird.cn",
				"这是站点描述",
				"https://example.com/avatar.png",
				"_No response_",
				"None",
				"None",
				"None",
				"_No response_",
			}))
		})
	})

	When("input is valid and group is custom", func() {
		It("should return valid result", func() {
			res, err := Extract(inputForTestWithCustomGroup)
			Expect(err).To(BeNil())
			Expect(res).To(Equal([]string{
				"测试站点",
				"blog.aflybird.cn",
				"这是站点描述",
				"https://example.com/avatar.png",
				customGroupForTest,
				"None",
				"None",
				"None",
				"_No response_",
			}))
		})
	})
})
