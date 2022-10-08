package utils

import (
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"strings"
	
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
	"github.com/spf13/viper"
)

func SaveImageFromRequest(c *gin.Context, httpRequestKey string) (string, error) {
	serverName := viper.GetString("server.name")

	avatarFile, handler, err := c.Request.FormFile(httpRequestKey)
	if err != nil {
		return "", err
	}
	defer avatarFile.Close()

	var img image.Image

	filenameParts := strings.Split(handler.Filename, ".")
	filenameExtension := strings.ToLower(filenameParts[len(filenameParts)-1])
	newFilename := uuid.NewV4().String()
	
	switch filenameExtension {
		case "jpg", "jpeg", "png":
			filenameExtension = "webp"
		case "ico","woff","swg","webp","webm","gif":
		default:
			return "", nil //make error later
	}

	options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 40)
	if err != nil {
		return "", err
	}

	resultFileName := newFilename + "." + filenameExtension
	output, err := os.Create(viper.GetString("img_path") + "/" + resultFileName)
	if err != nil {
		return "", err
	}
	defer output.Close()

	switch filenameExtension {
	case "jpg", "jpeg":
		img, err = jpeg.Decode(avatarFile)
		if err != nil {
			return "", err
		}
	case "png":
		img, err = png.Decode(avatarFile)
		if err != nil {
			return "", err
		}
	default:
		_, err := io.Copy(output, avatarFile)
		if err != nil {
			return "", err
		}
		return serverName + "/images/" + resultFileName, nil
	}

	if err := webp.Encode(output, img, options); err != nil {
		return "", err
	}
	return serverName + "/images/" + resultFileName, nil
}