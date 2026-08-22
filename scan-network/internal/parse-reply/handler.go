package parsereply
import (
  "sync"
  "fmt"
  "camera-scanner/scan-network/internal/get-rtsp-url"
)

//HANDLE FUNCTION TO BE CALLED BY HANDLER FOR EVERY REPLY
func handle(cameraReply string)error{

    camInfo,err := parseProbeReply(cameraReply)
    if err!=nil {
      fmt.Printf("[ERROR] %s",err)
      return err
    }

    capabilities ,err := getrtspurl.GetCapabilities(camInfo)
    if err!=nil {
      fmt.Println("[ERROR GETCAP] ",err)
      return err
    }

    mediaXAddrs,err := parseCapabilityResponse(capabilities)
    if err!=nil {
      fmt.Println("[ERROR PARSECAPRES] ",err)
    } 
    
    //Checking if mediaXAddrs got any media service XAddr
    fmt.Println("Media Length: ",len(mediaXAddrs.Media))
    return nil
}

//HANDLER TO PARSE CHANNEL DATA
func Handler(wg *sync.WaitGroup,c chan string){
  fmt.Printf("[INFO] Handler Started\n")
  defer wg.Done()
  defer fmt.Printf("[INFO] Handler Closed\n")
  for cameraReply := range c{
    handle(cameraReply)
  }
}

