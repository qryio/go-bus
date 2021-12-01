package examples

import (
	"fmt"
	"github.com/qryio/go-bus"
	"testing"
)

func TestSimple_1(t *testing.T) {
	b := bus.New()
	_ = b.Subscribe([]string{"my", "topic"}, func(t []string, s *bus.Subscription, p interface{}) {
		fmt.Printf("%v world\n", p)
		go s.Terminate()
	})
	b.Publish([]string{"my", "topic"}, "hello")
}