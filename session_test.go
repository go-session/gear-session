package gearsession

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/teambition/gear"
	"gopkg.in/session.v2"
)

func TestSession(t *testing.T) {
	cookieName := "test_gear_session"
	app := gear.New()
	app.Use(New(
		session.SetCookieName(cookieName),
		session.SetSign([]byte("sign")),
		session.SetCookieLifeTime(60),
		session.SetExpired(10),
	))

	app.Use(func(ctx *gear.Context) error {
		store := FromContext(ctx)

		if ctx.Query("login") == "1" {
			foo, ok := store.Get("foo")
			if !ok || foo != "bar" {
				t.Error("Not expected value:", foo)
				return nil
			}
			fmt.Fprint(ctx.Res, "ok")
			return nil
		}

		store.Set("foo", "bar")
		err := store.Save()
		if err != nil {
			return err
		}
		fmt.Fprint(ctx.Res, "ok")
		return nil
	})
	srv := app.Start()
	defer srv.Close()

	url := "http://" + srv.Addr().String()
	res, err := http.Get(url)
	if err != nil {
		t.Error(err)
		return
	}

	cookie := res.Cookies()[0]
	if cookie.Name != cookieName {
		t.Error("Not expected value:", cookie.Name)
		return
	}

	buf, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if string(buf) != "ok" {
		t.Error("Not expected value:", string(buf))
		return
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s?login=1", url), nil)
	if err != nil {
		t.Error(err)
		return
	}
	req.AddCookie(cookie)

	res, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
		return
	}

	buf, _ = ioutil.ReadAll(res.Body)
	res.Body.Close()
	if string(buf) != "ok" {
		t.Error("Not expected value:", string(buf))
		return
	}
}
