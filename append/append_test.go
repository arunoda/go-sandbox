package append

import (
	"os"
	"syscall"
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

	var fileData []byte
	var file *os.File
	var offset int64 = 0
	var length int = syscall.Getpagesize() * 10

	var position int64 = 0
	for n := 0; n < b.N; n++ {
		name := "Hello\n"
		if position == 0 || position+int64(len(name)) >= int64(length) {
			if fileData != nil {
				syscall.Munmap(fileData)
				file.Close()
			}

			position = 0

			f, err := os.OpenFile("data_mmap.txt", os.O_RDWR, 0644)
			check(err)

			mapd, err := syscall.Mmap(int(f.Fd()), offset, length, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
			check(err)

			fileData = mapd
			file = f
			offset += int64(length)
		}

		position = AppendMmap(fileData, name, position)
	}
}
