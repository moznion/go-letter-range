go-letter-range
==

[![Build Status](https://travis-ci.org/moznion/go-letter-range.svg?branch=master)](https://travis-ci.org/moznion/go-letter-range) [![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](http://godoc.org/github.com/moznion/go-letter-range/letterrange)

A library that lists letters according to specified letters range.

Example
--

```go
package main

import (
	"fmt"

	"github.com/moznion/go-letter-range/letterrange"
)

func main() {
	got, err := letterrange.GetRunesWithinRange('a', 'z')
	if err != nil {
		// do error handling
		return
	}
	fmt.Println(got) // out: a rune slice that contains 'a' to 'z' runes
}
```

```go
package main

import (
	"fmt"

	"github.com/moznion/go-letter-range/letterrange"
)

func main() {
	got, err := GetRunesWithinRangeString("u{0061}-u{007A}")
	if err != nil {
		// do error handling
		return
	}
	fmt.Println(got) // out: a rune slice that contains 'a' to 'z' runes
}
```

Docs
--

Please see: [![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](http://godoc.org/github.com/moznion/go-letter-range/letterrange)

License
--

```
The MIT License (MIT)
Copyright © 2018 moznion, http://moznion.net/ <moznion@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the “Software”), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
```

