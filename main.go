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
	r := gin.Default()
	r.GET("/image", func(c *gin.Context) {

		file, err := os.Create("img.png")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		resp, err := http.Get(c.Query("img"))
		if err != nil {
			panic(err)
		}
		io.Copy(file, resp.Body)
		size, err := strconv.ParseInt(c.Query("size"), 10, 0)
		img, err := imaging.Open("./" + file.Name())
		dstimg := imaging.Resize(img, int(size), 0, imaging.Box)
		imaging.Save(dstimg, "img.resized.png")
		c.File("./img.resized.png")
	})

	r.Run(":8080")
}
