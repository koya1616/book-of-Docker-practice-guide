// モンテカルロシミュレーションによって円周率の近似値を求める
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	count := 0
	for i := 0; i < 10000; i++ {
		x := rand.Float64()
		y := rand.Float64()
		if (x*x + y*y) < 1 {
			count++
		}
	}
	fmt.Println(4 * float64(count) / 10000)
}