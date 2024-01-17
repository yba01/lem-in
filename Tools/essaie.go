package Tools

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)
// 1-
// Get data from file and put it in string's slice
func GetData(path string) []string {
	var tab []string
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error Read file :", err)
		os.Exit(0)
	}
	cont := bufio.NewScanner(file)
	for cont.Scan() {
		tab = append(tab, cont.Text())
	}
	return tab
}
func AntNumber(tab []string) int {
	if len(tab) > 1 {
		number, err := strconv.Atoi(tab[0])
		if err != nil {
			fmt.Println("Error ant number :",err)
			os.Exit(0)
		}
		return number
	}
	return 0
}
func RoomAndLinks(tab []string) ([]string,[]string) {
	var index int
	for i, str := range tab {
		count := len(strings.Split(str, "-"))
		if count == 2 {
			index = i
			break
		}
	}
	return tab[1:index], tab[index:]
}
func StartAndEndRoom(tab []string) (string, string) {
	var start []string
	var end []string
	for i, str := range tab {
		if str == "##start" {
			start = strings.Split(tab[i+1], " ")
		}
		if str == "##end" {
			end = strings.Split(tab[i+1]," ")
		}
	}
	return start[0], end[0]
}
func NameRooms(tab []string) []string {
	var namerooms []string
	for _, str := range tab {
		value := strings.Split(str," ")
		if len(value) == 3 {
			namerooms = append(namerooms, value[0])
		}
	}
	return namerooms
}
func LedgeBetween(tab []string) map[int][]string {
	j:=0
	links := make(map[int][]string)
	for i:=0; i<len(tab);i++ {
		value := strings.Split(tab[i], "-")
		if len(value)==2 {
			links[j] = value
			j++
		}
	}
	return links
}
//2-Create graph
type Graph struct {
	Node 		map[string][]string
	StartNode 	string
	EndNode 	string
}

func NewGraph() *Graph {
	return &Graph{
		Node: make(map[string][]string),
	}
}

func (g *Graph) AddNode(node string){
	if _, exist := g.Node[node]; !exist {
		g.Node[node]=[]string{}
	}else {
		fmt.Println("invalid data format, duplicated room")
		os.Exit(0)
	}
}
func (g *Graph) AddEdge(from, to string) {
	if _, exist :=g.Node[from]; !exist {
		fmt.Println("invalid data format, unknown room")
		os.Exit(0)
	}
	if _, exist :=g.Node[to]; !exist {
		fmt.Println("invalid data format, unknown room")
		os.Exit(0)
	}

	g.Node[from] = append(g.Node[from], to)
	g.Node[to] = append(g.Node[to], from)
}

func (g *Graph) GetStartNode(node string) {
	if _, exist := g.Node[node]; exist {
		g.StartNode = node
	}else {
		fmt.Println("invalid data format, unknown room")
		os.Exit(0)
	}
}
func (g *Graph) GetEndNode(node string) {
	if _, exist := g.Node[node]; exist {
		g.EndNode = node
	}else {
		fmt.Println("invalid data format, unknown room")
		os.Exit(0)
	}
}
// 3-Find path using Beardth First Search Algorithm
func (g *Graph) FindPathBfs() [][]string {
	queue := [][]string{{g.StartNode}}
	Allpath := [][]string(nil)
	//parents := []string(nil)

	for len(queue)>0 {
		path := queue[0]
		queue = queue[1:]
		node := path[len(path)-1]
		//fmt.Println("this is path",path)
		//fmt.Println("this is parents list",parents)

		if node == g.EndNode {
			Allpath = append(Allpath, path)
		}
		
		for _, 	adjacent := range g.Node[node] {
			if (!Containt(path, adjacent)){
				Newpath := append([]string(nil),path...)
				Newpath = append(Newpath, adjacent)
				queue = append(queue, Newpath)
			}
		}
	}	
	return Allpath	
} 
func Containt(path []string, node string) bool {
	for _, element := range path {
		if element == node {
			return true
		}
	}
	return false
}
// Allows to know if two paths have same rooms...
func SameRoom(one, second []string, start, end string) bool {
	for _, pathone := range one {
		if pathone == start {
			continue
		}else if pathone == end {
			return false
		}else {
			for _, pathsec := range second {
				if pathone == pathsec {
					return true
				}
			}
		}
	}
	return false
}
// allows to know if a path 
func AvoidJams(pathget [][]string, path []string, start, end string) bool {
	for _, chemin := range pathget {
		if SameRoom(chemin, path, start, end) {
			return true
		}
	}
	return false
}
// Our path
func Maxlenght(tab map[int][][]string) [][]string {
	if len(tab) == 1 {
		return tab[0]
	}
	max := len(tab[0])
	index := 0
	for i,value := range tab {
		if len(value) > max {
			index = i
		}
	}
	solution := tab[index]
	for i:=0; i<len(solution); i++ {
		for j:=0; j<len(solution)-1;j++ {
			if len(solution[j+1])<len(solution[j]) {
				solution[j+1], solution[j] = solution[j], solution[j+1]
			}
		}
	}
	return solution
}

func PutAntInPath(tab [][]string, nant int) map[int][]int {
	rslt := make(map[int][]int)
	index := 0
	for i:= 1; i <= nant; i++ {
		if i==1 {
			rslt[index] = append(rslt[index], i)
		}else{
			value := len(tab[index])-2 + len(rslt[index])
			if index != len(tab)-1 {
				if len(rslt)-1 <  index + 1 {
					if value > len(tab[index + 1])-2 {
						rslt[index+1] = append(rslt[index+1], i)
						index += 1
					}else {
						index = 0
						rslt[index] = append(rslt[index], i)
					}
				}else {
					otherval := len(tab[index+1]) - 2 + len(rslt[index+1])
					if value > otherval {
						rslt[index+1] = append(rslt[index+1], i)
						index += 1
					}else {
						index = 0
						rslt[index] = append(rslt[index], i)
					}
				}
			}else {
				index = 0
				rslt[index] = append(rslt[index], i)
			}
		}
	}		
	return rslt
}