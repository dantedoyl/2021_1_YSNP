package http

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/go-park-mail-ru/2021_1_YSNP/internal/app/middleware"
	"github.com/microcosm-cc/bluemonday"
	"github.com/sirupsen/logrus"

	"github.com/go-park-mail-ru/2021_1_YSNP/internal/app/errors"
	"github.com/go-park-mail-ru/2021_1_YSNP/internal/app/models"
	"github.com/go-park-mail-ru/2021_1_YSNP/internal/app/product"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type ProductHandler struct {
	productUcase product.ProductUsecase
}

func NewProductHandler(productUcase product.ProductUsecase) *ProductHandler {
	return &ProductHandler{
		productUcase: productUcase,
	}
}

func (ph *ProductHandler) Configure(r *mux.Router, mw *middleware.Middleware) {
	r.HandleFunc("/product/list", ph.MainPageHandler).Methods(http.MethodPost)
	r.HandleFunc("/product/{id:[0-9]+}", ph.ProductIDHandler).Methods(http.MethodGet)
	r.HandleFunc("/product/create", mw.CheckAuthMiddleware(ph.ProductCreateHandler)).Methods(http.MethodPost)
	r.HandleFunc("/product/upload/{pid:[0-9]+}", mw.CheckAuthMiddleware(ph.UploadPhotoHandler)).Methods(http.MethodPost)
	r.HandleFunc("/categories", ph.CategoriesHandler).Methods(http.MethodGet)
	r.HandleFunc("/product/promote", ph.PromoteProductHandler).Methods(http.MethodPost)
}

