package camera
import (
  "strings"
)

type CameraInfo struct{
  Type string `xml:"Body>ProbeMatches>ProbeMatch>Types"`
  Scopes string `xml:"Body>ProbeMatches>ProbeMatch>Scopes"`
  XAddrs string `xml:"Body>ProbeMatches>ProbeMatch>XAddrs"`
  ScopesArr []string
}

//To separate scope strings into an array
func (c *CameraInfo) ResolveScopes(){
  c.ScopesArr = strings.Split(c.Scopes," ")
}
