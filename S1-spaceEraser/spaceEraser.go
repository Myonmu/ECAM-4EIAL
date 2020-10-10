package main

import (
	"io"
	"os"
	"strings"
)

type spaceEraser struct {
	r io.Reader
}

func (se spaceEraser) Read(target []byte) (int, error) {
	n, err := se.r.Read(target)
	if err != nil {
		return n, err
	}
	index := 0
	for i := 0; i < n; i++ {
		if target[i] != 32 {
			target[index] = target[i]
			index++
		}
	}
	return index, err
}

func main() {
	s := strings.NewReader("H e l l o w o r l d!")
	r := spaceEraser{s}
	io.Copy(os.Stdout, &r)
}
