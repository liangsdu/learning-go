// image包 

package main

import (
	"fmt"
	"image"
)

func main() {
	m := image.NewRGBA(100,100)
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0,0).RGBA())
}
