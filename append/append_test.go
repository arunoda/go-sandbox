package append

import (
	"github.com/edsrzf/mmap-go"
	"os"
	"testing"
)

func BenchmarkAppendFile(b *testing.B) {
	f, err := os.OpenFile("data.txt", os.O_RDWR, 0644)
	check(err)

	var position int64 = 0
	for n := 0; n < b.N; n++ {
		position = AppendFile(*f, "Hello\n", position)
	}
}

func BenchmarkAppendMmap(b *testing.B) {
	f, err := os.OpenFile("data_mmap.txt", os.O_RDWR, 0644)
	check(err)

	var fileData mmap.MMap
	var offset int64 = 0
	var length int = 10000

	var position int64 = 0
	for n := 0; n < b.N; n++ {
		name := "Hello\n"
		if position == 0 || position+int64(len(name)) >= int64(length) {
			position = 0
			mapd, err := mmap.MapRegion(f, length, mmap.RDWR, 0, offset)
			check(err)

			if fileData != nil {
				fileData.Unmap()
			}

			fileData = mapd
			offset += int64(length)
		}

		position = AppendMmap(fileData, name, position)
	}
}
