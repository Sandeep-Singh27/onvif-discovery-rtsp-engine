package parsereply
import (
  "fmt"
  "encoding/xml"
  "camera-scanner/scan-network/internal/camera"
)

//TO CONVERT XML TO PROPER FORMAT
func parseProbeReply(cameraReply string)(camera.CameraInfo, error){
  var camInfo camera.CameraInfo
  xml_byte := []byte(cameraReply)

  err := xml.Unmarshal(xml_byte, &camInfo)  

  if err!=nil{
    fmt.Println("[Error] ",err)
    return camera.CameraInfo{},err
  } else {
    //Resolving Scopes to array
    camInfo.ResolveScopes()
    fmt.Printf("Type: %s\n",camInfo.Type)
    fmt.Printf("XAddrs: %s\n",camInfo.XAddrs)
    fmt.Printf("ScopesArr: %s\n\n",camInfo.ScopesArr)
    return camInfo,nil
  }
}
