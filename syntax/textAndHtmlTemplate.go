package main

import (
	"fmt"
	htmlTemplate "html/template"
	"log"
	"os"
	"text/template"
	"time"
)

type User struct {
	Login bool
}

type Issues struct {
	Number int
	User
	Title     string
	CreatedAt time.Time
}

type Result struct {
	TotalCount int
	Items      []Issues
}

func main() {
	// ===================
	// text/template包   html/template包可以将程序变量的值带入到文本或者html模板中

	const temp1 = `{{.TotalCount}} issues:
	{{range .Items}} -------------------
	Number: {{.Number}}
	User: 	{{.User.Login}}
	Title: 	{{.Title | printf "%.64s"}}
	Age: 	{{.CreatedAt | daysAgo}} days
	{{end}}`

	report, err := template.New("report").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(temp1)
	if err != nil {
		log.Fatal()
	}
	fmt.Println(report)

	// 模版通常是在编译期间就固定下来的，因此无法解析模板将是程序中一个严重的bug
	// template.Must() 提供了一种比啊姐的错误处理方式，它接受一个模板和错误作为参数，
	// 检查错误是否为nil，如果不是nil，则宕机
	var report2 = template.Must(template.New("report2").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(temp1))
	fmt.Println(report2)

	var issues = Issues{1, User{true}, "test", time.Now()}
	var items = []Issues{issues, issues, issues}
	var result = Result{TotalCount: 3, Items: items}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}

	// html模板
	issuesHtml := htmlTemplate.Must(htmlTemplate.New("issuesList").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>
<h1>{{.TotalCount}}}</h1>
<table>
    <tr style="text-align: left">
        <th>#</th>
        <th>Number</th>
        <th>User</th>
        <th>Title</th>
    </tr>
    {{range .Items}}}
    <tr>
        <td>{{.Number}}</td>
        <td>{{.User.Login}}</td>
        <td>{{.Title}}</td>
    </tr>
    {{end}}
</table>
</body>
</html>
`))

	if err := issuesHtml.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}

}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
