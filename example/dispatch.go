package example

import (
	"github.com/Duke-wei/string-dispatch/tree"
	"strconv"
)

type Calculate interface {
	Cal(string, string) float64
}

type Manager struct {

}

func (m *Manager) Cal(s1 string,s2 string) float64{
	f1,_ := strconv.ParseFloat(s1,64)
	f2,_ := strconv.ParseFloat(s2,64)
	return 0.8*f1*f2
}

type Worker struct {

}

func (w *Worker) Cal(s1 string, s2 string) float64{
	f1,_ := strconv.ParseFloat(s1,64)
	f2,_ := strconv.ParseFloat(s2,64)
	return 0.2*f1*f2
}

func DispatchTree() *tree.Node{
	root := tree.NewTree()
	root.AddRoute("work",&Worker{})
	root.AddRoute("manager",&Manager{})
	return root
}
