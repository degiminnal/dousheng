package utils

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func SnapShotFromVideo(videoPath, snapShotPath string, frameNum int) (err error) {
	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return err
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return err
	}

	err = imaging.Save(img, snapShotPath)
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return err
	}
	return nil
}
