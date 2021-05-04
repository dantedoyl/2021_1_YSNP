package usecase

import (
	"database/sql"
	"mime/multipart"
	"time"

	errors "github.com/go-park-mail-ru/2021_1_YSNP/internal/app/tools/errors"
	"github.com/go-park-mail-ru/2021_1_YSNP/internal/app/trends"

	"github.com/jackc/pgx"

	"github.com/go-park-mail-ru/2021_1_YSNP/internal/app/models"
	"github.com/go-park-mail-ru/2021_1_YSNP/internal/app/product"
	"github.com/go-park-mail-ru/2021_1_YSNP/internal/app/tools/errors"
	"github.com/go-park-mail-ru/2021_1_YSNP/internal/app/upload"
)

type ProductUsecase struct {
	productRepo product.ProductRepository
	uploadRepo  upload.UploadRepository
	trendsRepo  trends.TrendsRepository
	
}

func NewProductUsecase(repo product.ProductRepository, uploadRepo upload.UploadRepository, trendsRepo trends.TrendsRepository) product.ProductUsecase {
	return &ProductUsecase{
		productRepo: repo,
		uploadRepo:  uploadRepo,
		trendsRepo: trendsRepo,
	}
}

func (pu *ProductUsecase) Create(product *models.ProductData) *errors.Error {
	product.Date = time.Now().UTC().Format("2006-01-02")

	err := pu.productRepo.Insert(product)
	if err != nil {
		return errors.UnexpectedInternal(err)
	}

	return nil
}
  
func (pu *ProductUsecase) Close(productID uint64, ownerID uint64) *errors.Error {
	product, errE := pu.GetByID(productID)
	if errE != nil {
		return errE
	}

	if product.OwnerID != ownerID {
		return errors.Cause(errors.WrongOwner)
	}

	err := pu.productRepo.Close(product)
	if err != nil {
		return errors.UnexpectedInternal(err)
	}

	return nil
}
  
func (pu *ProductUsecase) Edit(product *models.ProductData) *errors.Error {
	err := pu.productRepo.Update(product)
	if err != nil {
		return errors.UnexpectedInternal(err)
	}

	return nil
}

func (pu *ProductUsecase) UpdatePhoto(productID uint64, ownerID uint64, filesHeaders []*multipart.FileHeader) (*models.ProductData, *errors.Error) {
	product, errE := pu.GetByID(productID)
	if errE != nil {
		return nil, errE
	}

	if product.OwnerID != ownerID {
		return nil, errors.Cause(errors.WrongOwner)
	}

	imgUrls, err := pu.uploadRepo.InsertPhotos(filesHeaders, "static/product/")
	if err != nil {
		return nil, errors.UnexpectedInternal(err)
	}

//	oldPhotos := product.LinkImages
	product.LinkImages = imgUrls
	err = pu.productRepo.InsertPhoto(product)
	if err != nil {
		return nil, errors.UnexpectedInternal(err)
	}

/*	err = pu.uploadRepo.RemovePhotos(oldPhotos)
	if err != nil {
		return nil, errors.UnexpectedInternal(err)
	}
*/
	return product, nil
}

func (pu *ProductUsecase) GetProduct(productID uint64) (*models.ProductData, *errors.Error) {
	product, errE := pu.GetByID(productID)
	if errE != nil {
		return nil, errE
	}

	err := pu.productRepo.UpdateProductViews(productID, 1)
	if err != nil {
		return nil, errors.UnexpectedInternal(err)
	}

	return product, nil
}

func (pu *ProductUsecase) TrendList(userID *uint64) ([]*models.ProductListData, *errors.Error) {
	productIdArray, err := pu.trendsRepo.GetTrendsProducts(*userID)

	products, err := pu.productRepo.SelectTrands(productIdArray, userID)
	if err != nil {
		return nil, errors.UnexpectedInternal(err)
	}

	if len(products) == 0 {
		return []*models.ProductListData{}, nil
	}

	return products, nil
}

func (pu *ProductUsecase) ListLatest(userID *uint64, content *models.Page) ([]*models.ProductListData, *errors.Error) {
	products, err := pu.productRepo.SelectLatest(userID, content)
	if err != nil {
		return nil, errors.UnexpectedInternal(err)
	}

	if len(products) == 0 {
		return []*models.ProductListData{}, nil
	}

	return products, nil
}

func (pu *ProductUsecase) UserAdList(userId uint64, content *models.Page) ([]*models.ProductListData, *errors.Error) {
	products, err := pu.productRepo.SelectUserAd(userId, content)
	if err != nil {
		return nil, errors.UnexpectedInternal(err)
	}

	if len(products) == 0 {
		return []*models.ProductListData{}, nil
	}

	return products, nil
}

func (pu *ProductUsecase) GetUserFavorite(userID uint64, content *models.Page) ([]*models.ProductListData, *errors.Error) {
	products, err := pu.productRepo.SelectUserFavorite(userID, content)
	if err != nil {
		return nil, errors.UnexpectedInternal(err)
	}

	if len(products) == 0 {
		return []*models.ProductListData{}, nil
	}

	return products, nil
}

func (pu *ProductUsecase) LikeProduct(userID uint64, productID uint64) *errors.Error {
	err := pu.productRepo.InsertProductLike(userID, productID)
	if err != nil {
		return errors.UnexpectedInternal(err)
	}

	if pgErr, ok := err.(pgx.PgError); ok && pgErr.Code == "23505" {
		return errors.Cause(errors.ProductAlreadyLiked)
	}

	err = pu.productRepo.UpdateProductLikes(productID, +1)
	if err != nil {
		return errors.UnexpectedInternal(err)
	}

	return nil
}

func (pu *ProductUsecase) DislikeProduct(userID uint64, productID uint64) *errors.Error {
	err := pu.productRepo.DeleteProductLike(userID, productID)
	if err != nil {
		return errors.UnexpectedInternal(err)
	}

	err = pu.productRepo.UpdateProductLikes(productID, -1)
	if err != nil {
		return errors.UnexpectedInternal(err)
	}

	return nil
}

func (pu *ProductUsecase) SetTariff(productID uint64, tariff int) *errors.Error {
	err := pu.productRepo.UpdateTariff(productID, tariff)
	if err != nil {
		return errors.UnexpectedInternal(err)
	}

	return nil
}

func (pu *ProductUsecase) GetByID(productID uint64) (*models.ProductData, *errors.Error) {
	product, err := pu.productRepo.SelectByID(productID)
	switch {
	case err == sql.ErrNoRows:
		return nil, errors.Cause(errors.ProductNotExist)
	case err != nil:
		return nil, errors.UnexpectedInternal(err)
	}

	return product, nil
}
