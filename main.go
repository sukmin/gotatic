package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sukmin/gotatic/mymiddleware"
	"math/rand"
	"net"
	"net/http"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

//http://www.faviconr.com/
const favicon_base64 = `AAABAAEAEBAAAAEAIABoBAAAFgAAACgAAAAQAAAAIAAAAAEAIAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAATIT3/FiVC/yU2Vv85Unf/cYCQ/8vLxf/Awr3/i4qL/3Brd/9oXmf/fH+F/3Nvd/91g4r/eYSH/3p6c/9zdHD/EBw2/xYjP/8oP2L/WnWV/8LFxP/Y1tD/ycrE/5eamv9vaXL/aGNs/29nb/9vZWz/eISK/3d/f/98goH/eXly/xUlQf8cLUj/M0xw/6y0t//h3tf/393W/8W/uf+ShYj/b3B6/2Rufv9ubHL/fYSG/3+Lkf95hov/f4uQ/3yAfv8gN1v/KD9g/zlTc/++v7n/3tzU/9/e1//FurL/kICA/21vdv9MbIX/XV9m/21qcv+AipD/gY2T/4OPlP+CjpX/Hzli/y9Ld/83VoD/bYCT/9bTyv/f3dX/z83G/7Kzrf9yb3b/S1Vo/0BIV/85PUf/f4SC/5mjo/+YoaL/mKKi/x84Yf8yT3z/QmSS/z5fiP+FlKH/29jQ/9LSx/+zuLf/d4ql/2J0j/9gcY//R1Np/7O3r//Cxb3/xMa+/8THwf8eOGH/JEBt/ydFc/8tSXH/RWCB/3GInf93jqT/hZ62/3qQrf9wi6z/ZX+j/4CTpf/a2dH/2NnQ/9nZ0v/Z2tP/Hzpl/yQ/a/8iP23/KUZv/zxch/9GaJT/V3ad/36Vr/9lfZ3/Y4Sr/22Lrv/Ax8T/3dzU/9vb0v/b3NX/29zV/yFBcf8gPWv/I0Jx/ytKd/83Wor/PWCO/2Z/n/+vv83/dY6r/2WDp/95j6f/wsnI/9va0P/e3dT/2tvT/9ra0/8lRnb/I0V4/yxOgP8uUIH/LE59/0Jkk/+Lorr/qr3O/5Cnv/9nf6L/Y26A/73Hyv/e3ND/3tzS/9vb0//Z2dD/I0R0/yxSiv81WpH/LFGH/y9Vif8+YJD/iaW+/5evxf+Gobv/aoWl/1Zsi/+1wMT/393Q/9zVy//Z1c3/2tjP/yxShv86ZJz/OmGb/zRZkf81W4//RmaR/3WVtf+En7r/fJe0/3OLqP9ee5z/xMvJ/97XzP/Szsj/29LH/97cy/9BbKH/RHKp/0Boov86Ypn/PWKV/0xrlP9hgqn/YICn/2aGq/95k63/ZH6b/9PV0P/d1s3/2NHH/9XUyf++w8L/WoSx/1WBsP9IdKf/RG2h/1uArP+Bn7v/YYKq/1R1oP9VdaD/aIer/4+pv//h4Nf/497R/9vWz//Cyc7/rbbB/7+0rv+EjJ7/ip+5/4Gevf+KpsL/iKbC/2yIsP9Tcp7/Y4Oq/4Cduv+xv8f/5eDR/9zb1v/W2N7/yMzT/6Kstv/OxLz/sK+4/6+stf/Fx9H/usHO/6i4yv99mrv/dZe6/6q3vv/Nzsj/29fN/9jVzv/R0tb/0tXZ/8HGzP+lrrj/AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==
`

func main() {

	port := flag.String("p", "11007", "Http listen port. default 11007")
	basicAuth := flag.Bool("a", false, "Http basicAuth true/false. default false")
	flag.Parse()

	faviconBytes, _ := base64.StdEncoding.DecodeString(favicon_base64);
	workPath, _ := os.Getwd()

	fmt.Println("gotatic start")
	printLocalUrl(*port)
	fmt.Println("work path : " + workPath)

	e := echo.New()
	e.HideBanner = true

	e.Use(mymiddleware.Logrus())
	e.Use(mymiddleware.NoCache())

	if *basicAuth {
		getUsername := RandStringRunes(20)
		genPassword := RandStringRunes(21)
		fmt.Println("username : " + getUsername)
		fmt.Println("password : " + genPassword)
		e.Use(middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{
			Skipper: func(c echo.Context) bool {
				return c.Path() == "/favicon.ico"
			},
			Validator: func(username, password string, c echo.Context) (bool, error) {
				if username == getUsername && password == genPassword {
					return true, nil
				}
				return false, nil
			},
		}))
	}

	// 파비콘을 추가하였으나 FileServer에서 파비콘링크를 추가해주지 않기에 향후 추가해야 한다.
	e.GET("/favicon.ico", func(c echo.Context) error {
		return c.Blob(http.StatusOK, "image/x-icon", faviconBytes)
	})

	e.GET("/*", echo.WrapHandler(http.FileServer(http.Dir(workPath))))

	err := e.Start(":" + *port)
	if err != nil {
		panic(err)
	}

}

func RandStringRunes(n int) string {

	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func printLocalUrl(port string) {
	fmt.Println("access url list")
	format := "\thttp://%s:%s\n"
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Printf(format, "127.0.0.1", port)
		return
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.IsGlobalUnicast() {
			if ipnet.IP.To4() != nil {
				fmt.Printf(format, ipnet.IP.String(), port)
			}
		}
	}

	fmt.Printf(format, "127.0.0.1", port)

}
