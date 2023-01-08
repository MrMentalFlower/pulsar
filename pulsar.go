/*
 ____        __  __      __          ____        ______      ____
/\  _`\     /\ \/\ \    /\ \        /\  _`\     /\  _  \    /\  _`\
\ \ \L\ \   \ \ \ \ \   \ \ \       \ \,\L\_\   \ \ \_\ \   \ \ \_\ \
 \ \ ,__/    \ \ \ \ \   \ \ \  __   \/_\__ \    \ \  __ \   \ \ ,  /
  \ \ \/      \ \ \_\ \   \ \ \L\ \    /\ \L\ \   \ \ \/\ \   \ \ \\ \
   \ \_\       \ \_____\   \ \____/    \ `\____\   \ \_\ \_\   \ \_\ \_\
    \/_/        \/_____/    \/___/      \/_____/    \/_/\/_/    \/_/\/ /

Date: 2023 1-5
Creator: flower
Description: A simple CLI clock that just displays time

*/

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

package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

// newline in ascii is 0x0a

func main() {
	var tmp, currentTime string
	var timeFormat, help bool

	flag.BoolVar(&timeFormat, "12", true, "Set 12 hour time")
	flag.BoolVar(&timeFormat, "24", false, "Set 24 hour time")
	flag.Parse()

	for {
		if timeFormat {
			currentTime = time.Now().Format("03:04")
		} else {
			currentTime = time.Now().Format("15:04")
		}
		if !help {
			fmt.Println("-h for commands")
			help = true
			time.Sleep(1500000000)
		}

		if tmp != currentTime {
			CLS()
			fmt.Print(asciiPrint([]rune(currentTime)))

		}
		tmp = currentTime

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

// len(numbers[0]
// need a way to save state after the new line is reached and then continue from ther
func asciiPrint(input []rune) string {
	var output string
	var character int
	var tmp [20]int

	for {
		for i := 0; i < len(input); i++ {
			character = tmp[i]
			for {
				if character == len(numbers()[input[i]]) {
					return output
				}

				if numbers()[input[i]][character] == 0x0a {
					tmp[i] = character + 1
					//output = output + " "
					break
				}
				output = output + string(numbers()[input[i]][character])
				character++
			}
		}
		output = output + "\n"

	}

	return output
}

func numbers() map[rune]string {
	var nums = map[rune]string{
		'0': `
  ___     
 / __ \   
/\ \/\ \  
\ \ \ \ \ 
 \ \ \_\ \
  \ \____/
   \/___/ 
`,

		'1': `
  __     
 /' \    
/\_, \   
\/_/\ \  
   \ \ \ 
    \ \_\
     \/_/
`,

		'2': `
   ___     
 /'___` + "`" + `\   
/\_\ /\ \  
\/_/// /__ 
   // /_\ \
  /\______/
  \/_____/ 
`,

		'3': `
   __     
 /'__` + "`" + `\   
/\_\L\ \  
\/_/_\_<_ 
  /\ \L\ \
  \ \____/
   \/___/ 
`,

		'4': `
 __ __      
/\ \\ \     
\ \ \\ \    
 \ \ \\ \_  
  \ \__ ,__\
   \/_/\_\_/
      \/_/  
`,

		'5': `
 ______    
/\  ___\   
\ \ \__/   
 \ \___` + "`" + `` + "`" + `\ 
  \/\ \L\ \
   \ \____/
    \/___/ 
`,

		'6': `
  ____    
 /'___\   
/\ \__/   
\ \  _` + "`" + `` + "`" + `\ 
 \ \ \L\ \
  \ \____/
   \/___/ 
`,

		'7': `
 ________ 
/\_____  \
\/___//'/'
    /' /' 
   /' /'  
  /\_/    
  \//     
`,

		'8': `
   __     
 /'_ ` + "`" + `\   
/\ \L\ \  
\/_> _ <_ 
  /\ \L\ \
  \ \____/
   \/___/ 
`,

		'9': `
   __      
 /'_  \    
/\ \L\ \   
\ \___, \  
 \/__,/\ \ 
      \ \_\
       \/_/
`,

		':': `
 __    
/\_\   
\/_/   
       
    __ 
   /\_\
   \/_/
`,
	}
	return nums
}
