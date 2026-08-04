package discovery
import (
  "fmt"
  "net"
  "time"
  "errors"
  "sync"
  "camera-scanner/scan-network/internal/parse-camera-reply"
)

const (
  multicastAddress string = "239.255.255.250"
  multicastPort string = "3702"

  probeXML string = `
    <?xml version="1.0" encoding="UTF-8"?>

<e:Envelope
    xmlns:e="http://www.w3.org/2003/05/soap-envelope"
    xmlns:w="http://schemas.xmlsoap.org/ws/2004/08/addressing"
    xmlns:d="http://schemas.xmlsoap.org/ws/2005/04/discovery"
    xmlns:dn="http://www.onvif.org/ver10/network/wsdl">

    <e:Header>

        <w:MessageID>
            uuid:550e8400-e29b-41d4-a716-446655440000
        </w:MessageID>

        <w:To>
            urn:schemas-xmlsoap-org:ws:2005:04:discovery
        </w:To>

        <w:Action>
            http://schemas.xmlsoap.org/ws/2005/04/discovery/Probe
        </w:Action>

    </e:Header>

    <e:Body>

        <d:Probe>

            <d:Types>
                dn:NetworkVideoTransmitter
            </d:Types>

        </d:Probe>

    </e:Body>

</e:Envelope>
`
)

//LISTENING FOR THE REPLY
func listenConn(conn *net.UDPConn, wg *sync.WaitGroup, c chan string){
  fmt.Printf("[INFO] Listener Started\n\n")
  
  data_received := make([]byte,8192)

  defer wg.Done()
  defer close(c)
  defer fmt.Println("[INFO] Listener Closed.")
  
  for {
    n,addr,err := conn.ReadFrom(data_received)
    if err!=nil{
      if errors.Is(err,net.ErrClosed){
        fmt.Println("[INFO] Connection has been closed")
        break
      } else {
        fmt.Println("[ERROR] ",err)
      }
    } else {
      fmt.Println(addr)
      c <- fmt.Sprintf("%s",data_received[:n])
    }
  }
}

// SENDING PROBE MESSAGE AND DISCOVERING CAMERAS
func DiscoverCamera(){

  // Resolving address to listen on 
  laddr, err := net.ResolveUDPAddr("udp","0.0.0.0:7890")
  if err!=nil{
    fmt.Println("[ERROR] Error while resolving the address")
    panic(err)
  }

  // Making a UDP Listener
  conn, err := net.ListenUDP("udp", laddr)
  fmt.Printf("[INFO] listening on %s\n",conn.LocalAddr())
 

  // sending discovery message to the multicast address
  if err!=nil{

    panic(err)

  } else {

    //Resolving Multicast Address
    maddr,err := net.ResolveUDPAddr("udp",multicastAddress+":"+multicastPort)
    if err!=nil{
      fmt.Println("[Error] while resolving multicast address")
      panic(err)
    }

    // Sending the Discovery Message
    discoveryMessage := []byte(probeXML)
    _, err = conn.WriteTo(discoveryMessage,maddr)

    if err!=nil{
      fmt.Println("[ERROR] Something went wrong while sending discovery message")
      panic(err)
    } else {
      fmt.Printf("[INFO] Discovery message sent successfully to %s\n",maddr)
    }

    //making channel to send reply from listener to handler
    c := make(chan string)
    
    //Creating listener goroutine
    wg := &sync.WaitGroup{}
    wg.Add(2)
    go listenConn(conn,wg,c)
    go parsereply.Handler(wg,c)

    //Waiting 5 secs for the reply on our socket
    time.Sleep(10 * time.Second)
    
    //Closing the socket after waiting
    conn.Close()

    //Waiting for the goroutine to complete
    wg.Wait()
  }
}
