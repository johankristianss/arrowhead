package arrowhead

type Service interface {
	HandleRequest(params *Params) ([]byte, error)
}

type Params struct {
	QueryParams map[string]string `json:"params"`
	Payload     []byte            `json:"payload"`
}

func EmptyParams() *Params {
	return &Params{
		QueryParams: make(map[string]string),
		Payload:     nil,
	}
}
