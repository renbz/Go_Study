package main

import (
	"html/template"
	"net/http"
)

func testTemplate(w http.ResponseWriter, r *http.Request) {
	// 解析模板文件
	//t, _ := template.ParseFiles("index.html")
	t2 := template.Must(template.ParseFiles("index.html", "index2.html"))
	// 执行
	//t.Execute(w, "")
	t2.ExecuteTemplate(w, "index2.html", "我要去html2")
}

func main() {
	http.HandleFunc("/testTemplate", testTemplate)
	http.ListenAndServe(":8080", nil)
}
