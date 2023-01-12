package commands

//prints the ascii art line by
func AsciiPrint(input []rune) string {
	var output string
	var character int
	var tmp [20]int

	for {
		for i := 0; i < len(input); i++ {
			character = tmp[i]
			for {
				if character == len(font()[input[i]]) {
					return output
				}

				if font()[input[i]][character] == 0x0a {
					tmp[i] = character + 1
					//output = output + " "
					break
				}
				output = output + string(font()[input[i]][character])
				character++
			}
		}
		output = output + "\n"
	}
}

func font() map[rune]string {
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

		'A': `
 ______     
/\  _  \    
\ \ \L\ \   
 \ \  __ \  
  \ \ \/\ \ 
   \ \_\ \_\
    \/_/\/_/
`,

		'M': `
  _   _     
 / \_/ \    
/\ \__\ \   
\ \ \_/\ \  
 \ \ \\ \ \ 
  \ \_\\ \_\
   \/_/ \/_/
`,

		'P': `
 ______  
/\  _ ` + "`" + `\ 
\ \ \L\ \
 \ \ ,__/
  \ \ \/ 
   \ \_\ 
    \/_/ 
`,

		' ': `
  
  
  
  
  
  
  
`,
	}
	return nums
}
