package gorest

import (
    "net"
    "bufio"
    "fmt"
    "strings"
    "gocontrol"
)


func RunServer() {
	InitializeRoutes();
	server, err := net.Listen("tcp", ":80")
    if server == nil {
        panic("error listening: " + err.Error())
    }
    conns := clientConns(server)
    for {
        go handleRequest(<-conns)
    }
}


func clientConns(listener net.Listener) chan net.Conn {
    ch := make(chan net.Conn)
    go func() {
        for {
            client, err := listener.Accept()
            if client == nil {
                fmt.Printf("couldn't accept: " + err.Error())
                continue
            }
            fmt.Printf("Client connected from %v <-> %v\n", client.LocalAddr(), client.RemoteAddr())
            ch <- client
        }
    }()
    return ch
}

func handleRequest(client net.Conn) {

    requestData := make(map[string]string)

    b := bufio.NewReader(client)
    firstline, firsterr := b.ReadBytes('\n')

    if(firsterr != nil) {
        fmt.Println("Couldn't read input from client")
        return
    }

    hostRequest := string(firstline)
	requestParams := strings.Fields(hostRequest)

    controller := routes[requestParams[0] + " " + requestParams[1]]
    requestData["IP"] = client.RemoteAddr().String()
    for {
        line, err := b.ReadBytes('\n')
        if((line[0] == 13 && line[1] == 10) || err != nil) { // EOF or newline
            break
        }
        header := string(line)
        headerArray := strings.Split(header, ":")
        requestData[headerArray[0]] = strings.TrimSpace(headerArray[1])
    }
	
	if controller == "" {
    	render404(client)
    } else {
    	renderPage(controller, client, requestData)
	}
}

func render404(client net.Conn) {
	client.Write([]byte("HTTP/1.1 404 NOT FOUND\n"))
    client.Write([]byte("Server: Go! Server\n"))
    client.Write([]byte("Content-Type: text/html\n"))
    client.Write([]byte("Connection: keep-alive\n"))
    client.Write([]byte("\n"))
    client.Write([]byte("<html><h4>404 NOT FOUND</h4></html>"))
    client.Close()	
}

func renderPage(controller string, client net.Conn, headers map[string]string) {
	data := gocontrol.CallController(controller, headers)
	client.Write([]byte("HTTP/1.1 200 OK\n"))
    client.Write([]byte("Server: Go! Server\n"))
    client.Write([]byte("Content-Type: text/html\n"))
    client.Write([]byte("Connection: keep-alive\n"))
    client.Write([]byte("\n"))
    client.Write([]byte(data))
	client.Close()
}