package profile

import (
	"github.com/Chipazawra/czwrmailing/pkg/jwtmng"
)

type Profile struct {
	TokenManager *jwtmng.Mng
}

func New(tm *jwtmng.Mng) *Profile {
	return &Profile{TokenManager: tm}
}
