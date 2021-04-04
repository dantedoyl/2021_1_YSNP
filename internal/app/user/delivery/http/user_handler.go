package delivery

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2021_1_YSNP/internal/app/middleware"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-park-mail-ru/2021_1_YSNP/internal/app/errors"
	"github.com/go-park-mail-ru/2021_1_YSNP/internal/app/models"
	"github.com/go-park-mail-ru/2021_1_YSNP/internal/app/session"
	"github.com/go-park-mail-ru/2021_1_YSNP/internal/app/user"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	userUcase user.UserUsecase
	sessUcase session.SessionUsecase
}

func NewUserHandler(userUcase user.UserUsecase, sessUcase session.SessionUsecase) *UserHandler {
	return &UserHandler{
		userUcase: userUcase,
		sessUcase: sessUcase,
	}
}

func (uh *UserHandler) Configure(r *mux.Router, mw *middleware.Middleware) {
	r.HandleFunc("/signup", uh.SignUpHandler).Methods(http.MethodPost)
	r.HandleFunc("/upload", mw.CheckAuthMiddleware(uh.UploadAvatarHandler)).Methods(http.MethodPost)
	r.HandleFunc("/me", mw.CheckAuthMiddleware(uh.GetProfileHandler)).Methods(http.MethodGet)
	r.HandleFunc("/settings", mw.CheckAuthMiddleware(uh.ChangeProfileHandler)).Methods(http.MethodPost)
	r.HandleFunc("/settings/password", mw.CheckAuthMiddleware(uh.ChangeProfilePasswordHandler)).Methods(http.MethodPost)
}

func (uh *UserHandler) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	signUp := models.SignUpRequest{}
	err := json.NewDecoder(r.Body).Decode(&signUp)
	if err != nil {
		logrus.Error(err)
		errE := errors.UnexpectedBadRequest(err)
		w.WriteHeader(errE.HttpError)
		w.Write(errors.JSONError(errE))
		return
	}

	user := &models.UserData{
		Name:       signUp.Name,
		Surname:    signUp.Surname,
		Sex:        signUp.Sex,
		Email:      signUp.Email,
		Telephone:  signUp.Telephone,
		Password:   signUp.Password1,
		DateBirth:  signUp.DateBirth,
		LinkImages: signUp.LinkImages,
	}
	errE := uh.userUcase.Create(user)
	if errE != nil {
		logrus.Error(errE.Message)
		w.WriteHeader(errE.HttpError)
		w.Write(errors.JSONError(errE))
		return
	}

	session := models.CreateSession(user.ID)
	errE = uh.sessUcase.Create(session)
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
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
	w.Write(errors.JSONSuccess("Successful login."))
}

