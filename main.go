package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type tensor [][][][]float32

func main() {

	const h = 20
	const w = 30
	const c = 32

	// inStrArr := readFile("test/inpy.txt")
	// wgtStrArr := readFile("test/wgtpy.txt")
	// outStrArr := readFile("test/outpy.txt")

	// bias := make([]float32, 32)
	// inArr := strArrToFlat(inStrArr, 1, 16)
	// wgtArr := strArrToFlat(wgtStrArr, 32, 16)
	// outArr := strArrToFlat(outStrArr, 1, 32)
	// fmt.Println(inArr)
	// fmt.Println(outArr)
	// out := dense(inArr[0], wgtArr, 0, "")
	// fmt.Println(out)
	// out = dense(inArr[0], wgtArr, 1, "relu")
	// fmt.Println(out)
	// input := []float32{0.05, 0.1}
	// w1 := [][]float32{{0.15, 0.2}, {0.25, 0.3}}
	// w2 := [][]float32{{0.40, 0.45}, {0.50, 0.55}}
	// bias := []float32{0.35, 0.6}
	// out1 := dense(input, w1, bias[0], "sigmoid")
	// fmt.Println(out1)
	// out2 := dense(out1, w2, bias[1], "sigmoid")
	denseBackpropagation()

}

func dense(in []float32, wgt [][]float32, bias float32, act string) []float32 {

	var activ activation
	switch act {
	case "relu":
		activ = relu
	case "sigmoid":
		activ = sigmoid
	default:
		activ = pass
	}
	out := make([]float32, len(wgt))
	for i, k := range wgt {
		var temp float32 = 0
		for j := range in {
			temp += in[j] * k[j]
			// if i == 19 {
			// 	fmt.Println(in[j], k[j], temp)
			// }
		}
		out[i] = activate(temp, bias, activ)
		// if relu {
		// 	out[i] = max(0, temp+bias)
		// } else {
		// 	out[i] = temp + bias
		// }
	}
	return out
}

func denseBackpropagation() {

	input := []float32{0.05, 0.1}
	w1 := [][]float32{{0.15, 0.2}, {0.25, 0.3}}
	w2 := [][]float32{{0.40, 0.45}, {0.50, 0.55}}
	bias := []float32{0.35, 0.6}

	net1 := dense(input, w1, bias[0], "")
	fmt.Println(net1)
	out1 := dense(input, w1, bias[0], "sigmoid")
	fmt.Println(out1)
	net2 := dense(out1, w2, bias[1], "")
	fmt.Println(net2)
	out2 := dense(out1, w2, bias[1], "sigmoid")
	fmt.Println(out2)

	target := []float32{0.01, 0.99}
	outLoss := loss(out2, target)
	fmt.Println(outLoss)
	totalLoss := sumArr(outLoss)
	fmt.Println("Total loss:", totalLoss)

}

func loss(arr []float32, target []float32) []float32 {
	resArr := make([]float32, len(arr))
	for i := range arr {
		temp := target[i] - arr[i]
		resArr[i] = 0.5 * temp * temp
	}
	return resArr
}

func sumArr(arr []float32) float32 {
	result := float32(0.0)
	for i := range arr {
		result += arr[i]
	}
	return result
}

// func conv2d(in, wgt tensor, relu, padding bool) tensor {
//
// 	return nil
// }

func max(x, y float32) float32 {
	if x > y {
		return x
	}
	return y
}

type activation func(float32) float32

func sigmoid(x float32) float32 {
	return float32(1 / (1 + math.Exp(float64(x)*(-1))))
}
func relu(x float32) float32 {
	return max(0, x)
}
func pass(x float32) float32 {
	return x
}

func sigmoidDer(x float32) float32 {
	return x * (1 - x)
}

func activate(x, b float32, fn activation) float32 {
	raw := x + b
	return fn(raw)
}
func readFile(fileName string) []string {

	inRead, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	str := string(inRead)
	str = strings.ReplaceAll(str, "\n", " ")
	strArr := make([]string, 0)
	strArr = strings.Split(str, " ")

	return strArr
}

func strArrToFlat(s []string, x, y int) [][]float32 {
	if len(s) < x*y {
		log.Fatalf("func strArrToTensor: Input string length %d smaller than array size %d.\n", len(s), x*y)
	}
	si := 0
	arr := make([][]float32, x)
	for i := 0; i < x; i++ {
		arr[i] = make([]float32, y)
		for j := 0; j < y; j++ {
			if len(s[si]) > 0 {
				temp, err := strconv.ParseFloat(s[si], 32)
				if err != nil {
					log.Fatalln(err, i, j, temp, s[si], len(s[si]))
				}
				arr[i][j] = float32(temp)
			}
			si++
		}
	}
	return arr
}

func strArrToTensor(s []string, b, h, w, c int) tensor {

	if len(s) < b*h*w*c {
		log.Fatalf("func strArrToTensor: Input string length %d smaller than array size %d.\n", len(s), b*h*w*c)
	}
	arr := make([][][][]float32, b)
	si := 0
	for i := 0; i < b; i++ {
		arr[i] = make([][][]float32, h)
		for j := 0; j < h; j++ {
			arr[i][j] = make([][]float32, w)
			for k := 0; k < w; k++ {
				arr[i][j][k] = make([]float32, c)
				for l := 0; l < c; l++ {
					if len(s[si]) > 0 {
						temp, err := strconv.ParseFloat(s[si], 32)
						if err != nil {
							log.Fatalln(err, i, j, k, l, temp, s[si], len(s[si]))
						}
						arr[i][j][k][l] = float32(temp)
					}
					si++
				}
			}
		}
	}
	return arr
}
