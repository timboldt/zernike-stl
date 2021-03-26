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
	const (
		sideLen = 3
	)
	heights := [sideLen][sideLen]float32{
		{1.1, 1.3, 1.2},
		{1.3, 1.3, 1.4},
		{1.2, 1.4, 1.2},
	}
	for row := 1; row < sideLen; row++ {
		for col := 1; col < sideLen; col++ {
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

	// Add base.
	s.AppendTriangle(stl.Triangle{
		Vertices: [3]stl.Vec3{
			{0, 0, 0},
			{0, float32(sideLen - 1), 0},
			{float32(sideLen - 1), float32(sideLen - 1), 0},
		},
	})
	s.AppendTriangle(stl.Triangle{
		Vertices: [3]stl.Vec3{
			{0, 0, 0},
			{float32(sideLen - 1), float32(sideLen - 1), 0},
			{float32(sideLen - 1), 0, 0},
		},
	})

	// Add sides.
	for i := 1; i < sideLen; i++ {
		// Front side.
		s.AppendTriangle(stl.Triangle{
			Vertices: [3]stl.Vec3{
				{0, float32(i - 1), 0},
				{0, float32(i - 1), heights[0][i-1]},
				{0, float32(i), heights[0][i]},
			},
		})
		s.AppendTriangle(stl.Triangle{
			Vertices: [3]stl.Vec3{
				{0, float32(i - 1), 0},
				{0, float32(i), heights[0][i]},
				{0, float32(i), 0},
			},
		})
		// Back side.
		s.AppendTriangle(stl.Triangle{
			Vertices: [3]stl.Vec3{
				{float32(sideLen - 1), float32(i - 1), 0},
				{float32(sideLen - 1), float32(i), heights[sideLen-1][i]},
				{float32(sideLen - 1), float32(i - 1), heights[sideLen-1][i-1]},
			},
		})
		s.AppendTriangle(stl.Triangle{
			Vertices: [3]stl.Vec3{
				{float32(sideLen - 1), float32(i - 1), 0},
				{float32(sideLen - 1), float32(i), 0},
				{float32(sideLen - 1), float32(i), heights[sideLen-1][i]},
			},
		})
		// Left side.
		s.AppendTriangle(stl.Triangle{
			Vertices: [3]stl.Vec3{
				{float32(i - 1), 0, 0},
				{float32(i), 0, heights[i][0]},
				{float32(i - 1), 0, heights[i-1][0]},
			},
		})
		s.AppendTriangle(stl.Triangle{
			Vertices: [3]stl.Vec3{
				{float32(i - 1), 0, 0},
				{float32(i), 0, 0},
				{float32(i), 0, heights[i][0]},
			},
		})
		// Right side.
		s.AppendTriangle(stl.Triangle{
			Vertices: [3]stl.Vec3{
				{float32(i - 1), float32(sideLen - 1), 0},
				{float32(i - 1), float32(sideLen - 1), heights[i-1][sideLen-1]},
				{float32(i), float32(sideLen - 1), heights[i][sideLen-1]},
			},
		})
		s.AppendTriangle(stl.Triangle{
			Vertices: [3]stl.Vec3{
				{float32(i - 1), float32(sideLen - 1), 0},
				{float32(i), float32(sideLen - 1), heights[i][sideLen-1]},
				{float32(i), float32(sideLen - 1), 0},
			},
		})
	}

	// Clean it up and write it out.
	s.RecalculateNormals()
	s.WriteFile("/Users/timboldt/Desktop/test.stl")
}
