package append

import (
	"github.com/edsrzf/mmap-go"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func AppendFile(f os.File, name string, position int64) int64 {
	nameBytes := []byte(name)
	_, err := f.WriteAt(nameBytes, position)
	check(err)
	return position + int64(len(nameBytes))
}

func AppendMmap(fileData mmap.MMap, name string, position int64) int64 {
	nameBytes := []byte(name)
	for lc, val := range nameBytes {
		fileData[position+int64(lc)] = val
	}
	return position + int64(len(nameBytes))
}
