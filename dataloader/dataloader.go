package dataloader

import (
	"context"
	"myapp/service"
	"net/http"
	"time"
)

const loadersKey = "dataloaders"

type Loaders struct {
	ItemByUserID ItemByUserIDLoader
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), loadersKey, &Loaders{
			ItemByUserID: ItemByUserIDLoader{
				maxBatch: 100,
				wait:     1 * time.Millisecond,
				fetch:    service.ItemDataloaderByUserID,
			},
		})
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
