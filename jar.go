package cookiejar

import(
    "net/http"
    "net/url"
    "sync"
)

type CookieSlice []*http.Cookie

type CookieJar struct {
    data map[string]CookieSlice
    lock sync.Mutex
}

func (jar *CookieJar) SetCookies(u *url.URL, cookies CookieSlice) {
    jar.lock.Lock()
    jar.data[u.Host] = cookies
    jar.lock.Unlock()
}

func (jar *CookieJar) Cookies(u *url.URL) CookieSlice {
    // FIXME: This is a very naive implementation
    return jar.data[u.Host]
}

func NewCookieJar() CookieJar {
    return CookieJar{
        data: make(map[string]CookieSlice),
        lock: sync.Mutex{},
    }
}