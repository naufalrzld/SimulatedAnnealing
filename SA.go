package main

import (
  "fmt"
  "time"
  "math"
  "math/rand"
)

func randomFloat(min, max float64) float64 {

  return rand.Float64() * (max - min) + min
}

func f(x1, x2 float64) float64 {
  sin := math.Sin(x1)
  cos := math.Cos(x2)
  sqrt := math.Sqrt(math.Pow(x1, 2) + math.Pow(x2, 2))/math.Pi
  exp := math.Exp(math.Abs(1 - (sqrt)))

  return -math.Abs(sin * cos * exp)
}

func e(delD, temp float64) float64 {
  return math.Exp(-delD/temp)
}

func main() {
  rand.Seed(time.Now().UnixNano())
  var T float64 = 100
  var coolingRate float64 = 0.00182
  var x1 float64 = randomFloat(-10, 10)
  var x2 float64 = randomFloat(-10, 10)

  var initialState float64 = f(x1, x2)
  var currentState float64 = initialState
  var bestSoFar float64 = currentState
  for T >= 0.00182 {
    x1 = randomFloat(-10, 10)
    x2 = randomFloat(-10, 10)

    var newState float64 = f(x1, x2)
    var delD float64 = newState - currentState

    if (delD < 0) {
      currentState = newState
      bestSoFar = newState
      T = T - coolingRate;
    } else if (delD >= 0) {
      var e float64 = e(delD, T)
      var R float64 = randomFloat(0, 1)

      if (R <= e) {
        newState = currentState
        bestSoFar = newState
        T = T - coolingRate
      }
    }
  }
  fmt.Printf("Best So Far: %f\n", bestSoFar)
}
