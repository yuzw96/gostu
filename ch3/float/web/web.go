package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

var (
	width, height = 600, 320
	xyscale       = width / 2 / xyrange
	zscale        = float64(height) * 0.4
)

const (
	cells   = 100
	xyrange = 30.0
	angle   = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "image/svg+xml")
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}

	if r.Form["width"] != nil {
		width, _ = strconv.Atoi(r.Form["width"][0])
	}

	if r.Form["height"] != nil {
		height, _ = strconv.Atoi(r.Form["height"][0])
	}

	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner1(i+1, j)
			bx, by := corner1(i, j)
			cx, cy := corner1(i, j+1)
			dx, dy := corner1(i+1, j+1)
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(w, "</svg>\n")
}
func corner1(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f1(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,xy).
	sx := float64(width)/2 + (x-y)*cos30*float64(xyscale)
	sy := float64(height)/2 + (x+y)*sin30*float64(xyscale) - z*zscale
	return sx, sy
}

func f1(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
