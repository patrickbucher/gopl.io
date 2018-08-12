package lenconv

// MToF converts a meter into feet.
func MToF(m Meter) Foot { return Foot(m / OneFootInMeters) }

// FToM converts a foot into meters.
func FToM(f Foot) Meter { return Meter(f * OneFootInMeters) }
