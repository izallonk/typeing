package engine

import (
    "math/rand/v2"
)

func Randomrange(min,max float64) float64{
    return min + rand.Float64() * (max - min)
}
