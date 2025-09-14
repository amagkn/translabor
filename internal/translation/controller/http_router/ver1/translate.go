package ver1

import (
	"net/http"

	"github.com/amagkn/translabor/internal/translation/dto"
	"github.com/amagkn/translabor/pkg/base_errors"
	"github.com/amagkn/translabor/pkg/logger"
	"github.com/amagkn/translabor/pkg/validation"
)

func (h *Handlers) Translate(w http.ResponseWriter, r *http.Request) {
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

	var output dto.TranslateOutput
	translation, err := h.uc.Translate(input)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, errorPayload{Type: base_errors.InternalServer})
		return
	}

	output.Translation = translation

	successResponse(w, http.StatusOK, output)
}
