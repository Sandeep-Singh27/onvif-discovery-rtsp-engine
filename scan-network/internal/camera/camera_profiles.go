package camera

type CameraProfiles struct{
  Profile []profile `xml:"Body>GetProfilesResponse>Profiles"`
}

type profile struct{
  Token string `xml:"token,attr"`
}
