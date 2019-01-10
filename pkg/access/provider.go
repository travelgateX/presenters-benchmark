package access

type OrgProvider struct {
	Active int `json:"act"`
}

type Provider struct {
	OrgProvider
	Dll     string `json:"dll"`
	Context string `json:"context"`
}
