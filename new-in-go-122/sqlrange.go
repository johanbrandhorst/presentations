package main

import "log"

// START OMIT
type Point struct {
	X float64 `sql:"x"`
	Y float64 `sql:"y"`
}

func main() {
	var points []Point
	for p, err := range sqlrange.Query[Point](db, `select x, y from points`) {
		if err != nil {
			log.Fatal("failed to query points:", err)
		}
		points = append(points, p)
	}
}

// END OMIT
