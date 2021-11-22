## Dice - random string/int utility for the Go language

This package contains a random string and int generator. It is a wrapper around the crypto/rand package.

**NOTE**: This package is provided "as is" with no guarantee. Use it at your own risk and always test it yourself before using it in a production environment. If you find any issues, please [create a new issue](https://github.com/twiny/dice/issues/new).


### Usage
```
go get github.com/twiny/dice
```

```go
package main

import (
	"fmt"

	"github.com/twiny/dice"
)

func main() {
	for i := 0; i < 5; i++ {
		j := dice.RandInt(1, 234567890)
		fmt.Println(j)
	}
	//

	for i := 0; i < 5; i++ {
		j := dice.RandString(25)
		fmt.Println(j)
	}

	//
	fmt.Println("done :)")
}

// Output:
/*
	159191458
	40548335
	112436247
	211217306
	137384292
	EnZZq88MIR3sc4qUwJBHSkVmE
	QFYze2dCr1UP0czV62xzWX3eP
	UhIpkypBNe59LTZnQl9KkoxQk
	tyffLFt6y2B9TsZB0H0eoqH0C
	pCcAXLhqCOTN7w1htABSDEIfC
	done :)
*/
```


### Benchmark
```
goos: darwin
goarch: amd64
pkg: dice
cpu: Intel(R) Core(TM) i5-7360U CPU @ 2.30GHz
BenchmarkRandString-4   	 1260508	       945.0 ns/op	      64 B/op	       2 allocs/op
BenchmarkRandInt-4      	 1000000	      1205 ns/op	      56 B/op	       4 allocs/op
PASS
ok  	dice	3.493s
```