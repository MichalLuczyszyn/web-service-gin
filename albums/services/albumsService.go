package services

import (
	"context"
	"github.com/MichalLuczyszyn/web-service-gin/albums/domain"
)

type AlbumsRepository interface {
	Fetch(ctx context.Context, cursor string, num int64) (res []domain.album, nextCursor string, err error)
	GetById(ctx context.Context, id int64) (domain.album, error)
	Update(ctx context.Context, ar *domain.album) error
	Delete(ctx context.Context, id int64) error
}
