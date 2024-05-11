package usecase

import (
	"context"

	"github.com/RizqiSugiarto/GaleryService/internal/entity"
)

type (
	Gallery interface {
		CreateLink(ctx context.Context, userData entity.Gallery) error 
		UpdateLink(ctx context.Context, userId string ,userData entity.Gallery) error
		GetLinkByUserId(ctx context.Context, userId string) (entity.Gallery, error)
		DeleteLink(ctx context.Context, userId string) error 
	}
	GalleryRepo interface {
		Create(ctx context.Context, gl entity.Gallery) error
		Update(ctx context.Context, userId string, gl entity.Gallery) error
		GetById(ctx context.Context, userId string) (entity.Gallery, error)
		Delete(ctx context.Context, userId string) error
	}
)

