package main

import (
	"fmt"
	"strconv"

	"github.com/paulmach/go.geo"
)

func main() {
	var d string
	path := geo.NewPath()
	path.Push(geo.NewPoint(0, 0))
	var x = 0
	var y = 0
	var look = 0
	var lastPoint = geo.NewPoint(0, 0)
	var addPoint = geo.NewPoint(0, 0)
	var prevPoint = geo.NewPoint(0, 0)
	for {
		_, err := fmt.Scan(&d)
		if err != nil {
			break
		}
		var direction = d[0:1]
		if direction != "L" && direction != "R" {
			continue
		}
		var distanceString = d[1:len(d)]
		if distanceString[len(distanceString)-1] == ',' {
			distanceString = distanceString[:len(distanceString)-1]
		}
		var distance, _ = strconv.Atoi(distanceString)

		//fmt.Printf("%d, %d, %d\n", x, y, look)

		//fmt.Printf("%s,%d\n", direction, distance)
		if direction == "L" {
			look -= 1
		} else {
			look += 1
		}
		look %= 4
		if look < 0 {
			look += 4
		}
		switch look {
		case 0:
			y += distance
		case 1:
			x += distance
		case 2:
			y -= distance
		case 3:
			x -= distance
		}
		var newPoint = geo.NewPoint(float64(x), float64(y))
		line := geo.NewLine(lastPoint, newPoint)
		if path.Intersects(line) {
			points, _ := path.Intersection(line)
			x = int(points[0][0])
			y = int(points[0][1])
			break
		}
		addPoint = prevPoint
		prevPoint = lastPoint
		lastPoint = newPoint
		path.Push(addPoint)
	}

	if y < 0 {
		y = -y
	}
	if x < 0 {
		x = -x
	}
	fmt.Printf("%d\n", x+y)
}
