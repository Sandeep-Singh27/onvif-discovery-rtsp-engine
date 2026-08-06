package getrtspurl
import (
  "fmt"
  "net/http"
  "strings"
  "io"
  "errors"
  "camera-scanner/scan-network/internal/camera"
)

const (
  get_capability_message string = `<?xml version="1.0" encoding="UTF-8"?>

<soap:Envelope
    xmlns:soap="http://www.w3.org/2003/05/soap-envelope"
    xmlns:tds="http://www.onvif.org/ver10/device/wsdl">

    <soap:Body>
        <tds:GetCapabilities>
            <tds:Category>All</tds:Category>
        </tds:GetCapabilities>
    </soap:Body>

</soap:Envelope>`
)


func GetCapabilities(camInfo camera.CameraInfo)([]byte,error){
  if camInfo.XAddrs!=""{
    request_xml := strings.NewReader(get_capability_message)
    res, err := http.Post(camInfo.XAddrs,"application/soap+xml",request_xml)
    if err==nil{
      defer res.Body.Close()
      data_received,_ := io.ReadAll(res.Body)
      fmt.Printf("[%s] %s\n",camInfo.XAddrs,data_received)
      return data_received,nil
    }else{
      return nil,err
    }
  }else{
    return nil,errors.New("XAddrs is empty")
  }  
}

