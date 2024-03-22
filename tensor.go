package main

import (
	"fmt"
	"github.com/kshedden/gonpy"
)

type Tensor struct {
	shape []int
	arr   []float32
}

func (t *Tensor) at(coord ...int) float32 {
	ii := 0
	tempm := 1
	for i := len(t.shape) - 1; i >= 0; i-- {
		temp := coord[i]

		if i < len(t.shape)-1 {
			tempm *= t.shape[i+1]
			temp = tempm * coord[i]
		}
		ii += temp

	}
	// fmt.Println(ii)
	// fmt.Println(t.arr[ii])
	return t.arr[ii]
}

func (t *Tensor) readNpy(f string) {

	r, err := gonpy.NewFileReader(f)
	fmt.Println(err)
	fmt.Printf("%T\n", r)
	fmt.Println(r.Shape)
	fmt.Println(r.Dtype)
	data, err1 := r.GetFloat32()
	fmt.Println(err1)
	fmt.Println(data)
}
