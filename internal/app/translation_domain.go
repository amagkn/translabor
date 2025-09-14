package app

import (
	"github.com/amagkn/translabor/internal/translation/controller/http_router"
	"github.com/amagkn/translabor/internal/translation/usecase"
)

func ProductDomain(d Dependences) {
	translationUseCase := usecase.New()

	http_router.TranslationRoutes(d.RouterHTTP, translationUseCase)
}
