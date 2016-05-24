trylock - TryLock implementation for Go
=======================================

[![Build Status](https://travis-ci.org/DmitriyMV/trylock.svg?branch=master)](https://travis-ci.org/DmitriyMV/trylock)
[![GoDoc](https://godoc.org/github.com/DmitriyMV/trylock?status.svg)](https://godoc.org/github.com/DmitriyMV/trylock)
[![Coverage Status](https://img.shields.io/coveralls/DmitriyMV/trylock.svg?flat=1)](https://coveralls.io/github/DmitriyMV/trylock)

trylock uses unsafe, which is sorta "unsafe", but should work until `sync.Mutex`
will change its layout (I hope it never will).

# Usage

```go
type LockedStruct struct {
	mu trylock.Mutex
}

storage := &LockedStruct{}

if storage.mu.TryLock() {
	// do something with storage
} else {
	// return busy or use some logic for unavailable storage
}
```
