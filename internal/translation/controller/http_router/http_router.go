package http_router

import (
	"github.com/amagkn/translabor/internal/translation/controller/http_router/ver1"
	"github.com/amagkn/translabor/internal/translation/usecase"
	"github.com/go-chi/chi/v5"
)

func TranslationRoutes(r *chi.Mux, uc *usecase.UseCase) {
	v1 := ver1.New(uc)

	r.Route("/api/v1/translate", func(r chi.Router) {
		r.Post("/", v1.Translate)
	})
}
