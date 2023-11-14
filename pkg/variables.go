package kubazulo

var (
	Cfg_client_id   string
	Cfg_tenant_id   string
	Cfg_force_login string
)

var (
	AuthorizationURL string
	TokenURL         string
)

type Session struct {
	TokenStartTimestamp int64  `json:"tokenstartTimestamp"`
	ExpirationTimestamp int64  `json:"expirationTimestamp"`
	AccessToken         string `json:"access_token"`
	RefreshToken        string `json:"refresh_token"`
}

func FillVariables() {
	// AuthorizationURL is the endpoint used for initial login/auth
	AuthorizationURL = "https://login.microsoftonline.com/" + Cfg_tenant_id + "/oauth2/v2.0/authorize"
	// TokenURL is the endpoint for getting access/refresh tokens
	TokenURL = "https://login.microsoftonline.com/" + Cfg_tenant_id + "/oauth2/v2.0/token"
}