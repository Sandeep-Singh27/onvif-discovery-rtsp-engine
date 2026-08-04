package parsereply
import (
  "fmt"
  "encoding/xml"
  "strings"
)

const (
  cameraReply string = `<?xml version="1.0" encoding="UTF-8"?>
<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://www.w3.org/2003/05/soap-envelope" xmlns:SOAP-ENC="http://www.w3.org/2003/05/soap-encoding" xmlns:wa="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:wd="http://schemas.xmlsoap.org/ws/2005/04/discovery" xmlns:dn="http://www.onvif.org/ver10/network/wsdl" xmlns:tds="http://www.onvif.org/ver10/device/wsdl"><SOAP-ENV:Header><wa:MessageID>uuid:2419d68a-2dd2-21b2-a205-1d61ad35c0de</wa:MessageID>  <wa:RelatesTo>uuid:550e8400-e29b-41d4-a716-446655440000</wa:RelatesTo><wa:To SOAP-ENV:mustUnderstand="true">http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</wa:To><wa:Action SOAP-ENV:mustUnderstand="true">http://schemas.xmlsoap.org/ws/2005/04/discovery/ProbeMatches</wa:Action></SOAP-ENV:Header><SOAP-ENV:Body><wd:ProbeMatches><wd:ProbeMatch> <wa:EndpointReference><wa:Address>urn:uuid:2419d68a-2dd2-21b2-a205-009937b0f3ba</wa:Address></wa:EndpointReference><wd:Types>dn:NetworkVideoTransmitter tds:Device</wd:Types><wd:Scopes>onvif://www.onvif.org/Profile/Streaming onvif://www.onvif.org/Profile/T onvif://www.onvif.org/type/Network_Video_Transmitter  onvif://www.onvif.org/hardware/IPC onvif://www.onvif.org/name/ </wd:Scopes><wd:XAddrs>http://192.168.1.12:80/onvif/device_service</wd:XAddrs><wd:MetadataVersion>1</wd:MetadataVersion></wd:ProbeMatch></wd:ProbeMatches></SOAP-ENV:Body></SOAP-ENV:Envelope>
`
)

type CameraInfo struct{
  Type string `xml:"Body>ProbeMatches>ProbeMatch>Types"`
  Scopes string `xml:"Body>ProbeMatches>ProbeMatch>Scopes"`
  XAddrs string `xml:"Body>ProbeMatches>ProbeMatch>XAddrs"`
  ScopesArr []string
}

//To separate scope strings into an array
func (c *CameraInfo) resolveScopes(){
  c.ScopesArr = strings.Split(c.Scopes," ")
}

//TO CONVERT XML TO PROPER FORMAT
func parseCameraReply(cameraReply string){
  var camInfo CameraInfo
  xml_byte := []byte(cameraReply)

  err := xml.Unmarshal(xml_byte, &camInfo)  

  if err!=nil{
    fmt.Println("[Error] ",err)
  } else {
    //Resolving Scopes to array
    camInfo.resolveScopes()
    fmt.Printf("Type: %s\n",camInfo.Type)
    fmt.Printf("XAddrs: %s\n",camInfo.XAddrs)
    fmt.Printf("ScopesArr: %s\n",camInfo.ScopesArr)
  }
}

