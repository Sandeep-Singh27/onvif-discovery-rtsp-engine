package parsereply
import (
  "sync"
  "fmt"
  "camera-scanner/scan-network/internal/get-rtsp-url"
  "errors"
)

//HANDLE FUNCTION TO BE CALLED BY HANDLER FOR EVERY REPLY
func handle(cameraReply string)error{

    camInfo,err := parseProbeReply(cameraReply)
    if err!=nil {
      fmt.Printf("[ERROR PROBEREPLY] %s",err)
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
      return err
    } 
    fmt.Println("[*] MediaXAddrs: ",mediaXAddrs)

    if len(mediaXAddrs.Media) < 1{
      fmt.Println("[ERROR] No MediaXAddr found after GetCapabilities")
      return errors.New("Found no MediaXAddr after GetCapabilities")
    }
  
    profiles ,err := getrtspurl.GetProfiles(mediaXAddrs.Media[0].XAddrs)
    if err!=nil{
      fmt.Println("[ERROR GETPROFILES] ",err)
      return err
    }
    
    camProfiles,err := parseProfilesResponse(profiles)
    if err!=nil {
      fmt.Println("[ERROR PARSEPROFRES] ",err)
      return err
    }
    fmt.Println("[*] Profile Tokens: ",camProfiles)
    data,err := getrtspurl.RequestStreamURI(camProfiles.Profile[0].Token,camInfo)
    fmt.Println(string(data))

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

