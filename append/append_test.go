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
	check(err)
	f.Close()
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

	fileData, err := mmap.Map(f, mmap.RDWR, 0)
	check(err)

	position := int64(0)

	for n := 0; n < b.N; n++ {
		name := "Hello\n"
		position = AppendMmap(fileData, name, position)
	}
}
