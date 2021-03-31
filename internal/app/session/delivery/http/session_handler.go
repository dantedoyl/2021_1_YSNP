package delivery

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2021_1_YSNP/internal/app/errors"
	"github.com/go-park-mail-ru/2021_1_YSNP/internal/app/middleware"
	"github.com/go-park-mail-ru/2021_1_YSNP/internal/app/models"
	"github.com/go-park-mail-ru/2021_1_YSNP/internal/app/session"
	"github.com/go-park-mail-ru/2021_1_YSNP/internal/app/user"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type SessionHandler struct {
	sessUcase session.SessionUsecase
	userUcase user.UserUsecase
}

func NewSessionHandler(sessUcase session.SessionUsecase, userUcase user.UserUsecase) *SessionHandler {
	return &SessionHandler{
		sessUcase: sessUcase,
		userUcase: userUcase,
	}
}

func (sh *SessionHandler) Configure(r *mux.Router, mw *middleware.Middleware) {
	r.HandleFunc("/login", sh.LoginHandler).Methods(http.MethodPost)
	r.HandleFunc("/logout", mw.CheckAuthMiddleware(sh.LogoutHandler)).Methods(http.MethodPost)
}

func (sh *SessionHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	login := &models.LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		logrus.Error(err)
		errE := errors.UnexpectedBadRequest(err)
		w.WriteHeader(errE.HttpError)
		w.Write(errors.JSONError(errE))
		return
	}

	user, errE := sh.userUcase.GetByTelephone(login.Telephone)
	if errE != nil {
		logrus.Error(errE.Message)
		w.WriteHeader(errE.HttpError)
		w.Write(errors.JSONError(errE))
		return
	}

	errE = sh.userUcase.CheckPassword(user, login.Password)
	if errE != nil {
		logrus.Error(errE.Message)
		w.WriteHeader(errE.HttpError)
		w.Write(errors.JSONError(errE))
		return
	}

	session := models.CreateSession(user.ID)
	errE = sh.sessUcase.Create(session)
	if errE != nil {
		logrus.Error(errE.Message)
		w.WriteHeader(errE.HttpError)
		w.Write(errors.JSONError(errE))
		return
	}

	cookie := http.Cookie{
		Name:     "session_id",
		Value:    session.Value,
		Expires:  session.ExpiresAt,
		Secure:   false,
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
	w.Write(errors.JSONSuccess("Successful login."))
}

func (sh *SessionHandler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		errE := errors.Cause(errors.UserUnauthorized)
		logrus.Error(errE.Message)
		w.WriteHeader(errE.HttpError)
		w.Write(errors.JSONError(errE))
		return
	}

	errE := sh.sessUcase.Delete(session.Value)
	if errE != nil {
		logrus.Error(errE.Message)
		w.WriteHeader(errE.HttpError)
		w.Write(errors.JSONError(errE))
		return
	}

	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)

	w.WriteHeader(http.StatusOK)
	w.Write(errors.JSONSuccess("logout success"))
}
