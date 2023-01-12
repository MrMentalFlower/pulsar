/*
letters
font: lary 3d from texttool.com
https://www.texttool.com/ascii-font#p=display&f=Larry%203D&t=%5C
 ______       ____         ____         ____         _____        ____       ____        __  __       ______       _____      __  __      __                       _   _      __  __      _____       ____        _____       ____        ____        ______      __  __      __  __      __   __      __    __      ________
/\  _  \     /\  _`\      /\  _`\      /\  _`\      /\  ___\     /\  _`\    /\  _`\     /\ \/\ \     /\__  _\     /\___ \    /\ \/\ \    /\ \         /'\_/`\     / \_/ \    /\ \/\ \    /\  __`\    /\  _`\     /\  __`\    /\  _`\     /\  _`\     /\__  _\    /\ \/\ \    /\ \/\ \    /\ \ /\ \    /\ \  /\ \    /\_____  \
\ \ \L\ \    \ \ \L\ \    \ \ \/\_\    \ \ \/\ \    \ \ \___     \ \ \L\_\  \ \ \L\_\   \ \ \_\ \    \/_/\ \/     \/__/\ \   \ \ \/'/'   \ \ \       /\      \   /\ \__\ \   \ \ `\\ \   \ \ \/\ \   \ \ \L\ \   \ \ \/\ \   \ \ \L\ \   \ \,\L\_\   \/_/\ \/    \ \ \ \ \   \ \ \ \ \   \ `\`\/'/'   \ `\`\\/'/    \/____//'/'
 \ \  __ \    \ \  _ <'    \ \ \/_/_    \ \ \ \ \    \ \  ___\    \ \  _\/   \ \ \L_L    \ \  _  \      \ \ \        _\ \ \   \ \ , <     \ \ \  __  \ \ \__\ \  \ \ \_/\ \   \ \ , ` \   \ \ \ \ \   \ \ ,__/    \ \ \ \ \   \ \ ,  /    \/_\__ \      \ \ \     \ \ \ \ \   \ \ \ \ \   `\/ > <      `\ `\ /'          //'/'
  \ \ \/\ \    \ \ \L\ \    \ \ \L\ \    \ \ \_\ \    \ \ \____    \ \ \/     \ \ \/, \   \ \ \ \ \      \_\ \__    /\ \_\ \   \ \ \\`\    \ \ \L\ \  \ \ \_/\ \  \ \ \\ \ \   \ \ \`\ \   \ \ \_\ \   \ \ \/      \ \ \_\ \   \ \ \\ \     /\ \L\ \     \ \ \     \ \ \_\ \   \ \ \_/ \     \/'/\`\     `\ \ \         //'/'___
   \ \_\ \_\    \ \____/     \ \____/     \ \____/     \ \_____\    \ \_\      \ \____/    \ \_\ \_\     /\_____\   \ \____/    \ \_\ \_\   \ \____/   \ \_\\ \_\  \ \_\\ \_\   \ \_\ \_\   \ \_____\   \ \_\       \ \_____\   \ \_\ \_\   \ `\____\     \ \_\     \ \_____\   \ `\___/     /\_\\ \_\     \ \_\        /\_______\
    \/_/\/_/     \/___/       \/___/       \/___/       \/_____/     \/_/       \/___/      \/_/\/_/     \/_____/    \/___/      \/_/\/_/    \/___/     \/_/ \/_/   \/_/ \/_/    \/_/\/_/    \/_____/    \/_/        \/_____/    \/_/\/ /    \/_____/      \/_/      \/_____/    `\/__/      \/_/ \/_/      \/_/        \/_______/
                                                                                                                                                                  ?

Numbers
   _          ___          __        __ __         ______        ____       ________      __          __           __        __
 /' \       /'___`\      /'__`\     /\ \\ \       /\  ___\      /'___\     /\_____  \   /'_ `\      /'_ `\       /'__`\     /\_\
/\_, \     /\_\ /\ \    /\_\L\ \    \ \ \\ \      \ \ \__/     /\ \__/     \/___//'/'  /\ \L\ \    /\ \L\ \     /\ \/\ \    \/_/      _______
\/_/\ \    \/_/// /__   \/_/_\_<_    \ \ \\ \_     \ \___``\   \ \  _``\       /' /'   \/_> _ <_   \ \___, \    \ \ \ \ \            /\______\
   \ \ \      // /_\ \    /\ \L\ \    \ \__ ,__\    \/\ \L\ \   \ \ \L\ \     /' /'      /\ \L\ \   \/__,/\ \    \ \ \_\ \      __   \/______/
    \ \_\    /\______/    \ \____/     \/_/\_\_/     \ \____/    \ \____/    /\_/        \ \____/        \ \_\    \ \____/     /\_\
     \/_/    \/_____/      \/___/         \/_/        \/___/      \/___/     \//          \/___/          \/_/     \/___/      \/_/



*/

package commands

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

// reference time for the clock to know when to update
var referenceTime string

// Compares the reference time to the current time to see if the clock needs to be update
func WriteOut(currentTime string) {
	if currentTime != referenceTime {
		CLS()
		fmt.Print(AsciiPrint([]rune(currentTime)))
	}
	referenceTime = currentTime
}

// Checks if the flags are set or not and display's a quick little tip for help before running the clock and sets the time
// defaults to 24 hour format time
func SetTime(help bool, setTime [2]*bool) string {
	if *setTime[0] {
		return time.Now().Format("03:04  PM")
	} else if *setTime[1] && help || !help {
		return time.Now().Format("15:04")
	}

	if help {
		fmt.Println("-h for commands")
		time.Sleep(1500000000)
		return time.Now().Format("15:04")
	} else {
		return time.Now().Format("15:04")
	}

}

// Code from MasterDimmy https://github.com/MasterDimmy/go-cls/blob/main/cls.go
func CLS() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("clear") //Linux example, its tested
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls") //Windows example, its tested
	default:
		fmt.Println("CLS for ", runtime.GOOS, " not implemented")
		return
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
