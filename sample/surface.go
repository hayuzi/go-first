package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // 像素表示的画布大小
	cells         = 100                 // 网格单元的个数
	xyrange       = 30.0                // 坐标轴范围 -xyrange..+xyrange
	xyscale       = width / 2 / xyrange // x或者y轴上每个单位长度的像素
	zscale        = height * 0.4        // z轴上每个长度的像素
	angle         = math.Pi / 6         // x，y轴的角度 ，弧度方式浮点数表示（=30度）
)

var sin30, cos30 = math.Sin(angle), math.Sin(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 1; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// 网格单元顶点（i，j）的坐标（x，y）
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// 计算曲面高度z
	z := f(x, y)
	// 将xyz等角投射到二维svg回头平面上，坐标是（sx，sy）
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // 到（0，0）的距离
	return math.Sin(r) / r
}
