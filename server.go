package main
import "fmt"
import "net"

func main(){
	connection, _ := net.Listen("tcp", "localhost:9986")
	for(connection != nil){
		fmt.Println("Aguardando Clients...")
		client,_ :=connection.Accept()
		buffer:= make([]byte,1024)
		mLen, _ := client.Read(buffer)
		fmt.Println("Recebido: ", string(buffer[:mLen]))
	}
}