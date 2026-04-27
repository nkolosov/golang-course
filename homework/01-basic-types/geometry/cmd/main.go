package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"geometry-app/pkg/geometry"

	"github.com/spf13/pflag"
)

func main() {
	// Define flags
	distanceFlag := pflag.Bool("distance", false, "Calculate distance between two points")
	pointFlags := pflag.StringArray("point", []string{}, "Point in format X,Y (can be used multiple times)")
	areaFlag := pflag.Bool("area", false, "Calculate area of figure")
	circleFlag := pflag.String("circle", "", "Circle in format X,Y,R")
	centerFlag := pflag.String("center", "", "Center point of circle in format X,Y")
	radiusFlag := pflag.Float64("radius", 0, "Radius of circle")
	polygonFlag := pflag.Bool("polygon", false, "Use polygon figure")
	containsFlag := pflag.Bool("contains", false, "Check if point is inside figure")

	pflag.Parse()

	// Validate arguments
	if *distanceFlag {
		if len(*pointFlags) != 2 {
			fmt.Println("Error: --distance requires exactly 2 points")
			os.Exit(1)
		}
		handleDistance(*pointFlags)
	} else if *areaFlag {
		handleArea(circleFlag, centerFlag, radiusFlag, polygonFlag, pointFlags)
	} else if *containsFlag {
		if len(*pointFlags) != 1 {
			fmt.Println("Error: --contains requires exactly 1 point")
			os.Exit(1)
		}
		handleContains(circleFlag, centerFlag, radiusFlag, polygonFlag, pointFlags)
	} else {
		fmt.Println("Error: No valid command specified")
		pflag.Usage()
		os.Exit(1)
	}
}

func handleDistance(pointFlags []string) {
	p1, err := parsePoint(pointFlags[0])
	if err != nil {
		fmt.Printf("Error parsing first point: %v\n", err)
		os.Exit(1)
	}

	p2, err := parsePoint(pointFlags[1])
	if err != nil {
		fmt.Printf("Error parsing second point: %v\n", err)
		os.Exit(1)
	}

	distance := p1.DistanceTo(p2)
	fmt.Printf("Distance: %.2f\n", distance)
}

func handleArea(circleFlag *string, centerFlag *string, radiusFlag *float64, polygonFlag *bool, pointFlags *[]string) {
	var figure geometry.Figure

	if *circleFlag != "" {
		c, err := parseCircle(*circleFlag)
		if err != nil {
			fmt.Printf("Error parsing circle: %v\n", err)
			os.Exit(1)
		}
		figure = c
	} else if *centerFlag != "" && *radiusFlag != 0 {
		center, err := parsePoint(*centerFlag)
		if err != nil {
			fmt.Printf("Error parsing center point: %v\n", err)
			os.Exit(1)
		}
		c := geometry.Circle{
			Center: center,
			Radius: *radiusFlag,
		}
		figure = c
	} else if *polygonFlag {
		points, err := parsePoints(*pointFlags)
		if err != nil {
			fmt.Printf("Error parsing polygon points: %v\n", err)
			os.Exit(1)
		}
		p := geometry.Polygon{
			Points: points,
		}
		figure = p
	} else {
		fmt.Println("Error: --area requires specifying a figure (circle or polygon)")
		os.Exit(1)
	}

	area := figure.Area()
	fmt.Printf("Area: %.2f\n", area)
}

func handleContains(circleFlag *string, centerFlag *string, radiusFlag *float64, polygonFlag *bool, pointFlags *[]string) {
	point, err := parsePoint((*pointFlags)[0])
	if err != nil {
		fmt.Printf("Error parsing point: %v\n", err)
		os.Exit(1)
	}

	var figure geometry.Figure

	if *circleFlag != "" {
		c, err := parseCircle(*circleFlag)
		if err != nil {
			fmt.Printf("Error parsing circle: %v\n", err)
			os.Exit(1)
		}
		figure = c
	} else if *centerFlag != "" && *radiusFlag != 0 {
		center, err := parsePoint(*centerFlag)
		if err != nil {
			fmt.Printf("Error parsing center point: %v\n", err)
			os.Exit(1)
		}
		c := geometry.Circle{
			Center: center,
			Radius: *radiusFlag,
		}
		figure = c
	} else if *polygonFlag {
		points, err := parsePoints(*pointFlags)
		if err != nil {
			fmt.Printf("Error parsing polygon points: %v\n", err)
			os.Exit(1)
		}
		p := geometry.Polygon{
			Points: points,
		}
		figure = p
	} else {
		fmt.Println("Error: --contains requires specifying a figure (circle or polygon)")
		os.Exit(1)
	}

	contains := figure.Contains(point)
	fmt.Printf("Contains: %t\n", contains)
}

func parsePoint(pointStr string) (geometry.Point, error) {
	parts := strings.Split(pointStr, ",")
	if len(parts) != 2 {
		return geometry.Point{}, errors.New("invalid point format, expected X,Y")
	}

	x, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return geometry.Point{}, fmt.Errorf("invalid X coordinate: %v", err)
	}

	y, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return geometry.Point{}, fmt.Errorf("invalid Y coordinate: %v", err)
	}

	return geometry.Point{X: x, Y: y}, nil
}

func parseCircle(circleStr string) (geometry.Circle, error) {
	parts := strings.Split(circleStr, ",")
	if len(parts) != geometry.MinPoints {
		return geometry.Circle{}, errors.New("invalid circle format, expected X,Y,R")
	}

	centerX, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return geometry.Circle{}, fmt.Errorf("invalid center X coordinate: %v", err)
	}

	centerY, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return geometry.Circle{}, fmt.Errorf("invalid center Y coordinate: %v", err)
	}

	radius, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		return geometry.Circle{}, fmt.Errorf("invalid radius: %v", err)
	}

	return geometry.Circle{
		Center: geometry.Point{X: centerX, Y: centerY},
		Radius: radius,
	}, nil
}

func parsePoints(pointFlags []string) ([]geometry.Point, error) {
	points := make([]geometry.Point, 0, len(pointFlags))
	for _, pointStr := range pointFlags {
		if pointStr == "" {
			continue
		}
		point, err := parsePoint(pointStr)
		if err != nil {
			return nil, err
		}
		points = append(points, point)
	}
	return points, nil
}
