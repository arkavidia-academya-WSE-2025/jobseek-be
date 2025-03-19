package converter

import (
	"fp-academya-be/internal/entity"
	"fp-academya-be/internal/model"
)

func PostToResponse(post *entity.Post) *model.PostReponse {
	return &model.PostReponse{
		ID:        &post.ID,
		Title:     post.Title,
		Content:   post.Content,
		CreatedAt: &post.CreatedAt,
		UpdatedAt: &post.UpdatedAt,
		User:      UserToResponse(&post.User),
	}
}
