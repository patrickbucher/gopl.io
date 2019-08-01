package eval

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseVarDefs(varDefs string) (Env, error) {
	definitions := strings.Split(varDefs, ",")
	environment := make(map[Var]float64)
	for _, def := range definitions {
		parts := strings.Split(strings.TrimSpace(def), "=")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid definition '%s'", parts)
		}
		key := Var(parts[0])
		value, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			return nil, fmt.Errorf("parse '%s' as float: %v", parts[1], err)
			continue
		}
		environment[key] = value
	}
	return environment, nil
}
