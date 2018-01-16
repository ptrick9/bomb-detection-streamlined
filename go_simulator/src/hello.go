package main
import "fmt"
import "math/rand"

//var grid[][]Node
//var nodeList[]NodeParent
var maxX int
var maxY int



func rangeInt(min, max int) int {
	return rand.Intn(max-min) + min
}


func main() {
	maxX = 10
	maxY = 10

	//nodeList = make([]Node, 10)
	/*var nodeList []interface{move()
	sensorReport() float32}*/
	nodeList := make([]NodeMovement, 10)
	for i := 0; i < 10; i++ {
		if i%3 == 0 {
			nodeList[i] = &bn{NodeImpl:&NodeImpl{x: rand.Intn(10), y: rand.Intn(10), id: i}, x_speed: rangeInt(-3, 3), y_speed: rangeInt(-3, 3)}
			fmt.Print("added bn")

		} else if i%3 == 1 {
			nodeList[i] = &wn{NodeImpl:&NodeImpl{x: rand.Intn(10), y: rand.Intn(10), id: i}, speed: rangeInt(-3, 3), dir: rand.Intn(2)}
		} else {
			nodeList[i] = &rn{NodeImpl:&NodeImpl{x: rand.Intn(10), y: rand.Intn(10), id: i}}
		}
	}

	for i := 0; i < 10; i++ {
		for _, v := range nodeList {
			fmt.Println(v)
			v.move()
			fmt.Println(v.sensorReport())
		}
		fmt.Println("-------------")
	}

}
