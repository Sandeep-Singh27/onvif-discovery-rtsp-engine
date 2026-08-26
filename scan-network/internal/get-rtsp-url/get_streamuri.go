package getrtspurl

import (
    "io"
    "bytes"
	"encoding/xml"
    "net/http"
    "camera-scanner/scan-network/internal/camera"
)

type Envelope struct {
	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope Envelope"`

	Body Body `xml:"http://www.w3.org/2003/05/soap-envelope Body"`
}

type Body struct {
	GetStreamURI GetStreamURI `xml:"GetStreamUri"`
}

type GetStreamURI struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl GetStreamUri"`

	StreamSetup StreamSetup `xml:"StreamSetup"`
	ProfileToken string      `xml:"ProfileToken"`
}

type StreamSetup struct {
	Stream string    `xml:"Stream"`
	Transport Transport `xml:"Transport"`
}

type Transport struct {
	Protocol string `xml:"Protocol"`
}

//FUNCTION TO REQUEST RTSP URL OF CAMERA
func RequestStreamURI(token string ,camInfo camera.CameraInfo)([]byte, error){

  request := Envelope{
    Body: Body{
      GetStreamURI: GetStreamURI{
        StreamSetup: StreamSetup{
          Stream: "RTP-Unicast",
		  Transport: Transport{
            Protocol: "RTSP",
          },
		},
		ProfileToken: token,
	  },
    },
  } 

  request_xml, err := xml.MarshalIndent(request,"","  ")
  if err!=nil{
    return nil,err
  }
  request_xml = append([]byte(xml.Header),request_xml...)
  request_message := bytes.NewReader(request_xml)

  res, err := http.Post(camInfo.XAddrs,"application/soap+xml",request_message)
  if err!=nil{
    return nil,err
  }
  defer res.Body.Close()

  data_received ,_ := io.ReadAll(res.Body)
  return data_received,nil
  
}
