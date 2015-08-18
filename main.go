package main

import (
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strconv"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r := gin.Default()
	r.GET("/image", handleResize)
	r.Run(":" + port)
}

func handleResize(c *gin.Context) {
	size, err := strconv.ParseInt(c.Query("size"), 10, 0)

	resp, err := http.Get(c.Query("img"))
	if err != nil {
		panic(err)
	}

	imgName := createImageFromResponse(resp)
	resizedImg := resizeImage(imgName, int(size))
	c.File(resizedImg)
}

func createImageFromResponse(data *http.Response) string {
	filename := "img"
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.Copy(file, data.Body)
	return filename
}

func resizeImage(filename string, size int) string {
	resizedFile := filename + ".resized.png"
	img, err := imaging.Open(filename)
	if err != nil {
		panic(err)
	}
	dstimg := imaging.Resize(img, size, 0, imaging.Box)
	imaging.Save(dstimg, resizedFile)
	return resizedFile
}
