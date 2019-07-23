// Copyright 2013 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

//source from https://github.com/julienschmidt/httprouter/blob/master/tree.go

package tree

//change this type for your own api
type Handle interface{}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Node struct {
	path     string
	indices  string
	children []*Node
	handle   Handle
	priority uint32
}

func NewTree() *Node {
	return &Node{}
}

// increments priority of the given child and reorders if necessary
func (n *Node) incrementChildPrio(pos int) int {
	n.children[pos].priority++
	prio := n.children[pos].priority

	// adjust position (move to front)
	newPos := pos
	for newPos > 0 && n.children[newPos-1].priority < prio {
		// swap node positions
		n.children[newPos-1], n.children[newPos] = n.children[newPos], n.children[newPos-1]

		newPos--
	}

	// build new index char string
	if newPos != pos {
		n.indices = n.indices[:newPos] + // unchanged prefix, might be empty
			n.indices[pos:pos+1] + // the index char we move
			n.indices[newPos:pos] + n.indices[pos+1:] // rest without char at 'pos'
	}

	return newPos
}

// addRoute adds a node with the given handle to the path.
// Not concurrency-safe!
func (n *Node) AddRoute(path string, handle Handle) {
	fullPath := path
	n.priority++

	// non-empty tree
	if len(n.path) > 0 || len(n.children) > 0 {
	walk:
		for {
			// Find the longest common prefix.
			i := 0
			max := min(len(path), len(n.path))
			for i < max && path[i] == n.path[i] {
				i++
			}

			// Split edge
			if i < len(n.path) {
				child := Node{
					path:     n.path[i:],
					indices:  n.indices,
					children: n.children,
					handle:   n.handle,
					priority: n.priority - 1,
				}

				n.children = []*Node{&child}
				// []byte for proper unicode char conversion, see #65
				n.indices = string([]byte{n.path[i]})
				n.path = path[:i]
				n.handle = nil
			}

			// Make new node a child of this node
			if i < len(path) {
				path = path[i:]
				c := path[0]

				// Check if a child with the next path byte exists
				for i := 0; i < len(n.indices); i++ {
					if c == n.indices[i] {
						i = n.incrementChildPrio(i)
						n = n.children[i]
						continue walk
					}
				}
				n.indices += string([]byte{c})
				child := &Node{}
				n.children = append(n.children, child)
				n.incrementChildPrio(len(n.indices) - 1)
				n = child
				n.path = path[:]
				n.handle = handle
				return

			} else if i == len(path) { // Make node a (in-path) leaf
				if n.handle != nil {
					panic("a handle is already registered for path '" + fullPath + "'")
				}
				n.handle = handle
			}
			return
		}
	} else { // Empty tree
		n.path = path[:]
		n.handle = handle
	}
}

// Returns the handle registered with the given path (key). The values of
// wildcards are saved to a map.
func (n *Node) GetValue(path string) Handle {
walk: // outer loop for walking the tree
	for {
		if len(path) > len(n.path) {
			if path[:len(n.path)] == n.path {
				path = path[len(n.path):]
				c := path[0]
				for i := 0; i < len(n.indices); i++ {
					if c == n.indices[i] {
						n = n.children[i]
						continue walk
					}
				}
				//can not find in indices
				//return n handle
				return n.handle
			}
		} else if path == n.path {
			return n.handle
		}
		return nil
	}
}
