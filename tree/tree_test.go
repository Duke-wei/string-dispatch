package tree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func printChildren(n *Node, prefix string) {
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
	root := NewTree()
	root.AddRoute("wengwei", 1)
	root.AddRoute("wengwei2", 2)
	root.AddRoute("wengwei3", 3)
	root.AddRoute("wengwei4", 4)
	root.AddRoute("weiweng", 5)
	root.AddRoute("weiwang", 6)
	printChildren(root, "`-")
	qa.NotEqual(root, nil)
}

func Test_getValue(t *testing.T) {
	qa := assert.New(t)
	root := NewTree()
	root.AddRoute("wengwei", 1)
	root.AddRoute("wengweng", 2)
	h := root.GetValue("wengwei")
	qa.Equal(1, h)
	h = root.GetValue("wengweng")
	qa.Equal(2, h)
	h = root.GetValue("wengweng2")
	qa.Equal(2, h)
	h = root.GetValue("handler")
	qa.Equal(nil, h)
}
