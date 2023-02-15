package cookie

import (
	"net/http"
	"time"
)

type CookieService interface {
	SetCookie(string, string, string, time.Time) *http.Cookie // Set Cookie
	DelCookie(string, string) *http.Cookie                    // Delete Cookie
}
