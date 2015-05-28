## Append Speed Test (MMP vs File.*)

~~~
PASS
BenchmarkAppendFile  2000000           879 ns/op
BenchmarkAppendMmap 20000000            58.8 ns/op
ok      _/data/sandbox/go-sandbox/append    3.894s
~~~

Mmapping a block and writing to that is faster than a simple file append.

## Setup

First we need to allocate create two files

1. data.txt - a simple txt file
2. data_mmap.txt - pre allocated file with `mkfile 1G data_mmap.txt`

