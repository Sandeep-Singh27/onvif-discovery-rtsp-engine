package parsereply
import (
  "sync"
  "fmt"
  "camera-scanner/scan-network/internal/get-rtsp-url"
  "errors"
)

//HANDLE FUNCTION TO BE CALLED BY HANDLER FOR EVERY REPLY
func handle(cameraReply [2]string)error{

    camInfo,err := parseProbeReply(cameraReply[1])
    if err!=nil {
      fmt.Printf("[ERROR PROBEREPLY] %s",err)
      return err
    }

    //Assigning IP Address to CameraInfo 
    camInfo.IPAddr = cameraReply[0]

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

    //Initializing map for CameraInfo.RTSPURLs
    rtsp_urls := make(map[string]string)

    for _,v := range camProfiles.Profile{

      data,err := getrtspurl.RequestStreamURI(v.Token,camInfo)
      if err!=nil{
        return fmt.Errorf("Error while getting streamuri %s : %w", v.Token,err)
      }

      camUri, err := parseStreamUriResponse(data)
      if err!=nil{
        return fmt.Errorf("Error while parsing streamuri %s : %w", v.Token,err)
      }
      rtsp_urls[v.Token] = camUri.Uri
    }

    camInfo.RTSPUrls = rtsp_urls
    fmt.Println(camInfo)
    return nil
}

//HANDLER TO PARSE CHANNEL DATA
func Handler(wg *sync.WaitGroup,c chan [2]string){
  fmt.Printf("[INFO] Handler Started\n")
  defer wg.Done()
  defer fmt.Printf("[INFO] Handler Closed\n")
  for cameraReply := range c{
    err := handle(cameraReply)
    if err!=nil{
      fmt.Println("[ERROR HANDLING REPLY] ",err)
    }
  }
}

