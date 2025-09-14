package ver1

import (
	"github.com/amagkn/translabor/internal/translation/usecase"
)

type Handlers struct {
	uc *usecase.UseCase
}

func New(uc *usecase.UseCase) *Handlers {
	return &Handlers{uc: uc}
}
