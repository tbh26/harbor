package code

import (
	"errors"
	"fmt"
)

func ReadersWritersDemo() {
	fmt.Println("Hello code.ReadersWritersDemo() world! (ch7)")

	ReaderDemo()
	WriterDemo()

	fmt.Println()
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type MyReader struct {
	data string
	from int
}

func (r *MyReader) Read(p []byte) (int, error) {
	if p == nil {
		return -1, errors.New("nil target slice")
	}
	if len(r.data) <= 0 || r.from == len(r.data) {
		return 0, errors.New("end of file") // io.EOF
	}
	n := len(r.data) - r.from
	if len(p) < n {
		n = len(p)
	}
	for i := 0; i < n; i += 1 {
		b := byte(r.data[r.from])
		p[i] = b
		r.from += 1
	}
	if r.from == len(r.data) {
		return n, errors.New("end of file")
	}
	return n, nil
}

func ReaderDemo() {
	fmt.Println("\n=-= ReaderDemo =-=")

	//
	target := make([]byte, 5)
	empty := MyReader{}
	n, err := empty.Read(target)
	fmt.Println(" -- empty reader demo --")
	fmt.Printf(" - read %d: error: %v \n", n, err)

	//
	fmt.Println(" -- my-reader demo --")
	mr := MyReader{"Save the world with Go!", 0}
	n, err = mr.Read(target)
	for err == nil {
		fmt.Printf(" - %d read, target: \"%s\", error: %v \n", n, target, err)
		n, err = mr.Read(target)
	}
	fmt.Printf(" - %d read, target: \"%s\", error: %v \n", n, target, err)
	fmt.Printf(" - %d read, target:  \"%s\" , error: %v \n", n, target[:n], err)

	fmt.Println()
}

type MyWriter struct {
	data string
	size int
}

func (w *MyWriter) Write(p []byte) (int, error) {
	if p == nil {
		return -1, errors.New("nil target slice")
	}
	if len(p) == 0 {
		return 0, errors.New("end of file")
	}
	n := w.size
	var err error = nil
	if len(p) < w.size {
		n = len(p)
	} else {
		err = errors.New(fmt.Sprintf("payload too large (%d max, got %d)", w.size, len(p)))
	}
	w.data = w.data + string(p[0:n])
	return n, err
}

func WriterDemo() {
	fmt.Println("\n=-= WriterDemo =-=")

	msg := []byte("the world with Go!!")

	stock := MyWriter{"Save ", 6}
	i := 0
	var err error
	for err == nil && i < len(msg) {
		n, err := stock.Write(msg[i:])
		fmt.Printf(" - %d written, stock: \"%s\", error; %v\n", n, stock.data, err)
		i = i + n
	}
	fmt.Println(" ---")
	stock = MyWriter{"Save ", 6}
	offset := 5
	ml := len(msg)
	for i = 0; i < ml; i += offset {
		sliceEnd := i + offset
		if ml < sliceEnd {
			sliceEnd = len(msg)
		}
		n, e := stock.Write(msg[i:sliceEnd])
		fmt.Printf(" - %d written, stock: \"%s\", slice-end: %d, error; %v  (i: %d, message-length: %d)\n", n, stock.data, sliceEnd, e, i, ml)
	}

	fmt.Println()
}
