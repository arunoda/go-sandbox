package append

import (
	"github.com/edsrzf/mmap-go"
	"os"
	"testing"
)

func TestGenFile(t *testing.T) {
	f, err := os.OpenFile("data_mmap.txt", os.O_RDWR, 0644)
	// has no file
	if err != nil {
		f, err = os.OpenFile("data_mmap.txt", os.O_RDWR|os.O_CREATE, 0644)
		check(err)
		var size int = 1024 * 1024

		for i := 0; i < size; i++ {
			b := make([]byte, 1024)
			offset := i * 1024
			f.WriteAt(b, int64(offset))
		}
		f.Close()
	}

	f, err = os.OpenFile("data.txt", os.O_RDWR|os.O_CREATE, 0644)
	f.Close()
	check(err)
}

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
			if fileData != nil {
				fileData.Unmap()
			}

			position = 0
			mapd, err := mmap.MapRegion(f, length, mmap.RDWR, 0, offset)
			check(err)

			fileData = mapd
			offset += int64(length)
		}

		position = AppendMmap(fileData, name, position)
	}
}
