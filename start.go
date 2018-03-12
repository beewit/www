package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	//"github.com/beewit/beekit/utils"
	"fmt"
	"net/http"

	"github.com/beewit/beekit/utils"
	"github.com/beewit/beekit/utils/uhttp"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CSRF())
	e.Use(middleware.Gzip())

	e.Static("/page", "page")
	e.Static("/static", "static")

	e.File("/", "page/index.html")
	e.File("8VHBgcXdwx.txt", "8VHBgcXdwx.txt")
	//微信域名验证码
	e.File("MP_verify_3Z6AKFClzM8nQt3q.txt", "page/MP_verify_3Z6AKFClzM8nQt3q.txt")
	e.File("/.well-known/pki-validation/fileauth.txt", "fileauth.txt")

	var hiveHost = "http://hive.9ee3.com"

	e.GET("/api/help/list", func(c echo.Context) error {
		t := c.FormValue("t")
		url := fmt.Sprintf("%s/api/help/list?t=%s", hiveHost, t)
		b, err := uhttp.PostForm(url, nil)
		if err != nil {
			return utils.ErrorNull(c, "获取失败")
		}
		rp := utils.ToResultParam(b)
		return c.JSON(http.StatusOK, rp)
	})
	e.GET("/api/help/get", func(c echo.Context) error {
		id := c.FormValue("id")
		if id == "" || !utils.IsValidNumber(id) {
			return utils.ErrorNull(c, "数据参数错误")
		}
		url := fmt.Sprintf("%s/api/help/get?id="+id, hiveHost)
		b, err := uhttp.PostForm(url, nil)
		if err != nil {
			return utils.ErrorNull(c, "获取失败")
		}
		rp := utils.ToResultParam(b)
		return c.JSON(http.StatusOK, rp)
	})

	utils.Open("http://127.0.0.1:8080")

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
