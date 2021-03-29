package usecase

import (
	"time"

	"github.com/go-park-mail-ru/2021_1_YSNP/internal/app/errors"
	"github.com/go-park-mail-ru/2021_1_YSNP/internal/app/models"
	"github.com/go-park-mail-ru/2021_1_YSNP/internal/app/session"
)

type SessionUsecase struct {
	sessRepo session.SessionRepository
}

func NewSessionUsecase(repo session.SessionRepository) session.SessionUsecase {
	return &SessionUsecase{
		sessRepo: repo,
	}
}

func (su *SessionUsecase) Create(sess *models.Session) *errors.Error {
	err := su.sessRepo.Insert(sess)
	if err != nil {
		//TODO: создать ошибку
		return &errors.Error{Message: err.Error()}
	}
	return nil
}

func (su *SessionUsecase) Get(sessValue string) (*models.Session, *errors.Error) {
	sess, err := su.sessRepo.SelectByValue(sessValue)
	if err != nil {
		//TODO: создать ошибку
		return nil, &errors.Error{Message: err.Error()}
	}
	return sess, nil
}

func (su *SessionUsecase) Delete(sessionValue string) *errors.Error {
	if _, err := su.Get(sessionValue); err != nil {
		//TODO: создать ошибку
		return &errors.Error{Message: err.Error()}
	}

	err := su.sessRepo.DeleteByValue(sessionValue)
	if err != nil {
		//TODO: создать ошибку
		return &errors.Error{Message: err.Error()}
	}
	return nil
}

func (su *SessionUsecase) Check(sessValue string) (*models.Session, *errors.Error) {
	sess, err := su.Get(sessValue)
	if err != nil {
		//TODO: создать ошибку
		return nil, &errors.Error{Message: err.Error()}
	}

	if sess.ExpiresAt.Before(time.Now()) {
		err := su.Delete(sessValue)
		if err != nil {
			//TODO: создать ошибку
			return nil, &errors.Error{Message: err.Error()}
		}
		//TODO: создать ошибку
		return nil, &errors.Error{Message: "fff"} //созданная ошибка
	}

	return sess, nil
}
