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
    return camInfo,nil
  }
}
