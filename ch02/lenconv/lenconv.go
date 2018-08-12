// Package lenconv performs Foot and Meter conversions.
package lenconv

import "fmt"

type Foot float64
type Meter float64

const OneFootInMeters = 0.3048

func (f Foot) String() string  { return fmt.Sprintf("%gft", f) }
func (m Meter) String() string { return fmt.Sprintf("%gm", m) }
