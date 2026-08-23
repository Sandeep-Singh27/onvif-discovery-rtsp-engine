package parsereply
import (
  "camera-scanner/scan-network/internal/camera"
  "encoding/xml"
)

func parseProfilesResponse(profileResponse []byte)(camera.CameraProfiles,error){
  var camProfiles camera.CameraProfiles
  err := xml.Unmarshal(profileResponse,&camProfiles)
  if err!=nil{
    return camera.CameraProfiles{}, err
  }
  return camProfiles, nil
}

