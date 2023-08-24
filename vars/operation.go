package vars

type Operation struct {
	OP    string `json:"op"`
	Path  string `json:"path"`
	Value string `json:"value"`
}
