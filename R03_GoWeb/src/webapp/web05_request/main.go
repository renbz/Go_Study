package main

import (
	"fmt"
	"net/http"
)

// 创建处理器函数
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "你发送的请求的请求地址是:", r.URL.Path)
	fmt.Fprintln(w, "你发送的请求的请求地址后的查询字符串是:", r.URL.RawQuery)
	fmt.Fprintln(w, "请求头的所有信息是	:", r.Header)
	fmt.Fprintln(w, "请求头的Accept-Encoding的信息是:", r.Header["Accept-Encoding"])
	fmt.Fprintln(w, "请求头的Accept-Encoding的属性值是:", r.Header.Get("Accept-Encoding"))
	// 获取请求体内容的长度
	len := r.ContentLength
	// 创建byte切片
	body := make([]byte, len)
	// 将请求体重的内容读到body中
	r.Body.Read(body)
	// 在浏览器中显示请求体的内容
	fmt.Fprintln(w, "请求体重的内容有: ", string(body))
}
func main() {
	http.HandleFunc("/hello", handler)
	http.ListenAndServe(":8080", nil)
}
