package main

import "fmt"

func main() {
  var waterTemperatureEbullitionInKelvin float32 = 373.2

  fmt.Println(convertKelvinToCelsius(waterTemperatureEbullitionInKelvin))
}

func convertKelvinToCelsius(k float32) float32 {
  return k - 273
}
