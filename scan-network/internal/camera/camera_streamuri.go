package camera 

type StreamUri struct{
	Uri string `xml:"Body>GetStreamUriResponse>MediaUri>Uri"`
}