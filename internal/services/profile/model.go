package profile

import (
	"github.com/Chipazawra/czwrmailing/pkg/jwtmng"
)

type Profile struct {
	TokenManager *jwtmng.Mng
}

type Template struct {
	uuid    string
	rawdata string
	params  []string
}

func New(tm *jwtmng.Mng) *Profile {
	return &Profile{TokenManager: tm}
}
