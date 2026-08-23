package main//getrtspurl
import (
  "fmt"
  "net/http"
  "strings"
  "io"
)

const (
  get_profiles_message string = `<?xml version="1.0" encoding="UTF-8"?>
    <s:Envelope>
      <s:Body>
        <trt:GetProfiles/>
      </s:Body>
    </s:Envelope>
  `
)

func GetProfiles(mediaXAddr string)([]byte,error){
  request_xml := strings.NewReader(get_profiles_message)
  res, err := http.Post(mediaXAddr,"application/soap+xml",request_xml)
  if err!=nil{
    return nil,err
  }

  defer res.Body.Close()
  data_received,_ := io.ReadAll(res.Body)
  return data_received,nil
}

func main(){
  mediaXAddr := "http://192.168.1.12:80/onvif/Media"
  reply,_ := GetProfiles(mediaXAddr)
  reply_string := string(reply)
  fmt.Println(reply_string)
}
