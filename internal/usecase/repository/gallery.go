package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/RizqiSugiarto/GaleryService/internal/entity"
	"github.com/RizqiSugiarto/GaleryService/pkg/maria"
	"github.com/google/uuid"
)

type GalleryRepo struct {
	*maria.Maria
}

func New(ma *maria.Maria) *GalleryRepo {
	return &GalleryRepo{ma}
}

func(g *GalleryRepo) Create(ctx context.Context, gl entity.Gallery) error {
	sql, args, err := g.Builder.
			Insert("gallery").Columns("id, link, userId").
			Values(uuid.New(), gl.Link, gl.UserId).
			ToSql()
		
		if err != nil {
			return fmt.Errorf("galleryRepo - Create - g.Builder")
		}

		_, err = g.Db.Exec(sql, args...)

		if err != nil {
			return fmt.Errorf("galleryRepo - Create - g.Db.Exec")
		}

		return nil
}

func(g *GalleryRepo) Update(ctx context.Context, userId string, gl entity.Gallery) error {
	updateValue := map[string]interface{}{
		"link": gl.Link,
		"userId": gl.UserId,
		"updated_at": time.Now(),
	}

	sql, args, err := g.Builder.
			Update("gallery").
			SetMap(updateValue).
			ToSql()
		
		if err != nil {
			return fmt.Errorf("galleryrepo - Update - g.Builder: %v", err)
		}

		_, err = g.Db.Exec(sql, args...)

		if err != nil {
			return fmt.Errorf("galleryrepo - update - g.Db.Exec: %v", err)
		}

		return nil
}

func(g *GalleryRepo) GetById(ctx context.Context, userId string) (entity.Gallery, error) {
	
	var userData entity.Gallery
	
	sql, args, err := g.Builder.
			Select("*").
			From("gallery").
			Where(squirrel.Eq{"userId": userId}).
			ToSql()
		
		if err != nil {
			return entity.Gallery{}, fmt.Errorf("galleryrepo - GetById - g.Builder: %v", err)
		}

		rows, err := g.Db.Query(sql, args...)

		if err != nil {
			return entity.Gallery{}, fmt.Errorf("galleryrepo - GetById - g.Db.Query: %v", err)
		}

		if err := rows.Scan(&userData); err != nil {
			return entity.Gallery{}, fmt.Errorf("galleryrepo - GetById - rows.Scan: %v", err) 
		}

		return userData, nil
}

func(g *GalleryRepo) Delete(ctx context.Context, userId string) error {
	sql, args, err := g.Builder.
			Delete("gallery").
			Where(squirrel.Eq{"userId": userId}).
			ToSql()
			
		if err != nil {
			return fmt.Errorf("galleryrepo - Delete - g.Builder: %v", err)
		}

		_, err = g.Db.Exec(sql, args...)

		if err != nil {
			return fmt.Errorf("galleryrepo - Delete - g.Db.Exec: %v", err)
		}
		
		return nil
}