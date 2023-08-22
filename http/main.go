package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// type circle struct {
// 	radius float64
// }

// type rect struct {
// 	width  float64
// 	height float64
// }

// func (r rect) area() float64 {
// 	return r.width * r.height
// }

// func (c circle) area() float64 {
// 	return math.Pi * c.radius * c.radius
// }

// type shape interface {
// 	area() float64
// }

type logWriter struct{}

func main() {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	lw := logWriter{}
	// io.Copy(os.Stdout, resp.Body)
	io.Copy(lw, resp.Body)

	// c1 := circle{4.5}
	// r1 := rect{4, 7}

	// shapes := []shape{c1, r1}

	// for _, sh := range shapes {
	// 	fmt.Println(sh.area())
	// }

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("The size of the bs is:", len(bs))

	return len(bs), nil
}
