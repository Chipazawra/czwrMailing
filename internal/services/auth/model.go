package auth

import (
	"github.com/Chipazawra/czwrmailing/pkg/jwtmng"
)

type Auth struct {
	TokenManager *jwtmng.Mng
	config       AuthConf
}

type AuthConf struct {
	Users      map[string]string `yaml:"users"`
	JwtTTL     int               `yaml:"jwtttl"`
	RefreshTTL int               `yaml:"refreshttl"`
	Secret     string            `yaml:"secret"`
}

func New(tm *jwtmng.Mng, c *AuthConf) *Auth {
	return &Auth{TokenManager: tm, config: *c}
}
