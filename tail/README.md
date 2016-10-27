Package Tail
============
Implementation of a tail call optimization in go

How to use
----------
```go
// change the return type of your function to interface{}
func Factorial(n int, acc int64) interface{} {
	if n <= 0 {
		return acc
	}

	// add this wrapper to your tail call
	return tail.Flooper(func() interface{} {
		return Factorial(n-1, int64(n)*acc)
	})
}

func main() {
	tail.Loop(Factorial(10000, int64(1)))
}
```

### Of course you can make it pretty
```go
func Factorial(n int) int64 {
	return tail.Loop(factorial(n, int64(1))).(int64)
}

func factorial(n int, acc int64) interface{} {
	if n <= 0 {
		return acc
	}
	return tail.Flooper(func() interface{} {
		return factorial(n-1, int64(n)*acc)
	})
}

func main() {
	Factorial(10000)
}
```

### You could even hide the function
```go
func Factorial(n int) int64 {
	var factorial func(int,int64) interface{}
	factorial = func(n int, acc int64) interface{} {
		if n <= 0 {
			return acc
		}
		return tail.Flooper(func() interface{} {
			return factorial(n-1, int64(n)*acc)
		})
	}
	return tail.Loop(factorial(n, int64(1))).(int64)
}

func main() {
	Factorial(10000)
}
```
