package usecase

import (
	"context"
	"testing"

	"github.com/amagkn/translabor/internal/translation/dto"
	"github.com/amagkn/translabor/internal/translation/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockPostgres struct{ mock.Mock }

func (m *MockPostgres) SelectWord(ctx context.Context, query string) (entity.WordWithTranslation, error) {
	args := m.Called(ctx, query)
	return args.Get(0).(entity.WordWithTranslation), args.Error(1)
}
func (m *MockPostgres) InsertWord(ctx context.Context, input dto.SaveWordInput) (entity.WordWithTranslation, error) {
	args := m.Called(ctx, input)
	return args.Get(0).(entity.WordWithTranslation), args.Error(1)
}

type MockLingvaAPI struct{ mock.Mock }

func (m *MockLingvaAPI) Translate(input dto.TranslateInput) (string, error) {
	args := m.Called(input)
	return args.String(0), args.Error(1)
}

func TestUseCase_Translate(t *testing.T) {
	ctx := context.Background()
	input := dto.TranslateInput{Query: "hello"}

	t.Run("word exists in db", func(t *testing.T) {
		mockPostgres := &MockPostgres{}
		mockLingvaAPI := &MockLingvaAPI{}

		mockPostgres.On("SelectWord", ctx, "hello").Return(entity.WordWithTranslation{Translation: "привет"}, nil)

		u := UseCase{postgres: mockPostgres, lingvaAPI: mockLingvaAPI}
		out, err := u.Translate(ctx, input)
		assert.NoError(t, err)
		assert.Equal(t, "привет", out.Translation)
	})
}
