package main

import (
        "fmt"
	"flag"
        "net"
        "os"
)


func CheckError(err error) {
        if err != nil {
                fmt.Println("Error: ", err)
                os.Exit(1)
        }
}

func Response(sConn net.UDPConn, buf []byte, addr *net.UDPAddr, target string){
	//Request to Sserver
        ServerAddr, _ := net.ResolveUDPAddr("udp", target)
        LocalAddr, _ := net.ResolveUDPAddr("udp", ":0")
        Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	if err != nil{
                fmt.Println("Error: ", err)
		return
	}
        defer Conn.Close()

        _, err = Conn.Write(buf)
        buf = make([]byte, 1024)
        n, err := Conn.Read(buf)

	if err != nil{
                fmt.Println("Error: ", err)
		return
	}

	//Modify LI flag
        buf[0] = buf[0] & 077 // binary 0011 1111

	//Response to Client
        sConn.WriteToUDP(buf[0:n], addr)
}

func main() {
	const ntpPort = "123"
	var host string   = ""
	var verbose bool  = false

	flag.BoolVar(&verbose, "v", false, "Verbose")
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Println("You Specify NTP server")
		fmt.Println("ex) ./proxy example.com")
		os.Exit(0)
	}
	host = flag.Args()[0]

        ProxyAddr, err := net.ResolveUDPAddr("udp", ":" + ntpPort)
        CheckError(err)

        ServerConn, err := net.ListenUDP("udp", ProxyAddr)
        CheckError(err)
	fmt.Println("Info: Start Proxy on ", ProxyAddr, " to", host)
        defer ServerConn.Close()

        buf := make([]byte, 1024)
        for {
                //Recieve from client
                n, addr, err := ServerConn.ReadFromUDP(buf)
		if verbose {
                fmt.Println("Received from", addr)
		}
                if err != nil {
                        fmt.Println("Error: ", err)
			continue
                }
                go Response(*ServerConn, buf[0:n] , addr, host + ":" + ntpPort)
        }

}
