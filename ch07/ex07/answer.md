The help message contains °C, because the Celsius type implements a `String()`
method using that format, and the interface `flag.Value` requires this method.
