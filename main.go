package main

import (
	"fmt"
	"math"

	"github.com/hschendel/stl"
)

const (
	sideLen = 14
)

const (
	invalid = iota
	piston
	tip
	tilt
	defocus
	astigo
	astigv
	comav
	comah
	trefv
	trefo
	spherical
)

func main() {
	s := &stl.Solid{
		Name:    "Zernike Shapes",
		IsAscii: true,
	}
	addZernikeShape(s, piston, 0, 0)
	addZernikeShape(s, tip, 1, 0)
	addZernikeShape(s, tilt, 1, 1)
	addZernikeShape(s, defocus, 2, 0)
	addZernikeShape(s, astigo, 2, 1)
	addZernikeShape(s, astigv, 2, 2)
	addZernikeShape(s, comav, 3, 0)
	addZernikeShape(s, comah, 3, 1)
	addZernikeShape(s, trefv, 3, 2)
	addZernikeShape(s, trefo, 3, 3)
	addZernikeShape(s, spherical, 4, 0)

	// Clean it up and write it out.
	s.RecalculateNormals()
	s.WriteFile("/Users/timboldt/Desktop/test.stl")
}

func addZernikeShape(s *stl.Solid, noll int, row, col int) {
	xlateX := float32(col * (sideLen - 1))
	xlateY := float32(-row * (sideLen - 1))
	heights := getHeightMap(noll)
	for row := 1; row < sideLen; row++ {
		for col := 1; col < sideLen; col++ {
			if col == 0 {
				continue
			}
			s.AppendTriangle(stl.Triangle{
				Vertices: [3]stl.Vec3{
					{xlateX + float32(row-1), xlateY + float32(col-1), heights[row-1][col-1]},
					{xlateX + float32(row), xlateY + float32(col-1), heights[row][col-1]},
					{xlateX + float32(row), xlateY + float32(col), heights[row][col]},
				},
			})
			s.AppendTriangle(stl.Triangle{
				Vertices: [3]stl.Vec3{
					{xlateX + float32(row-1), xlateY + float32(col-1), heights[row-1][col-1]},
					{xlateX + float32(row), xlateY + float32(col), heights[row][col]},
					{xlateX + float32(row-1), xlateY + float32(col), heights[row-1][col]},
				},
			})
			fmt.Printf("%2d %2d %6f\n", row, col, heights[row][col])
		}
	}

	// Add base.
	s.AppendTriangle(stl.Triangle{
		Vertices: [3]stl.Vec3{
			{xlateX + 0, xlateY + 0, 0},
			{xlateX + 0, xlateY + float32(sideLen-1), 0},
			{xlateX + float32(sideLen-1), xlateY + float32(sideLen-1), 0},
		},
	})
	s.AppendTriangle(stl.Triangle{
		Vertices: [3]stl.Vec3{
			{xlateX + 0, xlateY + 0, 0},
			{xlateX + float32(sideLen-1), xlateY + float32(sideLen-1), 0},
			{xlateX + float32(sideLen-1), xlateY + 0, 0},
		},
	})

	// Add sides.
	for i := 1; i < sideLen; i++ {
		// Front side.
		s.AppendTriangle(stl.Triangle{
			Vertices: [3]stl.Vec3{
				{xlateX + 0, xlateY + float32(i-1), 0},
				{xlateX + 0, xlateY + float32(i-1), heights[0][i-1]},
				{xlateX + 0, xlateY + float32(i), heights[0][i]},
			},
		})
		s.AppendTriangle(stl.Triangle{
			Vertices: [3]stl.Vec3{
				{xlateX + 0, xlateY + float32(i-1), 0},
				{xlateX + 0, xlateY + float32(i), heights[0][i]},
				{xlateX + 0, xlateY + float32(i), 0},
			},
		})
		// Back side.
		s.AppendTriangle(stl.Triangle{
			Vertices: [3]stl.Vec3{
				{xlateX + float32(sideLen-1), xlateY + float32(i-1), 0},
				{xlateX + float32(sideLen-1), xlateY + float32(i), heights[sideLen-1][i]},
				{xlateX + float32(sideLen-1), xlateY + float32(i-1), heights[sideLen-1][i-1]},
			},
		})
		s.AppendTriangle(stl.Triangle{
			Vertices: [3]stl.Vec3{
				{xlateX + float32(sideLen-1), xlateY + float32(i-1), 0},
				{xlateX + float32(sideLen-1), xlateY + float32(i), 0},
				{xlateX + float32(sideLen-1), xlateY + float32(i), heights[sideLen-1][i]},
			},
		})
		// Left side.
		s.AppendTriangle(stl.Triangle{
			Vertices: [3]stl.Vec3{
				{xlateX + float32(i-1), xlateY + 0, 0},
				{xlateX + float32(i), xlateY + 0, heights[i][0]},
				{xlateX + float32(i-1), xlateY + 0, heights[i-1][0]},
			},
		})
		s.AppendTriangle(stl.Triangle{
			Vertices: [3]stl.Vec3{
				{xlateX + float32(i-1), xlateY + 0, 0},
				{xlateX + float32(i), xlateY + 0, 0},
				{xlateX + float32(i), xlateY + 0, heights[i][0]},
			},
		})
		// Right side.
		s.AppendTriangle(stl.Triangle{
			Vertices: [3]stl.Vec3{
				{xlateX + float32(i-1), xlateY + float32(sideLen-1), 0},
				{xlateX + float32(i-1), xlateY + float32(sideLen-1), heights[i-1][sideLen-1]},
				{xlateX + float32(i), xlateY + float32(sideLen-1), heights[i][sideLen-1]},
			},
		})
		s.AppendTriangle(stl.Triangle{
			Vertices: [3]stl.Vec3{
				{xlateX + float32(i-1), xlateY + float32(sideLen-1), 0},
				{xlateX + float32(i), xlateY + float32(sideLen-1), heights[i][sideLen-1]},
				{xlateX + float32(i), xlateY + float32(sideLen-1), 0},
			},
		})
	}
}

