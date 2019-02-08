package main

import (
	"bytes"
	"fmt"
	"sort"
	"text/tabwriter"
)

type Employee struct {
	Name   string
	Salary int
	IQ     uint8
}

var employees = []*Employee{
	{"Catbert", 250000, 159},
	{"Dilbert", 120000, 120},
	{"Wally", 120000, 97},
	{"Alice", 120000, 130},
	{"Boss", 250000, 81},
}

type byName []*Employee

func (e byName) Len() int           { return len(e) }
func (e byName) Less(i, j int) bool { return e[i].Name < e[j].Name }
func (e byName) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }

type bySalary []*Employee

func (e bySalary) Len() int           { return len(e) }
func (e bySalary) Less(i, j int) bool { return e[i].Salary < e[j].Salary }
func (e bySalary) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }

type byIQ []*Employee

func (e byIQ) Len() int           { return len(e) }
func (e byIQ) Less(i, j int) bool { return e[i].IQ < e[j].IQ }
func (e byIQ) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }

type payroll []*Employee

func (p payroll) String() string {
	const format = "%v\t%v\t%3v\n"
	buf := bytes.NewBufferString("")
	tw := new(tabwriter.Writer).Init(buf, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Name", "Salary", "IQ")
	fmt.Fprintf(tw, format, "----", "------", "--")
	for _, e := range p {
		fmt.Fprintf(tw, format, e.Name, e.Salary, e.IQ)
	}
	tw.Flush()
	return buf.String()
}

func main() {
	var p payroll
	p = employees
	fmt.Println(p)

	sort.Sort(byName(employees))
	fmt.Println(p)

	sort.Sort(byIQ(employees))
	fmt.Println(p)

	sort.Sort(bySalary(employees))
	fmt.Println(p)
}
