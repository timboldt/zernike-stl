package main

import "fmt"
import "github.com/hschendel/stl"

func main() {
	s := &stl.Solid{
		Name:    "Simple",
		IsAscii: true,
		Triangles: []stl.Triangle{
			{
				Normal: stl.Vec3{0, 0, -1},
				Vertices: [3]stl.Vec3{
					{0, 0, 0},
					{0, 1, 0},
					{1, 0, 0},
				},
			},
			{
				Normal: stl.Vec3{0, -1, 0},
				Vertices: [3]stl.Vec3{
					{0, 0, 0},
					{1, 0, 0},
					{0, 0, 1},
				},
			},
			{
				Normal: stl.Vec3{0.57735, 0.57735, 0.57735},
				Vertices: [3]stl.Vec3{
					{0, 0, 1},
					{1, 0, 0},
					{0, 1, 0},
				},
			},
			{
				Normal: stl.Vec3{-1, 0, 0},
				Vertices: [3]stl.Vec3{
					{0, 0, 0},
					{0, 0, 1},
					{0, 1, 0},
				},
			},
		},
	}
	s.WriteFile("/tmp/test.stl")
	fmt.Println("hi")
}
