package cookie

import (
	"net/http"
	"time"
)

type Cookie struct {
	*http.Cookie
}

func NewCookieService() CookieService {
	return &Cookie{new(http.Cookie)}
}

func (c *Cookie) SetCookie(name, value, path string, expires time.Time) *http.Cookie {
	c.Name = name
	c.Value = value
	c.Path = path
	c.Expires = expires.Add(60000000000 * 60 * 6) // 이후 6시간 동안 유효
	return c.Cookie
}

func (c *Cookie) DelCookie(name, path string) *http.Cookie {
	c.Name = name
	c.Value = ""
	c.Path = path
	c.Expires = time.Unix(0, 0)
	return c.Cookie
}
