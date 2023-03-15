package main

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestBlogFriends(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BlogFriends Suite")
}
