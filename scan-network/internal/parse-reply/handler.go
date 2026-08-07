package parsereply
import (
  "sync"
  "fmt"
  "camera-scanner/scan-network/internal/get-rtsp-url"
)

//HANDLER TO PARSE CHANNEL DATA
func Handler(wg *sync.WaitGroup,c chan string){
  fmt.Printf("[INFO] Handler Started\n")
  defer wg.Done()
  defer fmt.Printf("[INFO] Handler Closed\n")
  for i := range c{
    camInfo,err := parseProbeReply(i)
    if err!=nil{
      fmt.Printf("[ERROR] %s",err)
    }else{
      _ ,err := getrtspurl.GetCapabilities(camInfo)
      if err==nil{
        fmt.Println("[INFO] Successfully retrieved capabilities.")
      }else{
        fmt.Println("[ERROR GETRTSP] ",err)
      } 
    } 
  }
}
