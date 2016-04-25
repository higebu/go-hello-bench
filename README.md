# go-hello-bench

This is an repository of golang benchmark example.

Why Hello function called too many times...

* c1: the BenchmarkHello function called count
* c2: the Hello function called count

Result:

```sh
$ go test -bench .
2016/04/25 20:18:24 NumCPU: 4
testing: warning: no tests to run
PASS
BenchmarkHello-4    2016/04/25 20:18:24 c1: 0
2016/04/25 20:18:24 c2: 1
2016/04/25 20:18:24 c1: 1
2016/04/25 20:18:24 c2: 101
2016/04/25 20:18:24 c1: 2
2016/04/25 20:18:24 c2: 10101
2016/04/25 20:18:24 c1: 3
2016/04/25 20:18:25 c2: 1010101
2016/04/25 20:18:25 c1: 4
2016/04/25 20:18:27 c2: 6010101
 5000000           481 ns/op
 ok     github.com/higebu/go-hello-bench    2.778s
```
