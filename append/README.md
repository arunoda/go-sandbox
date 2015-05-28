## Append Speed Test (MMP vs File.*)

~~~
PASS
BenchmarkAppendFile  2000000           879 ns/op
BenchmarkAppendMmap 20000000            58.8 ns/op
ok      _/data/sandbox/go-sandbox/append    3.894s
~~~

Mmapping a block and writing to that is faster than a simple file append.
