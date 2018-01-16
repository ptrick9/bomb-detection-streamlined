package main
import "fmt"
import "math/rand"

var grid[][]Node
var nodeList[]Node
var maxX int
var maxY int



func rangeInt(min, max int) int {
	return rand.Intn(max-min) + min
}


func main() {
	maxX = 10
	maxY = 132
	grid = make([][]Node, 10)
	for i:= 0; i < 10; i++ {
		grid[i] = make([]Node, 10)
	}
	nodeList = make([]Node, 10)
	for i := 0; i < 10; i++ {
		if i%3 == 0 {
			nodeList[i] = &bn{x: rand.Intn(10), y: rand.Intn(10), id: i, x_speed: rangeInt(-3, 3), y_speed: rangeInt(-3, 3)}
			fmt.Print("added bn")

		} else if i%3 == 1 {
			nodeList[i] = &wn{x: rand.Intn(10), y: rand.Intn(10), id: i, speed: rangeInt(-3, 3), dir: rand.Intn(2)}
		} else {
			nodeList[i] = &rn{x: rand.Intn(10), y: rand.Intn(10), id: i}
		}
	}

	for i := 0; i < 10; i++ {
		for _, v := range nodeList {
			fmt.Println(v)
			v.move()
		}
		fmt.Println("-------------")
	}

}
