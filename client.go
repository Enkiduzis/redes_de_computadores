package main
import "net"
import "os"

func main(){
	server, _ := net.Dial("tcp", "localhost:9986")
	server.Write([]byte(os.Args[1]))
	
}