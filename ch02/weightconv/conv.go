package weightconv

// KToP converts a kilogram weight to pounds.
func KToP(k Kilogram) Pound { return Pound(k / OnePoundInKilograms) }

// PToK converts a pound weight to kilograms.
func PToK(p Pound) Kilogram { return Kilogram(p * OnePoundInKilograms) }
