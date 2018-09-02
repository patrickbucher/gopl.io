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
		<title>GitHub Issues for "{{.SearchTerms}}"</title>
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
				<td class="num">
					<a href="issue-{{.Id}}.html">{{.Number}}</a>
				</td>
				<td>{{.State}}</td>
				<td>
					<a href="user-{{.User.Id}}.html">{{.User.Name}}</a>
				</td>
				<td>
					<a href="milestone-{{.Milestone.Id}}.html">
						{{.Milestone.Number}} {{.Milestone.Title}}
					</a>
				</td>
				<td>{{.Title}}</td>
			</tr>
			{{end}}
		</table>
	</body>
</html>
`

var IssueListTemplate = template.Must(template.New("issueList").
	Funcs(template.FuncMap{
		"extractRepoPath": extractRepoPath,
	}).
	Parse(IssueListHTML))

var IssuePageHTML = `
<html>
	<head>
		<title>GitHub Issue #{{.Id}}</title>
		<style type="text/css">
		span { padding: 5px; border-radius: 5px; }
		th { text-align: left; }
		th, td { padding: 0 1em 0 0; }
		tr { margin: 5px 0; }
		</style>
	</head>
	<body>
		<h1>{{.Title}}</h1>
		<table>
			<tr>
				<th>Repository</th><td>{{.RepositoryURL | extractRepoPath}}</td>
			</tr>
			<tr>
				<th>Id</th><td>{{.Id}}</td>
			</tr>
			<tr>
				<th>Number</th><td>{{.Number}}</td>
			</tr>
			<tr>
				<th>User</th>
				<td><a href="user-{{.User.Id}}.html">{{.User.Name}}</a></td>
			</tr>
			<tr>
				<th>Label</th>
				<td>
				{{range .Labels}}
					<span style="background-color: #{{.Color}}">{{.Name}}</span>
				{{end}}
				</td>
			</tr>
			<tr>
				<th>State</th><td>{{.State}}</td>
			</tr>
			<tr>
				<th>Milestone</th>
				<td>
				{{if ne 0 .Milestone.Number }}
					<a href="milestone-{{.Milestone.Id}}.html">
						{{.Milestone.Number}} {{.Milestone.Title}}
					</a>
					{{end}}
				</td>
			</tr>
			<tr>
				<th>Created At</th><td>{{.CreatedAt}}</td>
			</tr>
			<tr>
				<th>Updated At</th><td>{{.UpdatedAt}}</td>
			</tr>
			{{if .ClosedAt}}
			<tr>
				<th>ClosedAt</th><td>{{.ClosedAt}}</td>
			</tr>
			{{end}}
		</table>
		<pre>
{{.Body}}
		</pre>
	</body>
</html>
`

var IssuePageTemplate = template.Must(template.New("issuePage").
	Funcs(template.FuncMap{
		"extractRepoPath": extractRepoPath,
	}).
	Parse(IssuePageHTML))

// TODO user template
// TODO milestone template

func extractRepoPath(url string) string {
	elems := strings.Split(url, "/")
	if len(elems) < 2 {
		return ""
	}
	return elems[len(elems)-2] + "/" + elems[len(elems)-1]
}