func (uh *UserHandler) UploadAvatarHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.ContextUserID).(uint64)
	if !ok {
		errE := errors.Cause(errors.UserUnauthorized)
		logrus.Error(errE.Message)
		w.WriteHeader(errE.HttpError)
		w.Write(errors.JSONError(errE))
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 3*1024*1024)
	err := r.ParseMultipartForm(3 * 1024 * 1024)
	if err != nil {
		logrus.Error(err)
		errE := errors.UnexpectedBadRequest(err)
		w.WriteHeader(errE.HttpError)
		w.Write(errors.JSONError(errE))
		return
	}

	file, handler, err := r.FormFile("file-upload")
	if err != nil {
		logrus.Error(err)
		errE := errors.UnexpectedBadRequest(err)
		w.WriteHeader(errE.HttpError)
		w.Write(errors.JSONError(errE))
		return
	}
	defer file.Close()
	extension := filepath.Ext(handler.Filename)

	r.FormValue("file-upload")

	str, err := os.Getwd()
	if err != nil {
		logrus.Error(err)
		errE := errors.UnexpectedInternal(err)
		w.WriteHeader(errE.HttpError)
		w.Write(errors.JSONError(errE))
		return
	}

	photoPath := "static/avatar/"
	os.Chdir(photoPath)

	photoID, err := uuid.NewRandom()
	if err != nil {
		logrus.Error(err)
		errE := errors.UnexpectedInternal(err)
		w.WriteHeader(errE.HttpError)
		w.Write(errors.JSONError(errE))
		return
	}

	f, err := os.OpenFile(photoID.String()+extension, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		logrus.Error(err)
		errE := errors.UnexpectedInternal(err)
		w.WriteHeader(errE.HttpError)
		w.Write(errors.JSONError(errE))
		return
	}
	defer f.Close()

	os.Chdir(str)

	_, err = io.Copy(f, file)
	if err != nil {
		_ = os.Remove(photoID.String() + extension)
		logrus.Error(err)
		errE := errors.UnexpectedInternal(err)
		w.WriteHeader(errE.HttpError)
		w.Write(errors.JSONError(errE))
		return
	}

	avatar := "/static/avatar/" + photoID.String() + extension

	_, errE := uh.userUcase.UpdateAvatar(userID, avatar)
	if errE != nil {
		logrus.Error(errE.Message)
		w.WriteHeader(errE.HttpError)
		w.Write(errors.JSONError(errE))
		return
	}

	body, err := json.Marshal(avatar)
	if err != nil {
		logrus.Error(err)
		errE := errors.UnexpectedInternal(err)
		w.WriteHeader(errE.HttpError)
		w.Write(errors.JSONError(errE))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (uh *UserHandler) GetProfileHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.ContextUserID).(uint64)
	if !ok {
		errE := errors.Cause(errors.UserUnauthorized)
		logrus.Error(errE.Message)
		w.WriteHeader(errE.HttpError)
		w.Write(errors.JSONError(errE))
		return
	}

	user, errE := uh.userUcase.GetByID(userID)
	if errE != nil {
		logrus.Error(errE.Message)
		w.WriteHeader(errE.HttpError)
		w.Write(errors.JSONError(errE))
		return
	}

	body, err := json.Marshal(user)
	if err != nil {
		logrus.Error(err)
		errE := errors.UnexpectedInternal(err)
		w.WriteHeader(errE.HttpError)
		w.Write(errors.JSONError(errE))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (uh *UserHandler) ChangeProfileHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.ContextUserID).(uint64)
	if !ok {
		errE := errors.Cause(errors.UserUnauthorized)
		logrus.Error(errE.Message)
		w.WriteHeader(errE.HttpError)
		w.Write(errors.JSONError(errE))
		return
	}

	changeData := models.SignUpRequest{}
	err := json.NewDecoder(r.Body).Decode(&changeData)
	if err != nil {
		logrus.Error(err)
		errE := errors.UnexpectedBadRequest(err)
		w.WriteHeader(errE.HttpError)
		w.Write(errors.JSONError(errE))
		return
	}

	user := &models.UserData{
		Name:      changeData.Name,
		Surname:   changeData.Surname,
		Sex:       changeData.Sex,
		Email:     changeData.Email,
		Telephone: changeData.Telephone,
		DateBirth: changeData.DateBirth,
	}

	_, errE := uh.userUcase.UpdateProfile(userID, user)
	if errE != nil {
		logrus.Error(errE.Message)
		w.WriteHeader(errE.HttpError)
		w.Write(errors.JSONError(errE))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(errors.JSONSuccess("Successful change."))
}

func (uh *UserHandler) ChangeProfilePasswordHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.ContextUserID).(uint64)
	if !ok {
		errE := errors.Cause(errors.UserUnauthorized)
		logrus.Error(errE.Message)
		w.WriteHeader(errE.HttpError)
		w.Write(errors.JSONError(errE))
		return
	}

	passwordData := models.PasswordChangeRequest{}
	err := json.NewDecoder(r.Body).Decode(&passwordData)
	if err != nil {
		logrus.Error(err)
		errE := errors.UnexpectedBadRequest(err)
		w.WriteHeader(errE.HttpError)
		w.Write(errors.JSONError(errE))
		return
	}

	_, errE := uh.userUcase.UpdatePassword(userID, passwordData.NewPassword)
	if errE != nil {
		logrus.Error(errE.Message)
		w.WriteHeader(errE.HttpError)
		w.Write(errors.JSONError(errE))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(errors.JSONSuccess("Successful change."))
}
