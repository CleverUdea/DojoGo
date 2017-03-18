//Bladimir velez Rivera

package main

import (
  "fmt"
  "math"
)

func Sqrt(x float64) float64 {
  //Aqui va el código que implementa raiz
      //hint 1: usar ciclo que itere 10 veces
      //hint 2: z:= float64(x/2)

      z:= float64(x/2)
      count := 0
      for deltaZ := z; math.Abs(deltaZ) > 1e-6; {
            z2 := z
            z = z - ((z*z) - x )/(2*z)
            deltaZ = z - z2
            count ++
        }
        //fmt.Println(count)

      return z
}

func main() {
  fmt.Println(Sqrt(2))
}
