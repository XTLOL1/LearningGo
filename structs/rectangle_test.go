package main

import "testing"

func TestStructs(t *testing.T) {
	areaTests := []struct {
		name string
		shape Shape
		want  float64
	}{
		{"Rectangle", Rectangle{12, 6}, 72.0},
		{"Circle", Circle{10}, 314.1592653589793},
		{"Triangle", Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func (t *testing.T){
			got := tt.shape.area()
			if got != tt.want {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
			}
		})
	}	

}
