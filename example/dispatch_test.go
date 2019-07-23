package example

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type Member struct {
	Name         string
	MonkeyPerMin string
	WorkMin      string
}

var Staff = []Member{
	{"work1", "12.5", "60"},
	{"work2", "13.5", "67"},
	{"manager1", "30.5", "70"},
	{"manager2", "28.5", "75"}}

var D = DispatchTree()
var W = &Worker{}
var M = &Manager{}

func TestDispatchTree(t *testing.T) {
	sum := 0.0
	for i := 0; i < len(Staff); i++ {
		h := D.GetValue(Staff[i].Name).(Calculate)
		sum += h.Cal(Staff[i].MonkeyPerMin, Staff[i].WorkMin)
	}
	qa := assert.New(t)
	qa.Equal(3748.9, sum)
}

func BenchmarkDispatchTree(b *testing.B) {
	b.StopTimer()
	b.StartTimer()
	for j := 0; j < b.N; j++ {
		sum := 0.0
		for i := 0; i < len(Staff); i++ {
			h := D.GetValue(Staff[i].Name).(Calculate)
			sum += h.Cal(Staff[i].MonkeyPerMin, Staff[i].WorkMin)
		}
	}
}

func BenchmarkDispatchTree2(b *testing.B) {
	b.StopTimer()
	b.StartTimer()
	for j := 0; j < b.N; j++ {
		sum := 0.0
		for i := 0; i < len(Staff); i++ {
			if strings.Contains(Staff[i].Name, "work") {
				sum += W.Cal(Staff[i].MonkeyPerMin, Staff[i].WorkMin)
			} else if strings.Contains(Staff[i].Name, "manager") {
				sum += M.Cal(Staff[i].MonkeyPerMin, Staff[i].WorkMin)
			}
		}
	}
}
