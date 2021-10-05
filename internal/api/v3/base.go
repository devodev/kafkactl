package v3

type V3Base struct {
	Kind     string `json:"kind"`
	Metadata struct {
		Self string `json:"self"`
		Next string `json:"next"`
	} `json:"metadata"`
}

type V3BaseData struct {
	Kind     string `json:"kind"`
	Metadata struct {
		Self         string `json:"self"`
		ResourceName string `json:"resource_name"`
	} `json:"metadata"`
}

type V3BaseDataRelated struct {
	Related string `json:"related"`
}
