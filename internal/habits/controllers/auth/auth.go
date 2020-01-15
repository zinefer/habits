package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"

	petname "github.com/dustinkirkland/golang-petname"

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
			u, _ := user.FindByProviderID(req.Context(), gothUser.Provider, gothUser.UserID)
			if u.ID == 0 {
				u = addUserToDatabase(ctx, gothUser)
			}

			postLogin(ctx, res, req, u)
		} else {
			gothic.BeginAuthHandler(res, req)
		}
	}
}

// FakeAuth is a helper for development to relieve the need to setup oAuth
// DO NOT RUN THIS IN PRODUCTION
func FakeAuth() func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		userName := chi.URLParam(req, "user")

		u, _ := user.FindByName(req.Context(), userName)
		if u.ID == 0 {
			http.Error(res, http.StatusText(404), 404)
			return
		}

		postLogin(req.Context(), res, req, u)
	}
}

// Callback completes an oauth handshake
func Callback() func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		provider := chi.URLParam(req, "provider")

		ctx := context.WithValue(req.Context(), gothic.ProviderParamKey, provider)
		req = req.WithContext(ctx)

		gothUser, err := gothic.CompleteUserAuth(res, req)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}

		u, _ := user.FindByProviderID(req.Context(), gothUser.Provider, gothUser.UserID)
		if u.ID == 0 {
			u = addUserToDatabase(ctx, gothUser)
		}

		postLogin(ctx, res, req, u)
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

func postLogin(ctx context.Context, res http.ResponseWriter, req *http.Request, u *user.User) {
	setCurrentUserSession(res, req, u)
	setCurrentUserCookie(res, u.Name)
	redirectToIndex(res)
}

func addUserToDatabase(ctx context.Context, gu goth.User) *user.User {
	u := user.New(gu.UserID, gu.Provider, gu.NickName, gu.Name, gu.Email)
	err := preventNameCollisions(ctx, u)
	if err != nil {
		panic(err)
	}

	err = u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return u
}

func makeBasicCurrentUserCookie() *http.Cookie {
	return &http.Cookie{
		Name:     "current_user",
		Path:     "/",
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

func preventNameCollisions(ctx context.Context, user *user.User) error {
	return preventNameCollisionsRecursive(ctx, user, 0)
}

func preventNameCollisionsRecursive(ctx context.Context, u *user.User, tries int) error {
	unique, err := user.IsNameAvailable(ctx, u.Name)
	if !unique || u.Name == "" {
		u.Name = petname.Generate(min(tries/5, 2), "-")
		return preventNameCollisionsRecursive(ctx, u, tries+1)
	}
	return err
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
