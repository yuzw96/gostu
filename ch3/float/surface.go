package main

import (
	"fmt"
	"math"
	"os"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / height
	zscale        = height * 0.4
	angle         = math.Pi / 6 //30
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	filename := "test.svg"
	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {
		fmt.Println("创建文件失败")
	}

	_, err = f.Write([]byte(fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)))

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			if isFinite(ax) || isFinite(ay) ||
				isFinite(bx) || isFinite(by) ||
				isFinite(cx) || isFinite(cy) ||
				isFinite(dx) || isFinite(dy) {
				continue
			}

			_, err = f.Write([]byte(fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)))
		}
	}
	_, err = f.Write([]byte(fmt.Sprintf("</svg>")))
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func isFinite(f float64) bool {
	if math.IsInf(f, 0) || math.IsNaN(f) {
		return true
	}
	return false
}
