package main

import (
	"net/http"
	"sync"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/jinzhu/gorm"
)

//IChiRouter interface for router
type IChiRouter interface {
	InitRouter() *chi.Mux
}

type router struct {
	db *gorm.DB
}

func (router *router) InitRouter() *chi.Mux {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root."))
	})

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("test")
	})

	container := NewInjectionContainer(router.db)

	productController := container.ProvideProductController()

	r.Route("/products", func(r chi.Router) {
		r.Get("/{tenantId:[0-9]+}", productController.GetByTenetID)
	})

	return r
}

var (
	m          *router
	routerOnce sync.Once
)

//ChiRouter singleton function to return router
func ChiRouter(db *gorm.DB) IChiRouter {
	if m == nil {
		routerOnce.Do(func() {
			m = &router{db: db}
		})
	}

	return m
}
