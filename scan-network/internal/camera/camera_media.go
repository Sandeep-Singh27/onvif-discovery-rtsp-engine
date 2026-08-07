package camera
import (
)

type CameraMedia struct{
	Media []media `xml:"Body>GetCapabilitiesResponse>Capabilities>Media"`
}

type media struct{
	XAddrs string `xml:"XAddr"`
}