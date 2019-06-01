package StringHelper

import (
	"image/jpeg"
	"zhenfangbian/web_api/utils/FileHelper"

	"github.com/astaxie/beego"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

/*
二维码
*/
type QrCode struct {
	URL    string
	Width  int
	Height int
	Ext    string
	Level  qr.ErrorCorrectionLevel
	Mode   qr.Encoding
}

const (
	EXT_JPG = ".jpg"
)

func NewQrCode(url string, width, height int, level qr.ErrorCorrectionLevel, mode qr.Encoding) *QrCode {
	return &QrCode{
		URL:    url,
		Width:  width,
		Height: height,
		Level:  level,
		Mode:   mode,
		Ext:    EXT_JPG,
	}
}

func GetQrCodePath() string {
	return beego.AppConfig.String("server_path::QrCodeSavePath")
}

func GetQrCodeFullPath() string {
	return beego.AppConfig.String("server_path::RuntimeRootPath")
}

func GetQrCodeFullUrl(name string) string {
	return beego.AppConfig.String("server_path::PrefixUrl") + GetQrCodePath() + name
}

func GetQrCodeFileName(value string) string {
	return Md5(value)
}

func (q *QrCode) GetQrCodeExt() string {
	return q.Ext
}

func (q *QrCode) CheckEncode(path string) bool {
	src := path + GetQrCodeFileName(q.URL) + q.GetQrCodeExt()
	if FileHelper.CheckNotExist(src) == true {
		return false
	}

	return true
}

/*
生成二维码
*/
func (q *QrCode) Encode(path string) (string, string, error) {
	name := GetQrCodeFileName(q.URL) + q.GetQrCodeExt()
	src := path + name
	if FileHelper.CheckNotExist(src) == true {
		code, err := qr.Encode(q.URL, q.Level, q.Mode)
		if err != nil {
			return "", "", err
		}

		code, err = barcode.Scale(code, q.Width, q.Height)
		if err != nil {
			return "", "", err
		}

		f, err := FileHelper.MustOpen(name, path)
		if err != nil {
			return "", "", err
		}
		defer f.Close()

		err = jpeg.Encode(f, code, nil)
		if err != nil {
			return "", "", err
		}
	}

	return name, path, nil
}
