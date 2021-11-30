package main

import "testing"

func TestPerimeter(t *testing.T) {
  rectangle := Rectangle{10., 10.}
  got := Perimeter(rectangle)
  want := 40.
  if got != want {
    t.Errorf("got %.2f but want %.2f", got, want)
  }
}

func TestArea(t *testing.T) {
  areaTests := []struct{
    name string
    shape Shape
    hasArea float64
  }{
    {name: "Rectangle", shape: Rectangle{Width: 3., Height: 5.}, hasArea: 15.},
    {name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
    {name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.},
  }
  for _, test := range areaTests {
    t.Run(test.name, func(t *testing.T) {
      got := test.shape.Area()
      if got != test.hasArea {
        t.Errorf("%#v got %g but want %g", test.shape, got, test.hasArea)
      }
    })
  }
}
