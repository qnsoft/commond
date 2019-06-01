package FileHelper

import (
	"image"
	"image/draw"
	"image/jpeg"
	"os"
	"zhenfangbian/web_api/utils/ErrorHelper"
)

/*
生成图片与图片层叠合成图
*/
func Pic_layer(Bg_pic, Ft_pic, New_pic string) string {
	//图片，网上随便找了一张
	img_file, err := os.Open(Bg_pic) //背景图
	ErrorHelper.CheckErr(err)        //打开图片出错
	defer img_file.Close()
	img, err := jpeg.Decode(img_file)
	ErrorHelper.CheckErr(err) //把图片解码为结构体时出错
	//水印,用的是我自己支付宝的二维码
	wmb_file, err := os.Open(Ft_pic) //前景图或水印图
	ErrorHelper.CheckErr(err)        //打开水印图片出错
	defer wmb_file.Close()
	wmb_img, err := jpeg.Decode(wmb_file)
	ErrorHelper.CheckErr(err) //把水印图片解码为结构体时出错
	//把水印写在右下角，并向0坐标偏移10个像素
	//offset := image.Pt(img.Bounds().Dx()-wmb_img.Bounds().Dx()-10, img.Bounds().Dy()-wmb_img.Bounds().Dy()-10)
	offset := image.Pt(wmb_img.Bounds().Dx()+90, img.Bounds().Dy()-1330)
	b := img.Bounds()
	//根据b画布的大小新建一个新图像
	m := image.NewRGBA(b)
	//image.ZP代表Point结构体，目标的源点，即(0,0)
	//draw.Src源图像透过遮罩后，替换掉目标图像
	//draw.Over源图像透过遮罩后，覆盖在目标图像上（类似图层）
	draw.Draw(m, b, img, image.ZP, draw.Src)
	draw.Draw(m, wmb_img.Bounds().Add(offset), wmb_img, image.ZP, draw.Over)
	//生成新图片new.jpg,并设置图片质量
	imgw, err := os.Create(New_pic)
	jpeg.Encode(imgw, m, &jpeg.Options{100})
	defer imgw.Close()
	return New_pic
}
