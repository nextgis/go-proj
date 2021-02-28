package proj

import (
	"fmt"
	"testing"
)

func TestGeodesicArea(t *testing.T) {
	p, err := NewEPSG(4326)
	if err != nil {
		t.Fatal(err)
	}
	if p == nil {
		t.Fatal("Projection is nil")
	}

	xs := []float64{51.82, 43.48, 75.38, 51.82 }
	ys := []float64{63.80, 55.62, 59.13, 63.80 }

	area, perimeter, err := p.GeodesicArea(xs, ys)
	if err != nil {
		t.Fatal(err)
	}
	if area < 0.01 || perimeter < 0.01 {
		t.Fatal("Failed to get area or perimeter")
	}
	fmt.Printf("Area: %f\nPerimeter: %f\n", area, perimeter)
}

func TestGeodesicDistance(t *testing.T) {
	p, err := NewEPSG(4326)
	if err != nil {
		t.Fatal(err)
	}
	if p == nil {
		t.Fatal("Projection is nil")
	}

	xs := []float64{51.82, 43.48, 75.38, 51.82 }
	ys := []float64{63.80, 55.62, 59.13, 63.80 }

	dist, err := p.GeodesicDistance(xs, ys)
	if err != nil {
		t.Fatal(err)
	}
	if dist < 0.01 {
		t.Fatal("Failed to get area or perimeter")
	}
	fmt.Printf("Distance: %f\n", dist)
}
