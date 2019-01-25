package main

import (
    "flag"
    "html/template"
    "log"
    "net/http"
)

// 127.0.0.1:1718
var addr = flag.String("addr", ":1718", "http service address") // Q=17, R=18

// 构建的HTML模版将会被服务器执行并显示在页面中。
var templ = template.Must(template.New("qr").Parse(templateStr))

func main() {
    flag.Parse()

    // 将 QR 函数绑定到服务器的根路径。
    http.Handle("/", http.HandlerFunc(QR))

    // 将在服务器运行时处于阻塞状态。
    err := http.ListenAndServe(*addr, nil)
    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}

// QR 仅接受包含表单数据的请求，并为表单值 s 中的数据执行模板。
func QR(w http.ResponseWriter, req *http.Request) {
    templ.Execute(w, req.FormValue("s"))
}

const templateStr = `
<html>
<head>
<title>QR Link Generator</title>
</head>
<body>
{{if .}}
<img src="http://chart.apis.google.com/chart?chs=300x300&cht=qr&choe=UTF-8&chl={{.}}" />
<br>
{{.}}
<br>
<br>
{{end}}
<form action="/" name=f method="GET"><input maxLength=1024 size=70
name=s value="" title="Text to QR Encode"><input type=submit
value="Show QR" name=qr>
</form>
</body>
</html>
`