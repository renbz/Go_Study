package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"log"
	"net/http"
)

func main() {
	// 创建一个服务
	ginServer := gin.Default()
	ginServer.Use(favicon.New("./favicon.ico"))
	// 加载静态页面
	ginServer.LoadHTMLGlob("template/*")
	// 响应一个页面给前端
	ginServer.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"msg": "这是Go后台传递来的数据",
		})
	})

	// 接收前端传过来的参数
	// url?userid=xxx&username=kuangshen
	ginServer.GET("/user/info", myHandler(), func(context *gin.Context) {
		userSession := context.MustGet("userSession").(string)
		log.Println("============》", userSession)
		userid := context.Query("userid")
		username := context.Query("username")
		context.JSON(http.StatusOK, gin.H{
			"userid":   userid,
			"username": username,
		})
	})

	//  /user/info/1/kuangshen
	ginServer.GET("/user/info/:userid/:username", func(context *gin.Context) {
		userid := context.Param("userid")
		username := context.Param("username")
		context.JSON(http.StatusOK, gin.H{
			"userid":   userid,
			"username": username,
		})
	})
	// 前端给后端传JSON
	ginServer.POST("/json", func(context *gin.Context) {
		b, _ := context.GetRawData()

		var m map[string]interface{}
		// 包装为json数据 []byte
		_ = json.Unmarshal(b, &m)
		context.JSON(http.StatusOK, m)

	})

	// 路由
	ginServer.GET("/test", func(context *gin.Context) {
		// 重定向301
		context.Redirect(http.StatusMovedPermanently, "https://www.kuangstudy.com")
	})
	// 路由组
	userGroup := ginServer.Group("/user")
	{
		userGroup.GET("/add")
		userGroup.GET("/login")
		userGroup.GET("/logout")
	}

	// 访问地址,处理我们的请求 request Response
	ginServer.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "hello world"})
	})

	ginServer.POST("/user", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "post user"})
	})
	// 服务器端口
	ginServer.Run(":8082")
}

// 自定义Go中间件、拦截器
func myHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 通过自定义中间件设置的值,在后续处理只要调用了这个中间件的都可以拿到这里的参数
		context.Set("userSession", "userid-1")
		context.Next() // 放行
		// context.Abort()  阻止
	}
}
