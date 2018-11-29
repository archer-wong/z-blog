package routers

import (
	"github.com/go-macaron/cache"
	"github.com/go-macaron/captcha"
	"github.com/go-macaron/gzip"
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
	"html/template"
	"log"
	"net/http"
	"z-blog/helper"
	"z-blog/modules/setting"
	"z-blog/web/controllers"
)

var (
	m *macaron.Macaron
)

type M struct {
	*macaron.Macaron
}

func New() *M {
	new := func() *macaron.Macaron {
		initServer()
		initCaptcha()
		iniStatic()
		initTmpl()
		initSession()
		initRouter()
		return m
	}
	return &M{new()}
}

func (m *M) ListenAndServe() {
	env := setting.Cfg.Section("app").Key("ENV").String()
	address := setting.Cfg.Section("app").Key("ADDRESS").String()

	if env != "debug" {
		macaron.Env = "production"
	}
	log.Printf("Server is running on %s... (%s)\n", address, macaron.Env)
	http.ListenAndServe(address, m)
}

func initServer() {
	m = macaron.New()
	m.Use(macaron.Logger())
	m.Use(macaron.Recovery())
}

func initCaptcha() {
	m.Use(cache.Cacher())
	m.Use(captcha.Captchaer(captcha.Options{
		// 获取验证码图片的 URL 前缀，默认为 "/captcha/"
		URLPrefix:            "/captcha/",
		// 表单隐藏元素的 ID 名称，默认为 "captcha_id"
		FieldIdName:        "captcha_id",
		// 用户输入验证码值的元素 ID，默认为 "captcha"
		FieldCaptchaName:    "captcha",
		// 验证字符的个数，默认为 6
		ChallengeNums:        4,
		// 验证码图片的宽度，默认为 240 像素
		Width:                80,
		// 验证码图片的高度，默认为 80 像素
		Height:                35,
		// 验证码过期时间，默认为 600 秒
		Expiration:            600,
		// 用于存储验证码正确值的 Cache 键名，默认为 "captcha_"
		CachePrefix:        "captcha_",
	}))
}

func iniStatic() {
	m.Use(gzip.Gziper(gzip.Options{CompressionLevel: 2}))

	publicDir := setting.Cfg.Section("app").Key("PUBLIC_DIR").String()
	m.Use(macaron.Static(publicDir,
		macaron.StaticOptions{
			// 请求静态资源时的 URL 前缀，默认没有前缀
			Prefix: "public",
			// 禁止记录静态资源路由日志，默认为不禁止记录
			SkipLogging: true,
			// 当请求目录时的默认索引文件，默认为 "index.html"
		}))
}

func initTmpl() {
	viewDir := setting.Cfg.Section("app").Key("VIEW_DIR").String()
	m.Use(macaron.Renderer(macaron.RenderOptions{
		// 模板文件目录，默认为 "templates"
		Directory: viewDir,
		// 模板文件后缀，默认为 [".tmpl", ".html"]
		Extensions: []string{".tmpl", ".html"},
		// 模板函数，默认为 []
		Funcs: []template.FuncMap{map[string]interface{}{
			"AppName": func() string {
				return "Z-BLOG"
			},
			"AppVer": func() string {
				return "1.0.0"
			},
			"URL": m.URLFor,
			"DateTime": helper.DateTime,
			"SubString": helper.SubString,
			"Categories": controllers.Categories,
			"TopArticles": controllers.TopArticles,
			"Links": controllers.Links,
			"Breadcrumb": func(first, second string) []string {
				values := []string{first, second}
				return values
			},
		}},
		// 模板语法分隔符，默认为 ["{{", "}}"]
		Delims: macaron.Delims{"{{", "}}"},
		// 追加的 Content-Type 头信息，默认为 "UTF-8"
		Charset: "UTF-8",
		// 渲染具有缩进格式的 JSON，默认为不缩进
		IndentJSON: true,
		// 渲染具有缩进格式的 XML，默认为不缩进
		IndentXML: true,
		// 渲染具有前缀的 JSON，默认为无前缀
		HTMLContentType: "text/html",
	}))
}

func initSession() {
	m.Use(session.Sessioner(session.Options{CookieName: "zblog_id"}))
}
