package utils

import (
	"encoding/base64"
	"github.com/skip2/go-qrcode"
)

func GenerateQRCode(data string) string {
	// 生成二维码图片，返回一个 *image.NRGBA 对象
	qrCodeImage, _ := qrcode.New(data, qrcode.Medium)

	// 将二维码图片保存为 PNG 图片，并将图片转换为 Base64 编码
	pngBase64 := encodeImageToBase64(qrCodeImage)

	// 将 Base64 编码的图片数据包装在 data URI 中，以便在 HTML 或 CSS 中使用
	pngDataURI := "data:image/png;base64," + pngBase64
	return pngDataURI
}

func encodeImageToBase64(img *qrcode.QRCode) string {
	// 将图片保存为 PNG 格式的字节数组
	pngByteArray, err := img.PNG(256)
	if err != nil {
		panic(err)
	}

	// 将字节数组转换为 Base64 编码
	pngBase64 := base64.StdEncoding.EncodeToString(pngByteArray)

	// 返回 Base64 编码的图片数据
	return pngBase64
}
