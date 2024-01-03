package Tools

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetData(file string) {
	opened, err := os.Open(file)
	if err != nil {

	}

	var tab []string
	number, counter := 0, 0
	var rooms []string
	var links map[int][]string

	scann := bufio.NewScanner(opened)
	i:=0
	for scann.Scan() {
		if counter == 0 {
			number, err := strconv.Atoi(scann.Text())
			if err != nil || number < 1 {
				fmt.Println("ERROR: invalid data format, invalid number of Ants")
				os.Exit(0)
			}
			counter++
		}else {
			text := scann.Text()
			isroom, name := CorrectRoom(text)
			if isroom {
				rooms = append(rooms, name)
				continue
			}
			if !isroom && name != "" {
				fmt.Println(name)
				os.Exit(0)
			}
			islink, rel := Islink(text)
			if islink {
				if len(rel)==2 {
					links[i] = rel
					i++
					continue
				}else {
					fmt.Println(rel)
					os.Exit(0)
				}
			} 
		}
		
	}
}

func CorrectRoom(text string) (bool,string){
	tab := strings.Split(text, " ")
	if len(tab) != 3 {
		return false, ""
	}
	_ , err := strconv.Atoi(tab[1])
	_ , err = strconv.Atoi(tab[2])
	if err != nil {
		return false, "ERROR: invalid data format, bad coordonate"
	}
	if (len(tab[0]) != 0) && ((string(tab[0][0]) == "L" ) || (string(tab[0][0]) == "#")) {
		return false, "ERROR: invalid data format,"
	} 
	return true, tab[0]
}

func Islink(text string)(bool, []string) {
	tab := strings.Split(text, "-")
	if len(tab) == 2 {
		if tab[0] != tab[1] {
			return true, tab
		}
		return true,[]string{"ERROR: invalid data format"}
	}
	return false, []string{""}
}