/* logrot.go

MIT License

Copyright (c) 2018 charlanxcc

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/

package logrot

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func parseSize(size string) (int, error) {
	m := 1
	size = strings.TrimSpace(size)
	switch size[len(size)-1:] {
	case "k": fallthrough
	case "K":
		m = 1024
		size = strings.TrimSpace(size[:len(size)-1])
	case "m": fallthrough
	case "M":
		m = 1024 * 1024
		size = strings.TrimSpace(size[:len(size)-1])
	case "g": fallthrough
	case "G":
		m = 1024 * 1024 * 1024
		size = strings.TrimSpace(size[:len(size)-1])
	}

	i, err := strconv.Atoi(size)
	if err != nil {
		return 0, err
	} else {
		return i * m, nil
	}
}

func rotate(name string, size, count int) error {
	fi, err := os.Stat(name)
	if err != nil {
		return nil
	}

	if fi.Size() < int64(size) {
		return nil
	}

	for i := count; i >= 0; i-- {
		var src, dst string

		if i == count {
			os.Remove(fmt.Sprintf("%s.%d", name, i))
			continue
		} else if i == 0 {
			src = name
		} else {
			src = fmt.Sprintf("%s.%d", name, i)
		}
		dst = fmt.Sprintf("%s.%d", name, i + 1)

		if _, err := os.Stat(src); err != nil && os.IsNotExist(err) {
			continue
		}

		err := os.Rename(src, dst)
		if err != nil {
			return err
		}
	}
	return nil
}

// r could be either os.Stdin or io.PipeReader from io.Pipe()
func LogRotate(r io.Reader, filename string, size, count int) error {
	if size <= 0 || count <= 0 {
		return fmt.Errorf("Invalid Parameters")
	}

	err := rotate(filename, size, count)
	if err != nil {
		return err
	}

	newline := []byte{'\n'}
	in := bufio.NewScanner(r)
	for {
		ix := 0

		out, err := os.OpenFile(filename, os.O_WRONLY | os.O_CREATE, 0600)
		if err != nil {
			return err
		}

		off, err := out.Seek(0, 2)
		if err != nil {
			return err
		}
		ix = int(off)

		closed := false
		for {
			if in.Scan() == false {
				closed = true
				break
			}

			l, err := out.Write(in.Bytes())
			if err != nil {
				out.Close()
				return err
			}
			ix += l
			l, err = out.Write(newline)
			if err != nil {
				out.Close()
				return err
			}

			if ix >= size {
				break
			}
		}

		out.Close()
		err = rotate(filename, size, count)
		if err != nil {
			return err
		}

		if closed {
			break
		}
	}

	return nil
}

func usage() {
	fmt.Printf("Usage: logrot <file-name> <size> <count>\n")
}

func Main() {
	if len(os.Args) != 4 {
		usage()
		os.Exit(0)
	}

	filename := os.Args[1]
	size, err1 := parseSize(os.Args[2])
	count, err2 := strconv.Atoi(os.Args[3])
	if err1 != nil || err2 != nil || size <= 0 || count <= 0 {
		usage()
		os.Exit(1)
	}

	err := LogRotate(os.Stdin, filename, size, count)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
}

/* EOF */
