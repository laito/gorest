package gorest


import (
    "bufio"
    "fmt"
    "io"
    "log"
    "os"
    "strings"
    "strconv"
)

var routes = map[string]string{}

type route struct {
    url string
    controller string
}


func check(e error) {
	if e != nil {
		panic(e)
	}
}


func handleRoute(line int, route string) {
	validRoutes := [5]string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	route = strings.TrimSpace(route)
	if route > "" && !(strings.HasPrefix(route, "#"))	 {
		for _, request := range validRoutes {
			if strings.HasPrefix(route, request) {
				routeData := strings.Fields(route)
				routes[routeData[0]+" "+routeData[1]] = routeData[3]
				break
			} else {
				fmt.Println("Invalid Route on line " + strconv.Itoa(line) + " in go.routes. Ignoring.")
				break
			}
		}
	}
}

func InitializeRoutes() {
	file, err := os.Open("go.routes")
	check(err)
	bufferReader := bufio.NewReader(file)
	counter := 0
	for {
		switch line, err := bufferReader.ReadString('\n'); err {
		case nil: //valid line
			counter++
			handleRoute(counter, line)
		case io.EOF:
			counter++
			handleRoute(counter, line)
			return
		default:
			log.Fatal(err)
		}
	}	
}