package gps

import (
	"log"
	"testing"
)

func TestGPSDistance(t *testing.T) {
	g1 := NewGPS(45.24220494, 19.70152105)
	g2 := NewGPS(45.25219151, 19.83728467)
	diss := HaversineDistance(g1, g2)
	log.Println("distance: ", diss)

	if diss != 10.68634049002477 {
		t.Fatalf("distance error")
	}
}
