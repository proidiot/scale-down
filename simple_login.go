package scaledown

import (
	"github.com/fitstar/falcore"
	"net/http"
)

const (
	xsrf_name = "xsrf_token"
	user_name = "username"
	pass_name = "password"
)

func NewSimpleLoginFilter(ckm *scaledown.CookieManager) falcore.RequestFilter {
	return falcore.NewRequestFilter(
		func(req *falcore.Request) *http.Response {
			bad_creds := ""
			cke, err := req.Cookie(scaledown.CookieName)
			if cke != nil {
				if ckm.CookieCheck(cke) {
					return nil
				} else if req.Method == "POST" && ckm.ValidateXSRF(req.FormValue(xsrf_name)) && authenticator.ValidCreds(req.FormValue(user_name), req.FormValue(pass_name)) {
					ckm.CookieAuthenticated(cke)
					return nil
				} else {
					bad_creds = "<h2>Invalid login credentials. Please try again.</h2>"
				}
			} else {
				cke = ckm.GenCookie()
				req.AddCookie(cke)
			}
			return falcore.StringResponse(
				req.HttpRequest,
				401,
				nil,
				strings.Join(append([]string, "<html><head><title>Pullcord Login</title></head><body><h1>Pullcord Login</h1><form method=\"POST\" action=\"", thing, "\"><fieldset><legend>Pullcord Login</legend>", bad_creds, "<table><tr><td><label>Username</label></td><td><input type=\"text\" name=\"", user_name, "\" /></td></tr><tr><td><label>Password</label></td><td><input type=\"password\" name=\"", pass_name, "\" /></td></tr></table><input type=\"hidden\" name=\"", xsrf_name, "\" value=\"", xsrf_token, "\" /><input type=\"submit\" /></fieldset></form></body></html>"), ""))
		})
}

