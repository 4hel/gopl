// Surface computes an SVG rendering a 3-D surface function
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axis (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			_, _, z := getXYZ(i, j)
			points := [8]float64{ax, ay, bx, by, cx, cy, dx, dy}
			valid := true
			for _, val := range points {
				if math.IsNaN(val) {
					valid = false
				}
			}
			if !valid {
				continue
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' "+
				"style='fill:#%06x' />\n",
				ax, ay, bx, by, cx, cy, dx, dy, color(z))
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	x, y, z := getXYZ(i, j)
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func getMinMaxZ() (float64, float64) {
	var min, max float64
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			_, _, z := getXYZ(i, j)
			if z < min {
				min = z
			}
			if z > max {
				max = z
			}
		}
	}
	return min, max
}

func getXYZ(i, j int) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)
	return x, y, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func g(x, y float64) float64 {
	r := math.Hypot(x, y)
	return -0.01 * (r * r)
}

// z coordinate is -1 .. +1
func color(z float64) uint64 {
	// redFactor -> 0..1, blueFactor 1..0
	redFactor := (z + 1) * 0.5
	blueFactor := (-z + 1) * 0.5

	// 00..ff, ff..00
	redByte := int(255.0 * redFactor)
	blueByte := int(255.0 * blueFactor)

	// left shift of red byte to its position
	var redWord uint64 = uint64(redByte) << 16
	var blueWord uint64 = uint64(blueByte)

	//ORing the words to get final color like 0x7f007f
	return blueWord | redWord
}
