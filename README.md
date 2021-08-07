# Benchmarking string concatenation in Go

Experiment to benchmark efficient string concatenation in Go

* `fmt.Sprintf`
* `strings.Builder{}`
* `strings.Join`
* `+=`

## Results

```bash
$ go test -benchmem -bench=.
goos: darwin
goarch: amd64
pkg: github.com/caitlinelfring/go-strings-benchmark
cpu: Intel(R) Core(TM) i7-4870HQ CPU @ 2.50GHz
BenchmarkBuilder-8          	 7667181	       150.0 ns/op	     120 B/op	       4 allocs/op
BenchmarkBuilderWithCap-8   	23462280	        49.25 ns/op	      48 B/op	       1 allocs/op
BenchmarkPlusEqual-8        	 3133431	       375.4 ns/op	     264 B/op	       8 allocs/op
BenchmarkFmtSprint-8        	 3582502	       336.7 ns/op	      48 B/op	       1 allocs/op
BenchmarkJoin-8             	 9333879	       126.7 ns/op	      48 B/op	       1 allocs/op
PASS
ok  	github.com/caitlinelfring/go-strings-benchmark	7.112s
```
