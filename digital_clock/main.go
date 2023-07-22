package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

type placeholder [5]string

func main() {
	d := digits()
	c := colon()

	for {
		clearConsole()

		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		clock := [...]placeholder{
			d[hour/10], d[hour%10],
			c,
			d[min/10], d[min%10],
			c,
			d[sec/10], d[sec%10],
		}

		for line := 0; line < 5; line++ {
			for index, digit := range clock {
				next := clock[index][line]
				if digit == c && sec%2 == 0 {
					next = "   "
				}
				fmt.Print(next, "  ")
			}

			fmt.Println()
		}

		time.Sleep(time.Second)
	}
}

func digits() [10]placeholder {
	zero := placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one := placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two := placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three := placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four := placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five := placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six := placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven := placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight := placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine := placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	return [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}
}

func colon() placeholder {
	return placeholder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}
}

func clearConsole() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
