package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"gopl.io/ch07/eval"
)

func main() {
	http.HandleFunc("/", input)
	http.HandleFunc("/calc", calculate)
	http.ListenAndServe("0.0.0.0:8000", nil)
}

func input(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "only GET supported", http.StatusMethodNotAllowed)
		return
	}
	inputHTML, err := os.Open("input.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("open input.html: %v", err), http.StatusInternalServerError)
		return
	}
	defer inputHTML.Close()
	io.Copy(w, inputHTML)
}

func calculate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "only POST supported", http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, fmt.Sprintf("parsing form: %v", err), http.StatusBadRequest)
		return
	}
	formula := strings.TrimSpace(r.FormValue("formula"))
	if formula == "" {
		http.Error(w, "formula missing", http.StatusBadRequest)
		return
	}
	variables := strings.TrimSpace(r.FormValue("variables"))
	environment, err := eval.ParseVarDefs(variables)
	if err != nil {
		msg := fmt.Sprintf("parsing definitions '%s': %v", variables, err)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	expression, err := eval.Parse(formula)
	if err != nil {
		msg := fmt.Sprintf("parsing formula '%s': %v", formula, err)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	result := expression.Eval(environment)
	w.Write([]byte(fmt.Sprintf("%s (%s) = %g", formula, variables, result)))
}
