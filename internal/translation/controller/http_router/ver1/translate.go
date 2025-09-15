package ver1

import (
	"context"
	"net/http"

	"github.com/amagkn/translabor/internal/translation/dto"
	"github.com/amagkn/translabor/pkg/base_errors"
	"github.com/amagkn/translabor/pkg/logger"
	"github.com/amagkn/translabor/pkg/validation"
)

func (h *Handlers) Translate(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	input := dto.TranslateInput{}

	invalidFields, err := validation.ValidateStructWithDecodeJSONBody(r.Body, &input)
	if err != nil {
		logger.Error(err, "validation.ValidateStructWithDecodeJSONBody")
		if invalidFields != nil {
			errorResponse(w, http.StatusBadRequest, errorPayload{Type: base_errors.Validation, Details: invalidFields})
		} else {
			errorResponse(w, http.StatusBadRequest, errorPayload{Type: base_errors.InvalidJSON})
		}
		return
	}

	output, err := h.uc.Translate(ctx, input)
	if err != nil {
		logger.Error(err, "h.uc.Translate")
		errorResponse(w, http.StatusBadRequest, errorPayload{Type: base_errors.InternalServer})
		return
	}

	successResponse(w, http.StatusOK, output)
}
