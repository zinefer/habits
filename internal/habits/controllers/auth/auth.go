package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"

	"github.com/zinefer/habits/internal/habits/middlewares/session"
	"github.com/zinefer/habits/internal/habits/models/user"
)

// SignIn initiates an oauth handshake
func SignIn() func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		provider := chi.URLParam(req, "provider")

		ctx := context.WithValue(req.Context(), gothic.ProviderParamKey, provider)
		req = req.WithContext(ctx)

		// try to get the user without re-authenticating
		if gothUser, err := gothic.CompleteUserAuth(res, req); err == nil {
			postLogin(ctx, res, req, gothUser)
		} else {
			gothic.BeginAuthHandler(res, req)
		}
	}
}

// Callback completes an oauth handshake
func Callback() func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		provider := chi.URLParam(req, "provider")

		ctx := context.WithValue(req.Context(), gothic.ProviderParamKey, provider)
		req = req.WithContext(ctx)

		user, err := gothic.CompleteUserAuth(res, req)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}

		postLogin(ctx, res, req, user)
	}
}

// SignOut destroys session
func SignOut() func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		sess := session.GetSessionFromContext(req)	
		sess.Options.MaxAge = -1
		sess.Save(req, res)

		c := makeBasicCurrentUserCookie()
		c.MaxAge = -1
		http.SetCookie(res, c)

		gothic.Logout(res, req)
		redirectToIndex(res)
	}
}

func postLogin(ctx context.Context, res http.ResponseWriter, req *http.Request, user goth.User) {
	u := saveUserToDatabase(ctx, user)
	setCurrentUserSession(res, req, u)
	setCurrentUserCookie(res, user.NickName)
	redirectToIndex(res)
}

func saveUserToDatabase(ctx context.Context, gu goth.User) *user.User {
	u := user.New(gu.Name, gu.NickName, gu.Email, gu.Provider)
	_, err := u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return u
}

func makeBasicCurrentUserCookie() *http.Cookie {
	return &http.Cookie{
		Name:   "current_user",
		Path:   "/",
		HttpOnly: false,
	}
}

func setCurrentUserCookie(res http.ResponseWriter, user string) {
	c := makeBasicCurrentUserCookie()
	c.Value = user
	http.SetCookie(res, c)
}

func setCurrentUserSession(res http.ResponseWriter, req *http.Request, user *user.User) {
	sess := session.GetSessionFromContext(req)
	sess.Values["current_user"] = user
	err := sess.Save(req, res)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func redirectToIndex(res http.ResponseWriter) {
	res.Header().Set("Location", "/")
	res.WriteHeader(http.StatusTemporaryRedirect)
}
