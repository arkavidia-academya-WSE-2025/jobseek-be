package repository

import (
	"fp-academya-be/internal/entity"
	"fp-academya-be/internal/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PostRepository struct {
	Repository[entity.Post]
	Log *logrus.Logger
}

func NewPostRepository(log *logrus.Logger) *PostRepository {
	return &PostRepository{
		Log: log,
	}
}

func (r *PostRepository) FindById(db *gorm.DB, post *entity.Post, id string) error {
	return db.Preload("User").Where("id = ?", id).Take(post).Error
}

func (r *PostRepository) Search(db *gorm.DB, request *model.SearchPostRequest) ([]entity.Post, int64, error) {
	var posts []entity.Post
	if err := db.Scopes(r.FilterPost(request)).Preload("User").Offset((request.Page - 1) * request.Size).Limit(request.Size).Find(&posts).Error; err != nil {
		return nil, 0, err
	}

	var total int64 = 0
	if err := db.Model(&entity.Post{}).Scopes(r.FilterPost(request)).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return posts, total, nil
}

func (r *PostRepository) FilterPost(request *model.SearchPostRequest) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {

		if title := request.Title; title != "" {
			tx = tx.Where("title LIKE ?", "%"+title+"%")
		}

		if content := request.Content; content != "" {
			tx = tx.Where("content LIKE ?", "%"+content+"%")
		}

		return tx
	}
}
