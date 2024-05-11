package usecase

import (
	"context"

	"github.com/RizqiSugiarto/GaleryService/internal/entity"
)

type GalleryUsecase struct {
	gr GalleryRepo
}

func New(gr GalleryRepo) *GalleryUsecase {
	return &GalleryUsecase{gr: gr}
}

func(g *GalleryUsecase) CreateLink(ctx context.Context, userData entity.Gallery) error {
	return g.gr.Create(ctx, userData)
}

func(g *GalleryUsecase) UpdateLink(ctx context.Context, userId string ,userData entity.Gallery) error {
	return g.gr.Update(ctx, userId, userData)
}

func(g *GalleryUsecase) GetLinkByUserId(ctx context.Context, userId string) (entity.Gallery, error) {
	return g.gr.GetById(ctx, userId)
}

func(g *GalleryUsecase) DeleteLink(ctx context.Context, userId string) error {
	return g.gr.Delete(ctx, userId)
}