package examples

import (
	"fmt"
	"github.com/tsouza/go-bus"
	"testing"
)

func TestSimple_1(t *testing.T) {
	b := bus.New()
	sub := b.Subscribe([]string{"my", "topic"}, func(t []string, p interface{}) {
		fmt.Printf("%v world\n", p)
	})
	b.Publish([]string{"my", "topic"}, "hello")
	sub.Terminate()
}