package units_test

import (
	"testing"

	"github.com/cfindlayisme/go-types/units"
	numerical "github.com/cfindlayisme/go-utils/numerical"
)

func TestDistanceConversions(t *testing.T) {
	d := units.NewDistanceFromMeters(10)
	if got := d.ToFeet(); !numerical.AlmostEqualFloats(got, 32.8084, 0.0001) {
		t.Errorf("ToFeet() = %v, want %v", got, 32.8084)
	}

	d2 := units.NewDistanceFromFeet(32.8084)
	if got := d2.ToMeters(); !numerical.AlmostEqualFloats(got, 10, 0.0001) {
		t.Errorf("ToMeters() = %v, want %v", got, 10)
	}
}

func TestDistanceAdd(t *testing.T) {
	d1 := units.NewDistanceFromMeters(10)
	d2 := units.NewDistanceFromMeters(5)
	result := d1.Add(d2)
	if got := result.ToMeters(); !numerical.AlmostEqualFloats(got, 15, 0.0001) {
		t.Errorf("Add() = %v, want %v", got, 15)
	}
}

func TestDistanceSubtract(t *testing.T) {
	d1 := units.NewDistanceFromMeters(10)
	d2 := units.NewDistanceFromMeters(5)
	result := d1.Subtract(d2)
	if got := result.ToMeters(); !numerical.AlmostEqualFloats(got, 5, 0.0001) {
		t.Errorf("Subtract() = %v, want %v", got, 5)
	}
}

func TestDistanceMixedUnits(t *testing.T) {
	d1 := units.NewDistanceFromMeters(10)
	d2 := units.NewDistanceFromFeet(16.4042) // approximately 5 meters
	result := d1.Add(d2)
	if got := result.ToMeters(); !numerical.AlmostEqualFloats(got, 15, 0.0001) {
		t.Errorf("Add() with mixed units = %v, want %v", got, 15)
	}

	result = d1.Subtract(d2)
	if got := result.ToMeters(); !numerical.AlmostEqualFloats(got, 5, 0.0001) {
		t.Errorf("Subtract() with mixed units = %v, want %v", got, 5)
	}
}
