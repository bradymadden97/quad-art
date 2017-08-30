// main.go
package main

import (
	"container/heap"
	"log"
)

type Img struct {
	hist   [][]int   //Histogram of image stored as [R, G, B]
	pix    int       //Number of pixels in image
	color  []float64 //Average color stored as [R, G, B]
	error  float64   //Calculated error between average pixels and image
	width  int       //Picture width
	height int       //Picture height
	c1     *Img      //Pointer to child 1
	c2     *Img      //Pointer to child 2
	c3     *Img      //Pointer to child 3
	c4     *Img      //Pointer to child 4
}

func main() {
	n, i, g, gf, gp, b := initializeFlags()
	if *n == "" {
		log.Fatal(" -f <input image> required")
	}

	headNode, err := initialize(*n)
	if err != nil {
		log.Fatal(err)
	}

	mh := make(MinHeap, 1)
	mh[0] = headNode
	heap.Init(&mh)

	err = iterate(&mh, headNode, *i, *n, *b)
	if err != nil {
		log.Fatal(err)
	}

	if *g {
		err = toGIF(*n, *gf, *gp)
		if err != nil {
			log.Fatal(err)
		}
	}
}
