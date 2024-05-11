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
			Insert("gallery").Columns("id, link, user_id, created_at, updated_at").
			Values(uuid.New(), gl.Link, gl.UserId, time.Now(), time.Now()).
			ToSql()
		
		if err != nil {
			return fmt.Errorf("galleryRepo - Create - g.Builder: %v", err)
		}

		_, err = g.Db.Exec(sql, args...)
		if err != nil {
			return fmt.Errorf("galleryRepo - Create - g.Db.Exec: %v", err)
		}

		return nil
}

func (g *GalleryRepo) Update(ctx context.Context, userId string, gl entity.Gallery) error {
    tx, err := g.Db.Begin()

    if err != nil {
        return fmt.Errorf("galleryrepo - Update - begin transaction: %v", err)
    }

    defer func() {
        if err != nil {
            tx.Rollback()
        }
    }()

    var exists bool

    sqlExist, _, err := g.Builder.
        Select("COUNT(*)").
        From("gallery").
        Where(squirrel.Eq{"user_id": userId}).
        Limit(1).
        ToSql()

    err = tx.QueryRowContext(ctx, sqlExist, userId).Scan(&exists)

    if err != nil {
        return fmt.Errorf("galleryrepo - Update - checking user existence: %v", err)
    }

    if !exists {
        return fmt.Errorf("data not found")
    }

    updateValue := map[string]interface{}{
        "link":       gl.Link,
        "updated_at": time.Now(),
    }

    sqlUpdate, args, err := g.Builder.
        Update("gallery").
        SetMap(updateValue).
        Where(squirrel.Eq{"user_id": userId}).
        ToSql()

    if err != nil {
        return fmt.Errorf("galleryrepo - Update - g.Builder: %v", err)
    }
    _, err = tx.ExecContext(ctx, sqlUpdate, args...)

    if err != nil {
        return fmt.Errorf("galleryrepo - Update - tx.Exec: %v", err)
    }

    if err := tx.Commit(); err != nil {
        return fmt.Errorf("galleryrepo - Update - commit transaction: %v", err)
    }

    return nil
}


func(g *GalleryRepo) GetById(ctx context.Context, userId string) (entity.Gallery, error) {
	
	var userData entity.Gallery
	
	sql, args, err := g.Builder.
			Select("*").
			From("gallery").
			Where(squirrel.Eq{"user_id": userId}).
			ToSql()
		
		if err != nil {
			return entity.Gallery{}, fmt.Errorf("galleryrepo - GetById - g.Builder: %v", err)
		}

		rows:= g.Db.QueryRow(sql, args...)

		if err := rows.Scan(&userData.Id, &userData.Link, &userData.UserId, &userData.Created_at, &userData.Updated_at); err != nil {
			return entity.Gallery{}, err 
		}

		return userData, nil
}

func(g *GalleryRepo) Delete(ctx context.Context, userId string) error {
	tx, err := g.Db.Begin()

    if err != nil {
        return fmt.Errorf("galleryrepo - Delete - begin transaction: %v", err)
    }

    defer func() {
        if err != nil {
            tx.Rollback()
        }
    }()

    var exists bool

    sqlExist, _, err := g.Builder.
        Select("COUNT(*)").
        From("gallery").
        Where(squirrel.Eq{"user_id": userId}).
        Limit(1).
        ToSql()

    err = tx.QueryRowContext(ctx, sqlExist, userId).Scan(&exists)

    if err != nil {
        return fmt.Errorf("galleryrepo - Delete - checking user existence: %v", err)
    }

    if !exists {
        return fmt.Errorf("data not found")
    }

	sql, args, err := g.Builder.
			Delete("gallery").
			Where(squirrel.Eq{"user_id": userId}).
			ToSql()
			
		if err != nil {
			return fmt.Errorf("galleryrepo - Delete - g.Builder: %v", err)
		}

		_, err = tx.ExecContext(ctx, sql, args...)

		if err != nil {
			return fmt.Errorf("galleryrepo - Delete - g.Db.Exec: %v", err)
		}

		if err := tx.Commit(); err != nil {
			return fmt.Errorf("galleryrepo - Delete - commit transaction: %v", err)
		}
		
		return nil
}