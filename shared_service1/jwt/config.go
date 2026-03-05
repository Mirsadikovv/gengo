package jwt

type JwtConfig struct {
	Secret         string `json:"secret" env:"SECRET"`
	Expired        int64  `json:"expired" env:"EXPIRED"`
	RefreshExpired int64  `json:"refreshExpired" env:"REFRESH_EXPIRED"`
}

func (j *JwtConfig) SecretData() string {
	return j.Secret
}

func (j *JwtConfig) Exp() int64 {
	return j.Expired
}

func (j *JwtConfig) RefreshExp() int64 {
	return j.RefreshExpired
}
