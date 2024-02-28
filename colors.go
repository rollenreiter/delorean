// coming soon
package main

import "runtime"

type St struct {
	Reset     string
	Bold      string
	Dim       string
	Underline string
	Blink     string
	Reverse   string
	Hidden    string
}

type Bg struct {
	Black string
	White string

	DarkGray  string
	LightGray string

	Red     string
	Green   string
	Yellow  string
	Blue    string
	Magenta string
	Cyan    string

	LightRed     string
	LightGreen   string
	LightYellow  string
	LightBlue    string
	LightMagenta string
	LightCyan    string
}

type Fg struct {
	Black string
	White string

	LightGray string
	DarkGray  string

	Red     string
	Green   string
	Yellow  string
	Blue    string
	Magenta string
	Cyan    string

	LightRed     string
	LightGreen   string
	LightYellow  string
	LightBlue    string
	LightMagenta string
	LightCyan    string
}

type Format struct {
	St
	Fg
	Bg
}

// GetColors returns a Formats type that has escape sequences which support the user's terminal and operating system
func InitColors() Format {
	var f Format
	if runtime.GOOS != "windows" {
		f = Format{
			St{
				"\033[0m",
				"\033[1m",
				"\033[2m",
				"\033[4m",
				"\033[5m",
				"\033[7m",
				"\033[8m",
			},
			Fg{
				"\033[30m",
				"\033[97m",
				"\033[90m",
				"\033[37m",
				"\033[31m",
				"\033[32m",
				"\033[33m",
				"\033[34m",
				"\033[35m",
				"\033[36m",
				"\033[91m",
				"\033[92m",
				"\033[93m",
				"\033[94m",
				"\033[95m",
				"\033[96m",
			},
			Bg{
				"\033[40m",
				"\033[107m",
				"\033[100m",
				"\033[47m",
				"\033[41m",
				"\033[42m",
				"\033[43m",
				"\033[44m",
				"\033[45m",
				"\033[46m",
				"\033[101m",
				"\033[102m",
				"\033[103m",
				"\033[104m",
				"\033[105m",
				"\033[106m",
			},
		}
	} else {
		f = Format{
			St{"", "", "", "", "", "", ""},
			Fg{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""},
			Bg{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""},
		}
	}
	return f
}
