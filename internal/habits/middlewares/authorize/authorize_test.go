package authorize_test

import (
	"encoding/gob"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/go-chi/chi"
	"github.com/zinefer/habits/internal/habits/helpers/test"
	"github.com/zinefer/habits/internal/habits/middlewares/authorize"
	"github.com/zinefer/habits/internal/habits/middlewares/session"
	"github.com/zinefer/habits/internal/habits/models/user"
)

type TestSuite struct {
	suite.Suite
	ts *httptest.Server
	r  *chi.Mux
	s  *sessions.CookieStore
}

func (suite *TestSuite) setupRouter(authed bool) {
	suite.r.Use(session.SessionContextMiddleware(suite.s))

	if authed {
		suite.r.Use(fakeUserSessionSetterMiddleware())
	}

	suite.r.Use(authorize.AuthorizeMiddleware())

	suite.r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root"))
	})

	suite.ts = httptest.NewServer(suite.r)
}

func (suite *TestSuite) SetupTest() {
	gob.Register(&user.User{})
	suite.r = chi.NewRouter()
	suite.s = sessions.NewCookieStore([]byte("authorizetest"))
}

func (suite *TestSuite) TestAuthorizeMiddleware401() {
	suite.setupRouter(false)

	resp, _ := test.Request(suite.T(), suite.ts, "GET", "/", nil)
	assert.Equal(suite.T(), http.StatusUnauthorized, resp.StatusCode)
}

func (suite *TestSuite) TestAuthorizeMiddleware200() {
	suite.setupRouter(true)

	resp, body := test.Request(suite.T(), suite.ts, "GET", "/", nil)
	assert.Equal(suite.T(), 200, resp.StatusCode)
	assert.Equal(suite.T(), "root", body)
}

func (suite *TestSuite) TearDownTest() {
	suite.ts.Close()
}

func fakeUserSessionSetterMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sess := session.GetSessionFromContext(r)
			sess.Values["current_user"] = &user.User{}
			err := sess.Save(r, w)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			next.ServeHTTP(w, r)
		})
	}
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
