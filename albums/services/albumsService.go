package albums;

import (
	"context"
	"github.com/MichalLuczyszyn/web-service-gin/albums/domain"
)

type AlbumsRepository interface {
	Fetch(ctx context.Context, cursor string, num int64) (res []album, nextCursor string, err error)
	GetById(ctx context.Context, id int64) (album, error)
	Update(ctx context.Context, ar *album) error
	Delete(ctx context.Context, id int64) error
}