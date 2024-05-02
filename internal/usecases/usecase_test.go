package usecases_test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"reddit/internal/entities"
	"reddit/internal/repositories/mocks"
	"reddit/internal/usecases"
	"testing"
)

func TestCreatePost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockPostRepository(ctrl)

	mockRepo.EXPECT().CreatePost("Test Title", "Test Content", true).Return(&entities.Post{
		ID:            1,
		Title:         "Test Title",
		Content:       "Test Content",
		AllowComments: true,
	}, nil)

	postUseCase := usecases.NewPostUseCase(mockRepo)

	result, err := postUseCase.CreateNewPost("Test Title", "Test Content", true)

	assert.NoError(t, err)
	assert.Equal(t, "Test Title", result.Title)
	assert.Equal(t, "Test Content", result.Content)
	assert.True(t, result.AllowComments)
}
