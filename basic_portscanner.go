package main
import (
	"net"
	"fmt"
)
func main(){

	conn, err := net.Dial("tcp", "anantshri.info:443")
	if err != nil {
		fmt.Printf("Port Closed")
	} else {
		fmt.Printf("Port Open %s\n", conn.RemoteAddr().String())
	}
    defer conn.Close()
	

}
