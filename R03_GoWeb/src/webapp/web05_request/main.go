package main

import (
	"R03_GoWeb/src/R00_webapp/model"
	json2 "encoding/json"
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
	fmt.Fprintln(w, "请求体中的内容有: ", string(body))
	// 解析表单, 在调用r.From之前必须执行该操作
	/*r.ParseForm()
	// 获取请求参数
	fmt.Fprintln(w, "请求参数有:", r.Form)*/
	// 通过直接调用FormVale方法和PostFormValue方法直接获取请求参数的值
	fmt.Fprintln(w, "URL中的user请求参数的值是: ", r.FormValue("user"))
	fmt.Fprintln(w, "Form表单中的username请求参数的值是: ", r.PostFormValue("username"))

}

func testJsonRes(w http.ResponseWriter, r *http.Request) {
	// 设置响应内容的类型
	w.Header().Set("Content-Type", "application/json")
	// 创建User
	user := model.User{
		ID:       1,
		Username: "admin",
		Password: "123456",
		Email:    "admin@qq.com",
	}
	// 将User转换为Json个数
	json, _ := json2.Marshal(user)
	// 将json格式的数据响应给客户端
	w.Write(json)
}

func restRedirect(w http.ResponseWriter, r *http.Request) {
	// 设置响应头中的Location属性
	w.Header().Set("Location", "https://www.baidu.com")
	// 设置响应的状态码
	w.WriteHeader(302)
}

func main() {
	http.HandleFunc("/hello", handler)
	http.HandleFunc("/testJsonRes", testJsonRes)
	http.HandleFunc("/restRedirect", restRedirect)
	http.ListenAndServe(":8080", nil)
}
