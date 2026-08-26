package parsereply
import (
	"encoding/xml"
	"camera-scanner/scan-network/internal/camera"
)

func parseStreamUriResponse(response []byte)(camera.StreamUri, error){
	var camUri camera.StreamUri
	err := xml.Unmarshal(response,&camUri)
	if err!=nil{
		return camera.StreamUri{},err
	}
	return camUri,nil
}