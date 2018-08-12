// Package weightconv performs Kilogram and Pound conversions.
package weightconv

import "fmt"

type Kilogram float64
type Pound float64

const OnePoundInKilograms = 0.45359237

func (k Kilogram) String() string { return fmt.Sprintf("%gkg", k) }
func (p Pound) String() string    { return fmt.Sprintf("%glb", p) }
