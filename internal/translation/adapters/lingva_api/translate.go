package lingva_api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/amagkn/translabor/internal/translation/dto"
	"github.com/amagkn/translabor/pkg/base_errors"
	"github.com/amagkn/translabor/pkg/logger"
)

func (l *LingvaAPI) Translate(input dto.TranslateInput) (string, error) {
	lingvaURL := fmt.Sprintf("https://lingva.ml/api/v1/%s/%s/%s", input.Source, input.Target, input.Query)
	logger.Info(fmt.Sprintf("Init request: %s", lingvaURL))

	res, err := http.Get(lingvaURL)
	if err != nil {
		return "", base_errors.WithPath("http.Get", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", base_errors.WithPath("io.ReadAll", err)
	}

	var output dto.LingvaTranslateOutput
	if err := json.Unmarshal(body, &output); err != nil {
		return "", base_errors.WithPath("json.Unmarshal", err)
	}

	return output.Translation, nil
}