func (ph *ProductHandler) CategoriesHandler(w http.ResponseWriter, r *http.Request) {
	
	var categories []*models.Category
	categories = append(categories, &models.Category {Title: "Транспорт"}, &models.Category {Title: "Недвижмость"}, &models.Category {Title: "Хобби и отдых"}, &models.Category {Title: "Работа"}, &models.Category {Title: "Для дома и дачи"}, &models.Category {Title: "Бытовая электрика"}, &models.Category {Title: "Личные вещи"}, &models.Category {Title: "Животные"})
	body, err := json.Marshal(categories)
	if err != nil {
		logrus.Error(err)
		errE := errors.UnexpectedInternal(err)
		w.WriteHeader(errE.HttpError)
		w.Write(errors.JSONError(errE))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (ph *ProductHandler) MainPageHandler(w http.ResponseWriter, r *http.Request) {
	logger := r.Context().Value(middleware.ContextLogger).(*logrus.Entry)

	page := models.Page{}
	err := json.NewDecoder(r.Body).Decode(&page)
	if err != nil {
		logger.Error(err)
		errE := errors.UnexpectedBadRequest(err)
		w.WriteHeader(errE.HttpError)
		w.Write(errors.JSONError(errE))
		return
	}
	logger.Info("page ", page)

	_, err = govalidator.ValidateStruct(page)
	if err != nil {
		if allErrs, ok := err.(govalidator.Errors); ok {
			logger.Error(allErrs.Errors())
			errE := errors.UnexpectedBadRequest(allErrs)
			w.WriteHeader(errE.HttpError)
			w.Write(errors.JSONError(errE))
			return
		}
	}

	products, errE := ph.productUcase.ListLatest(&page.Content)
	if errE != nil {
		logger.Error(errE.Message)
		w.WriteHeader(errE.HttpError)
		w.Write(errors.JSONError(errE))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(products)
	if err != nil {
		logger.Error(err)
		errE := errors.UnexpectedInternal(err)
		w.WriteHeader(errE.HttpError)
		w.Write(errors.JSONError(errE))
		return
	}
}

func (ph *ProductHandler) ProductIDHandler(w http.ResponseWriter, r *http.Request) {
	logger := r.Context().Value(middleware.ContextLogger).(*logrus.Entry)

	vars := mux.Vars(r)
	productID, _ := strconv.ParseUint(vars["id"], 10, 64)
	logger.Info("product id ", productID)

	product, errE := ph.productUcase.GetByID(productID)
	if errE != nil {
		logger.Error(errE.Message)
		w.WriteHeader(errE.HttpError)
		w.Write(errors.JSONError(errE))
		return
	}
	logger.Debug("product by id ", product)

	body, err := json.Marshal(product)
	if err != nil {
		logger.Error(err)
		errE := errors.UnexpectedInternal(err)
		w.WriteHeader(errE.HttpError)
		w.Write(errors.JSONError(errE))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (ph *ProductHandler) ProductCreateHandler(w http.ResponseWriter, r *http.Request) {
	logger := r.Context().Value(middleware.ContextLogger).(*logrus.Entry)
	defer r.Body.Close()

	userID, ok := r.Context().Value(middleware.ContextUserID).(uint64)
	if !ok {
		errE := errors.Cause(errors.UserUnauthorized)
		logger.Error(errE.Message)
		w.WriteHeader(errE.HttpError)
		w.Write(errors.JSONError(errE))
		return
	}
	logger.Info("user id ", userID)

	productData := &models.ProductData{}
	err := json.NewDecoder(r.Body).Decode(&productData)
	if err != nil {
		logger.Error(err)
		errE := errors.UnexpectedBadRequest(err)
		w.WriteHeader(errE.HttpError)
		w.Write(errors.JSONError(errE))
		return
	}
	logger.Info("product data ", productData)
	productData.OwnerID = userID

	sanitizer := bluemonday.UGCPolicy()
	productData.Name = sanitizer.Sanitize(productData.Name)
	productData.Description = sanitizer.Sanitize(productData.Description)
	productData.Category = sanitizer.Sanitize(productData.Category)
	logger.Debug("sanitize user data ", productData)

	_, err = govalidator.ValidateStruct(productData)
	if err != nil {
		if allErrs, ok := err.(govalidator.Errors); ok {
			logger.Error(allErrs.Errors())
			errE := errors.UnexpectedBadRequest(allErrs)
			w.WriteHeader(errE.HttpError)
			w.Write(errors.JSONError(errE))
			return
		}
	}

	errE := ph.productUcase.Create(productData)
	if errE != nil {
		logger.Error(errE.Message)
		w.WriteHeader(errE.HttpError)
		w.Write(errors.JSONError(errE))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(errors.JSONSuccess("Successful creation.", productData.ID))
}

func (ph *ProductHandler) UploadPhotoHandler(w http.ResponseWriter, r *http.Request) {
	logger := r.Context().Value(middleware.ContextLogger).(*logrus.Entry)

	vars := mux.Vars(r)
	productID, _ := strconv.ParseUint(vars["pid"], 10, 64)
	logger.Info("product id ", productID)

	userId, ok := r.Context().Value(middleware.ContextUserID).(uint64)
	if !ok {
		errE := errors.Cause(errors.UserUnauthorized)
		logger.Error(errE.Message)
		w.WriteHeader(errE.HttpError)
		w.Write(errors.JSONError(errE))
		return
	}
	logger.Info("user id ", userId)

	r.Body = http.MaxBytesReader(w, r.Body, 10*1024*1024)
	err := r.ParseMultipartForm(10 * 1024 * 1024)
	if err != nil {
		logger.Error(err)
		errE := errors.UnexpectedBadRequest(err)
		w.WriteHeader(errE.HttpError)
		w.Write(errors.JSONError(errE))
		return
	}

	files := r.MultipartForm.File["photos"]
	imgs := make(map[string][]string)
	for i := range files {
		file, err := files[i].Open()
		if err != nil {
			logger.Error(err)
			errE := errors.UnexpectedBadRequest(err)
			w.WriteHeader(errE.HttpError)
			w.Write(errors.JSONError(errE))
			return
		}
		logger.Debug("photo ", files[i].Header)
		defer file.Close()
		extension := filepath.Ext(files[i].Filename)

		str, err := os.Getwd()
		if err != nil {
			logger.Error(err)
			errE := errors.UnexpectedInternal(err)
			w.WriteHeader(errE.HttpError)
			w.Write(errors.JSONError(errE))
			return
		}

		photoPath := "static/product/"
		os.Chdir(photoPath)

		photoID, err := uuid.NewRandom()
		if err != nil {
			logger.Error(err)
			errE := errors.UnexpectedInternal(err)
			w.WriteHeader(errE.HttpError)
			w.Write(errors.JSONError(errE))
			return
		}
		logger.Debug("new photo name ", photoID)

		f, err := os.OpenFile(photoID.String()+extension, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			logger.Error(err)
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
			logger.Error(err)
			errE := errors.UnexpectedInternal(err)
			w.WriteHeader(errE.HttpError)
			w.Write(errors.JSONError(errE))
			return
		}

		imgs["linkImages"] = append(imgs["linkImages"], "/static/product/"+photoID.String()+extension)
	}

	//if len(imgs) == 0 {
	//	logger.Error(err)
	//	w.WriteHeader(http.StatusBadRequest)
	//	w.Write(errors.JSONError(errors.Error{Message: "http: no such file"}.Error()))
	//	return
	//}

	_, errE := ph.productUcase.UpdatePhoto(productID, imgs["linkImages"])
	if errE != nil {
		logger.Error(errE.Message)
		w.WriteHeader(errE.HttpError)
		w.Write(errors.JSONError(errE))
		return
	}

	body, err := json.Marshal(imgs)
	if err != nil {
		logger.Error(err)
		errE := errors.UnexpectedInternal(err)
		w.WriteHeader(errE.HttpError)
		w.Write(errors.JSONError(errE))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (ph *ProductHandler) PromoteProductHandler(w http.ResponseWriter, r *http.Request) {
	logger := r.Context().Value(middleware.ContextLogger).(*logrus.Entry)

	err := r.ParseForm()
	if err != nil {
		logger.Error(err)
		errE := errors.UnexpectedBadRequest(err)
		w.WriteHeader(errE.HttpError)
		w.Write(errors.JSONError(errE))
		return
	}

	label := r.PostFormValue("label")
	if label == "" {
		errE := errors.Cause(errors.PromoteEmptyLabel)
		logger.Error(errE)
		w.WriteHeader(errE.HttpError)
		w.Write(errors.JSONError(errE))
		return
	}
	logger.Debug(label)

	data := strings.Split(label, ",")
	productID, err := strconv.ParseUint(data[0], 10, 64)
	if err != nil {
		logger.Error(err)
		errE := errors.UnexpectedBadRequest(err)
		w.WriteHeader(errE.HttpError)
		w.Write(errors.JSONError(errE))
		return
	}
	logger.Info(productID)

	tariff, err := strconv.Atoi(data[1])
	if err != nil {
		logger.Error(err)
		errE := errors.UnexpectedBadRequest(err)
		w.WriteHeader(errE.HttpError)
		w.Write(errors.JSONError(errE))
		return
	}
	logger.Info(tariff)

	errE := ph.productUcase.SetTariff(productID, tariff)
	if errE != nil {
		logger.Error(errE.Message)
		w.WriteHeader(errE.HttpError)
		w.Write(errors.JSONError(errE))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(errors.JSONSuccess("Successful promotion."))
}
