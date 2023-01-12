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
Version: 1.1
*/

package main

import (
	"flag"

	commands "github.com/MrMentalFlower/pulsar/pkgs"
)

// where the code
func main() {
	var help bool = true
	var timeMode [2]*bool

	timeMode[0] = flag.Bool("12", false, "sets time to 12 hour format")
	timeMode[1] = flag.Bool("24", false, "sets time to 24 hour format")
	flag.Parse()
	commands.WriteOut(commands.SetTime(help, timeMode))
	help = false

	for {

		commands.WriteOut(commands.SetTime(help, timeMode))
	}
}
