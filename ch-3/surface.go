package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                //axis ranges (-xyrange ..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x x or y unit
	zscale        = height * 0.4        //pixels per x unit
	angle         = math.Pi / 6         // angle of x, y axes(=30)

)

var sin30, cos30 = math.Sin(angle), math.Cos((angle)) // sin 30, cos 30

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func corner(i, j int) (float64, float64) {
	//find point(x,y) at the corner of the cell (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	//compute surface height z
	z := f(x, y)

	//project(x,y,z) isometrically onto 2-D SVG canvas(sx, sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) //distance form (0,0)
	return math.Sin(r) / r
}

//handler echoes the http request
func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "images/svg+xml")
	make_svg(w)
}

func make_svg(out io.Writer) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}

	fmt.Fprintf(out, "</svg>")
}
