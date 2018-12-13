package admin

import (
	"context"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"gopkg.in/macaron.v1"
	"io"
	"os"
	"time"
	"z-blog/modules/setting"
)

func QiniuUpload(ctx *macaron.Context) {
	ctx.Req.ParseMultipartForm(32 << 20)
	file, handler, err := ctx.Req.FormFile("editormd-image-file")

	if err != nil {
		ctx.JSON(200, map[string]interface{}{
			"success": 0,
			"message": "上传图片到服务器失败！",
		})
		return
	}
	defer file.Close()

	originName := handler.Filename
	fileName := time.Now().Format("2006-01-02-15-04-05") + "_" + originName

	f, err := os.OpenFile("storage/qiniu/"+fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		ctx.JSON(200, map[string]interface{}{
			"success": 0,
			"message": "上传图片到服务器失败！",
		})
		return
	}
	defer f.Close()
	io.Copy(f, file)

	localFile := "storage/qiniu/" + fileName
	sec := setting.Cfg.Section("qiniu")
	bucket := sec.Key("BUCKET").String()
	accessKey := sec.Key("ACCESS_KEY").String()
	secretKey := sec.Key("SECRET_KEY").String()
	zone := sec.Key("ZONE").String()
	useHttps, _ := sec.Key("USE_HTTPS").Bool()
	useCdnDomains, _ := sec.Key("USE_CDN_DOMAINS").Bool()
	domain := sec.Key("DOMAIN").String()

	putPolicy := storage.PutPolicy{
		Scope:               bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)

	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房

	if zone == "storage.ZoneHuadong" {
		cfg.Zone = &storage.ZoneHuadong
	} else if zone == "storage.ZoneHuabei" {
		cfg.Zone = &storage.ZoneHuabei
	} else if zone == "storage.ZoneHuanan" {
		cfg.Zone = &storage.ZoneHuanan
	} else if zone == "storage.ZoneBeimei" {
		cfg.Zone = &storage.ZoneBeimei
	} else {
		cfg.Zone = &storage.ZoneHuadong
	}
	// 是否使用https域名
	cfg.UseHTTPS = useHttps
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = useCdnDomains
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}

	err = formUploader.PutFile(context.Background(), &ret, upToken, fileName, localFile, &putExtra)
	if err != nil {
		ctx.JSON(200, map[string]interface{}{
			"success": 0,
			"message": "上传图片到七牛失败！",
		})
		return
	}

	ctx.JSON(200, map[string]interface{}{
		"success": 1,
		"message": "上传图片到服务器成功！",
		//"url": "http://markdown.archerwong.cn/" + fileName,
		"url": domain + ret.Key,
	})

}

