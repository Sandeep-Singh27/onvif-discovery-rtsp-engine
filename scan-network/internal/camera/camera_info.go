package camera
import (
  "fmt"
  "strings"
)

type CameraInfo struct{
  IPAddr string
  Type string `xml:"Body>ProbeMatches>ProbeMatch>Types"`
  Scopes string `xml:"Body>ProbeMatches>ProbeMatch>Scopes"`
  XAddrs string `xml:"Body>ProbeMatches>ProbeMatch>XAddrs"`
  ScopesArr []string
  RTSPUrls map[string]string
}

//To separate scope strings into an array
func (c *CameraInfo) ResolveScopes(){
  arr := strings.Split(c.Scopes," ")
  arr2 := make([]string,0)
  for _,v := range arr{
    if v!=""{
      arr2 = append(arr2,v)
    }
  }
  c.ScopesArr = arr2
}

func (c CameraInfo) String() string{
   output := fmt.Sprintf("IP Address: %s\nType: %s\nXAddrs: %s\nScopesArr:\n",c.IPAddr,c.Type,c.XAddrs)
   for _,v := range c.ScopesArr{
    output += fmt.Sprintf("\t%s\n",v)
   }
   output += "RTSPUrls: "
   if c.RTSPUrls != nil{
    output += "\n"
    for k,v := range c.RTSPUrls{
      output += fmt.Sprintf("\t%s : %s\n",k,v)
    }
   }
   return output
}