package lib

import "fmt"

/*
The function Printcolor is use to replace fmt.Print and its print whit color we define in this func
For that we used the ANSI Escape Sequences
*/
func PrintColor(s, color string) {
	color_map := map[string]string{
		"Black":  "\033[0;30m",
		"Red":    "\033[0;31m",
		"Green":  "\033[0;32m",
		"Yellow": "\033[0;33m",
		"Blue":   "\033[0;34m",
		"Purple": "\033[0;35m",
		"Cyan":   "\033[0;36m",
		"White":  "\033[0;37m"}
	fmt.Printf("%v%v", color_map[color], s)
	fmt.Printf("%v", color_map["White"])
}
