package examples

import (
	"fmt"
	"github.com/tsouza/go-bus"
	"testing"
)

func TestWildcard_1(t *testing.T) {
	b := bus.New()
	sub := b.Subscribe([]string{"my", "*"}, func(p interface{}) {
		fmt.Printf("%v world\n", p)
	})
	b.Publish([]string{"my", "topic"}, "hello")
	sub.Terminate()
}