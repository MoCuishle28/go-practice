package main


import(
	"math"
)


// 2 points sum
func cal(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}


func main() {

}