func getHeightMap(noll int) [sideLen][sideLen]float32 {
	const (
		offset = float32(sideLen-1) / 2
		scale  = 2.0
	)
	heights := [sideLen][sideLen]float32{}
	for row := 1; row < sideLen-1; row++ {
		for col := 1; col < sideLen-1; col++ {
			rho, phi := cartesianToPolar(float64(row), float64(col))
			switch noll {
			case tip:
				heights[row][col] = float32(2 * rho * math.Cos(phi))
			case tilt:
				heights[row][col] = float32(2 * rho * math.Sin(phi))
			case defocus:
				heights[row][col] = -1 * float32(math.Sqrt(3)*(2*rho*rho-1))
			case astigo:
				heights[row][col] = float32(math.Sqrt(6) * rho * rho * math.Sin(2*phi))
			case astigv:
				heights[row][col] = float32(math.Sqrt(6) * rho * rho * math.Cos(2*phi))
			case comav:
				heights[row][col] = float32(math.Sqrt(8) * (3*rho*rho*rho - 2*rho) * math.Sin(phi))
			case comah:
				heights[row][col] = float32(math.Sqrt(8) * (3*rho*rho*rho - 2*rho) * math.Cos(phi))
			case trefv:
				heights[row][col] = float32(math.Sqrt(8) * rho * rho * rho * math.Sin(3*phi))
			case trefo:
				heights[row][col] = float32(math.Sqrt(8) * rho * rho * rho * math.Cos(3*phi))
			case spherical:
				heights[row][col] = float32(math.Sqrt(5) * (6*rho*rho*rho*rho - 6*rho*rho + 1))
			default:
				heights[row][col] = 0
			}
			heights[row][col] = heights[row][col]*scale + offset
		}
	}
	// Edges and corners.
	for i := 0; i < sideLen; i++ {
		heights[0][i] = 1
		heights[sideLen-1][i] = 1
		heights[i][0] = 1
		heights[i][sideLen-1] = 1
	}
	heights[12][1] = 1
	heights[1][12] = 1
	heights[12][12] = 1
	// Handle the edges.
	for i := 1; i < sideLen-1; i++ {
		heights[1][i] = (heights[0][i] + heights[1][i] + heights[2][i]) / 3.0
		heights[sideLen-2][i] = (heights[sideLen-3][i] + heights[sideLen-2][i] + heights[sideLen-1][i]) / 3.0
		heights[i][1] = (heights[i][0] + heights[i][1] + heights[i][2]) / 3.0
		heights[i][sideLen-2] = (heights[i][sideLen-3] + heights[i][sideLen-2] + heights[i][sideLen-1]) / 3.0
	}
	return heights
}

func cartesianToPolar(row, col float64) (float64, float64) {
	width := float64(sideLen - 1)
	radius := width / 2
	x := row - (width+1)/2
	y := col - (width+1)/2
	return math.Sqrt(x*x+y*y) / math.Sqrt(radius*radius+radius*radius), math.Atan2(y, x)
}
