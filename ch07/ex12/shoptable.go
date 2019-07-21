package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", http.HandlerFunc(db.list))
	http.HandleFunc("/price", http.HandlerFunc(db.price))
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

var tableHTML = `<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="utf-8">
		<title>Exercise 7.12: Shop HTML Table</title>
	</head>
	<body>
		<h1>Exercise 7.12: Shop HTML Table</h1>
		<table>
			<tr>
				<th>Item</th>
				<th>Price</th>
			</tr>
			{{ range $item, $price := . }}
			<tr>
				<td>{{ $item }}</td>
				<td>{{ $price }}</td>
			</tr>
			{{ end }}
		</table>
	</body>
</html>
`

var tableTemplate = template.Must(template.New("shop").Parse(tableHTML))

func (db database) list(w http.ResponseWriter, req *http.Request) {
	err := tableTemplate.Execute(w, db)
	if err != nil {
		msg := fmt.Sprintf("error executing template: %v", err)
		http.Error(w, msg, http.StatusInternalServerError)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "n such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}
