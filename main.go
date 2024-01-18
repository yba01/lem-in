package main

import (
	"fmt"
	"main/Tools"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	namefile:=os.Args[1]
	data := Tools.GetData(namefile)
	antnumber := Tools.AntNumber(data)
	allrooms, allinks := Tools.RoomAndLinks(data)
	start,end := Tools.StartAndEndRoom(allrooms)
	namerooms := Tools.NameRooms(allrooms)
	edge := Tools.LedgeBetween(allinks)
//	fmt.Println(data)
	fmt.Println(antnumber)
//	fmt.Println(allrooms)
//	fmt.Println(allinks)
//	fmt.Println(start)
//	fmt.Println(end)
//	fmt.Println(namerooms)
//  	fmt.Println(edge)
	graph := Tools.NewGraph()
	for _,node := range namerooms {
		graph.AddNode(node)
	}
	for _, link := range edge {
		graph.AddEdge(link[0], link[1])
	}
	graph.GetStartNode(start)
	graph.GetEndNode(end)

//	fmt.Println("Graph Nodes and Edges:")
//	for node, edges := range graph.Node {
//		fmt.Printf("Node %s -> Edges: %s\n", node, edges)
//	}
//	fmt.Printf("Start Room: %s\n", graph.StartNode)
//	fmt.Printf("End Room: %s\n", graph.EndNode)

	paths := graph.FindPathBfs()
	fmt.Println(paths)
	tab := make(map[int][][]string)
	i:=0
	for _, chem := range paths {
		chemins := [][]string{chem}
		for _, str := range paths {
			if !Tools.AvoidJams(chemins, str, graph.StartNode, graph.EndNode) {
				chemins = append(chemins, str)
			}
		}
		tab[i] = chemins
		i += 1
	}
		fmt.Println(tab)
		solution := Tools.Maxlenght(tab)
		
		fmt.Println(solution)
		antbypath := Tools.PutAntInPath(solution,antnumber)
		fmt.Println(antbypath)

}
