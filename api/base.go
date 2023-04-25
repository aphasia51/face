package api

import (
	"face/dao"
	"face/forms"
	"face/scripts"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var (
	sms       = make(chan bool, 1)
	timeStamp = make(chan int, 1)
)

func DownLoad(ctx *gin.Context) {
	part, exists := <-timeStamp
	fmt.Println(part)
	fmt.Println(part, "paaaaaaaaaaaaaaaaaaaaart")
	if exists {
		filename := fmt.Sprintf("out_%v.mp4", part)
		ctx.File(fmt.Sprintf("./out_data/%v", filename))
	} else {
		zap.S().Info("文件下载地址生成失败")
		return
	}
}

func Notice(ctx *gin.Context) {
	client, _ := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

	for a := range sms {
		if a {
			err := client.WriteMessage(
				websocket.TextMessage,
				[]byte("视频剪切完成~, 请到链接下载"),
			)
			if err != nil {
				log.Println(err)
			}
			return
		}
	}
}

func OneCut(ctx *gin.Context) {
	var db = dao.DB
	settings := forms.Settings{}
	if err := ctx.ShouldBind(&settings); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "缺少剪切参数",
		})
		zap.S().Info("缺少剪切参数")
		sms <- false
		return
	}

	Stamp := time.Now().Unix()
	timeStamp <- int(Stamp)

	scripts.ExampleShowProgress(
		"./in_data/in1.mp4",
		fmt.Sprintf("./out_data/out_%v.mp4", Stamp),
		settings.Start,
		settings.Duration,
	)

	record := forms.Record{
		// Id:       1,
		UserName: ctx.GetString("username"),
		Start:    settings.Start,
		Duration: settings.Duration,
		Date:     time.Now().Format("2006-01-02 15:04:05"),
	}

	res := db.Create(&record)
	if res.Error != nil {
		zap.S().Info("用户记录存储失败!")
	}
	zap.S().Info("用户记录存储成功!")

	// err := ffmpeg.Input("./in_data/in1.mp4", ffmpeg.KwArgs{"ss": settings.Start}).
	// 	Output(fmt.Sprintf("./out_data/out_%v.mp4", timeStamp), ffmpeg.KwArgs{"t": settings.Duration}).
	// 	GlobalArgs("-progress").
	// 	OverWriteOutput().
	// 	Run()
	//
	// if err != nil {
	//
	// 	fmt.Println("Cut error")
	// 	sms <- false
	// 	return
	// }

	sms <- true

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "剪切完成，已通知您下载",
		"addr": fmt.Sprintf(
			"%v:%v/download",
			"localhost",
			dao.MySQLConfig.Port,
		),
	})
}
