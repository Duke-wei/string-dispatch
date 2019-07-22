package tree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func printChildren(n *node, prefix string) {
	fmt.Printf(" %02d %s%s[%d] %v \r\n", n.priority, prefix, n.path, len(n.children), n.handle)
	for l := len(n.path); l > 0; l-- {
		prefix += "-"
	}
	for _, child := range n.children {
		printChildren(child, prefix)
	}
}

func Test_min(t *testing.T) {
	qa := assert.New(t)
	qa.True(true, min(3, 4), 3)
	qa.False(false, min(3, 4), 4)
}

func Test_addRoute(t *testing.T) {
	qa := assert.New(t)
	root := &node{}
	root.addRoute("wengwei", 1)
	root.addRoute("wengwei2", 2)
	root.addRoute("wengwei3", 3)
	root.addRoute("wengwei4", 4)
	root.addRoute("weiweng", 5)
	root.addRoute("weiwang", 6)
	printChildren(root,"`-")
	qa.NotEqual(root, nil)
}
