package main

var IndexHtml = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>Choose Your Own Adventure</title>
	</head>
	<body>
		<u>
		{{range $k,$v := .Stories}}
		  <li> <a href="/story/{{$k}}">{{$k}}</a> </li>
		{{end}}
		</ul>
	</body>
</html>
<!DOCTYPE html>
`

var StoryHtml = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.ArcTitle}}</title>
	</head>
	<body>
		<h1>{{.Story.Title}}</h1>
		{{range .Story.Stories}}
		  <p>{{.}}</p>
		{{end}}
		<h2>options</h2>
		<ul>
			{{range .Story.Options}}
			<a href="{{if eq .Arc "home"}} /stories {{else}} /story/{{.Arc}} {{end}}">{{.Text}}</a><br>
			{{end}}
		</ul>
	</body>
</html>
<!DOCTYPE html>
`
