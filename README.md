# go-bus

![test](https://github.com/tsouza/go-bus/actions/workflows/test/badge.svg)

`go-bus` is a lightweight golang library for implementing the [pub/sub pattern](https://en.wikipedia.org/wiki/Publish%E2%80%93subscribe_pattern). It supports wildcard topics as well as dynamic topic subscription.

## Usage

Basic pub/sub usage:

```go
package main

import (
	"fmt"
	"github.com/tsouza/go-bus"
)

func main() {
	b := bus.New()
	sub := b.Subscribe([]string{"my", "topic"}, func(p interface{}) {
		fmt.Printf("%v world", p)
	})
	b.Publish([]string{"my", "topic"}, "hello")
	sub.Terminate()
}
```

Check out more [examples](examples)

## License

Code and documentation released under [Apache License 2.0](LICENSE)