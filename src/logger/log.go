package logger

import (
	"fmt"
	"os"
	"time"
)

var (
	greenBg      = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	whiteBg      = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellowBg     = string([]byte{27, 91, 57, 48, 59, 52, 51, 109})
	redBg        = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blueBg       = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	magentaBg    = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	cyanBg       = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	green        = string([]byte{27, 91, 51, 50, 109})
	white        = string([]byte{27, 91, 51, 55, 109})
	yellow       = string([]byte{27, 91, 51, 51, 109})
	red          = string([]byte{27, 91, 51, 49, 109})
	blue         = string([]byte{27, 91, 51, 52, 109})
	magenta      = string([]byte{27, 91, 51, 53, 109})
	cyan         = string([]byte{27, 91, 51, 54, 109})
	reset        = string([]byte{27, 91, 48, 109})
	disableColor = false
)

func Info(str string) {
	now := time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
	fmt.Print(now)
	fmt.Print(green, " Info :", reset)
	fmt.Println(str)
}

func Error(str string) {
	now := time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
	fmt.Print(now)
	fmt.Print(red, " Error:", reset)
	fmt.Println(str)
}

func Fatal(str string) {
	now := time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
	fmt.Print(now)
	fmt.Print(red, " Fatal:", reset)
	fmt.Println(str)
	os.Exit(1)
}
