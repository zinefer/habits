package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"

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
			setCurrentUserCookie(res, gothUser.NickName)
			redirectToIndex(res)
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

		saveUserToDatabase(ctx, user)
		setCurrentUserCookie(res, user.NickName)
		redirectToIndex(res)
	}
}

// SignOut destroys session
func SignOut() func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		gothic.Logout(res, req)
		redirectToIndex(res)
	}
}

func saveUserToDatabase(ctx context.Context, gu goth.User) {
	u := user.New(gu.Name, gu.NickName, gu.Email, gu.Provider)
	_, err := u.Save(ctx)
	if err != nil {
		panic(err)
	}
}

func setCurrentUserCookie(res http.ResponseWriter, user string) {
	cookie := &http.Cookie{Name: "current_user", Value: user, HttpOnly: false, Path: "/"}
	http.SetCookie(res, cookie)
}

func redirectToIndex(res http.ResponseWriter) {
	res.Header().Set("Location", "/")
	res.WriteHeader(http.StatusTemporaryRedirect)
}
