package main

import (
	"flag"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	port := flag.String("p", "11007", "Http listen port. default 11007")
	basicAuth := flag.Bool("a", true, "Http basicAuth true/false. default true")
	directoryListing := flag.Bool("d", true, "Directory listing true/false. default true")
	flag.Parse()

	ip := LocalIp()
	workPath, _ := os.Getwd()

	fmt.Println("gotatic start")
	fmt.Println("access url : http://" + ip + ":" + *port)
	fmt.Println("work path : " + workPath)
	fmt.Println("directoryListing mode :" + strconv.FormatBool(*directoryListing))

	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.Logger())
	if *basicAuth {
		getUsername := RandStringRunes(20)
		genPassword := RandStringRunes(21)
		fmt.Println("username : " + getUsername)
		fmt.Println("password : " + genPassword)
		e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
			if username == getUsername && password == genPassword {
				return true, nil
			}
			return false, nil
		}))
	}

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   workPath,
		Browse: *directoryListing,
	}))

	e.Logger.Fatal(e.Start(":" + *port))

}

func RandStringRunes(n int) string {

	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func LocalIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.IsGlobalUnicast() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "127.0.0.1"
}
