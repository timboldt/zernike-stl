package main

import (
	"fmt"

	"github.com/hschendel/stl"
)

func main() {
	s := &stl.Solid{
		Name:    "Zernike Shapes",
		IsAscii: true,
	}
	heights := [][]float32{
		{1, 3, 2},
		{3, 3, 4},
		{2, 4, 2},
	}
	for row := range heights {
		if row == 0 {
			continue
		}
		for col := range heights[row] {
			if col == 0 {
				continue
			}
			s.AppendTriangle(stl.Triangle{
				Vertices: [3]stl.Vec3{
					{float32(row - 1), float32(col - 1), heights[row-1][col-1]},
					{float32(row), float32(col - 1), heights[row][col-1]},
					{float32(row), float32(col), heights[row][col]},
				},
			})
			s.AppendTriangle(stl.Triangle{
				Vertices: [3]stl.Vec3{
					{float32(row - 1), float32(col - 1), heights[row-1][col-1]},
					{float32(row), float32(col), heights[row][col]},
					{float32(row - 1), float32(col), heights[row-1][col]},
				},
			})
			fmt.Printf("row %d col %d\n", row, col)
		}
	}
	// 	Triangles: []stl.Triangle{
	// 		{
	// 			Vertices: [3]stl.Vec3{
	// 				{0, 0, 0},
	// 				{0, 1, 0},
	// 				{1, 0, 0},
	// 			},
	// 		},
	// 		{
	// 			Vertices: [3]stl.Vec3{
	// 				{0, 0, 0},
	// 				{1, 0, 0},
	// 				{0, 0, 1},
	// 			},
	// 		},
	// 		{
	// 			Vertices: [3]stl.Vec3{
	// 				{0, 0, 1},
	// 				{1, 0, 0},
	// 				{0, 1, 0},
	// 			},
	// 		},
	// 		{
	// 			Vertices: [3]stl.Vec3{
	// 				{0, 0, 0},
	// 				{0, 0, 1},
	// 				{0, 1, 0},
	// 			},
	// 		},
	// 	},
	// }
	s.RecalculateNormals()
	s.WriteFile("/Users/timboldt/Desktop/test.stl")
}
