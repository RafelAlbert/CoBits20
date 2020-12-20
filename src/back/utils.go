package main

import "math"

func dist(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
  return math.Sqrt((x1-x2)*(x1-x2)+(y1-y2)*(y1-y2))
}
