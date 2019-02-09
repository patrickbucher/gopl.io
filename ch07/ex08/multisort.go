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
	Sex    string
}

var employees = []*Employee{
	{"Arthur", 170000, "M"},
	{"Alice", 160000, "F"},
	{"Benno", 150000, "M"},
	{"Bertha", 140000, "F"},
	{"Charles", 130000, "M"},
	{"Charlene", 120000, "F"},
	{"Daniel", 110000, "M"},
	{"Dorothea", 100000, "F"},
}

type Payroll []*Employee

type byName []*Employee

func (e byName) Len() int           { return len(e) }
func (e byName) Less(i, j int) bool { return e[i].Name < e[j].Name }
func (e byName) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }

type bySalary []*Employee

func (e bySalary) Len() int           { return len(e) }
func (e bySalary) Less(i, j int) bool { return e[i].Salary < e[j].Salary }
func (e bySalary) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }

type bySex []*Employee

func (e bySex) Len() int           { return len(e) }
func (e bySex) Less(i, j int) bool { return e[i].Sex < e[j].Sex }
func (e bySex) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }

type Multisort struct {
	criteria []sort.Interface
}

func (m Multisort) apply(criterion sort.Interface) {
	criteria := make([]sort.Interface, len(m.criteria))
	// recycle the criteria used before
	for _, s := range m.criteria {
		if criterion != s {
			// except the latest
			criteria = append(criteria, s)
		}
	}
	// use the fresh criterion at the end of the sort chain
	criteria = append(criteria, criterion)
	for _, s := range criteria {
		sort.Sort(s)
	}
	m.criteria = criteria
}

func (p Payroll) String() string {
	const format = "%v\t%v\t%v\n"
	buf := bytes.NewBufferString("")
	tw := new(tabwriter.Writer).Init(buf, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Name", "Salary", "Sex")
	fmt.Fprintf(tw, format, "----", "------", "---")
	for _, e := range p {
		fmt.Fprintf(tw, format, e.Name, e.Salary, e.Sex)
	}
	tw.Flush()
	return buf.String()
}

func main() {
	var sorter Multisort

	payroll := Payroll(employees)
	fmt.Println(payroll)

	salaryCriterion := bySalary(payroll)
	sexCriterion := bySex(payroll)
	nameCriterion := byName(payroll)

	// chained sort: "stable"
	sorter.apply(nameCriterion)
	sorter.apply(salaryCriterion)
	sorter.apply(sexCriterion)
	fmt.Println(payroll)

	// multiple sorts in a row: could be unstable... (is not)
	sort.Sort(nameCriterion) // back to start
	sort.Sort(salaryCriterion)
	sort.Sort(sexCriterion)
	fmt.Println(payroll)
}
