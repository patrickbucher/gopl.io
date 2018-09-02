package ex14

import (
	"html/template"
	"strings"
)

// TODO link number to issue details
// TODO link user to user details
// TODO link milestone to milestone details
var IssueListHTML = `
<html>
	<head>
		<title>GitHub Issues</title>
		<style type="text/css">
		th { text-align: left; }
		th, td { padding: 0 1em 0 0; }
		.num { text-align: right; }
		</style>
	</head>
	<body>
		<h1>{{.Count}} Issues</h1>
		<p>For Query: <em>{{.SearchTerms}}</em></p>
		<table>
			<tr>
				<th>Repository</th>
				<th class="num">#</th>
				<th>State</th>
				<th>User</th>
				<th>Milestone</th>
				<th>Title</th>
			</tr>
			{{range .Issues}}
			<tr>
				<td>{{.RepositoryURL | extractRepoPath}}</td>
				<td class="num">{{.Number}}</td>
				<td>{{.State}}</td>
				<td>{{.User.Name}}</td>
				<td>{{.Milestone.Number}} {{.Milestone.Title}}</td>
				<td>{{.Title}}</td>
			</tr>
			{{end}}
		</table>
	</body>
</html>
`

var IssueListTemplate = template.Must(template.New("issueList").
	Funcs(template.FuncMap{
		"extractRepoPath": func(url string) string {
			elems := strings.Split(url, "/")
			if len(elems) < 2 {
				return ""
			}
			return elems[len(elems)-2] + "/" + elems[len(elems)-1]
		},
	}).
	Parse(IssueListHTML))

// TODO single issue template
// TODO user template
// TODO milestone template
