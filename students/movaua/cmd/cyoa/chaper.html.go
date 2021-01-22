package main

const (
	chapterHTML = `
<!DOCTYPE html>
<html>

<head>
  <meta charset='utf-8'>
  <title>CYOA - {{ .Title }}</title>
  <meta name='viewport' content='width=device-width, initial-scale=1'>
</head>

<body>

  <h2>{{ .Title }}</h2>

  <p>{{ .Story }}</p>

  <ul>
    {{range .Options }}
    <li><a href="/{{ .Chaper }}">{{ .Text }}</a></li>
    {{end}}
  </ul>

</body>

</html>`
)
