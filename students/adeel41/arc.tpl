<html>
    <head>
        <title>{{.Title}}</title>
        <style>
            .arcLink { margin-bottom: 10px; display: block; }
        </style>
    </head>
    <body>
        <h2>{{.Title}}</h2>
        <p>{{.Paragraph}}</p>
        {{range .Options}}
        <a class="arcLink" href="/?arc={{.Arc}}">{{.Text}}</a>
        {{end}}
    </body>
</html>