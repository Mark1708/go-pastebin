package mapper

import (
	"time"

	"github.com/Mark1708/go-pastebin/internal/paste/dto/request"
	"github.com/Mark1708/go-pastebin/internal/paste/dto/response"
	"github.com/Mark1708/go-pastebin/internal/paste/dto/visibility"
	"github.com/Mark1708/go-pastebin/internal/paste/models"
)

func UpdateFromDTO(p *models.Paste, dto request.Dto) {
	if dto.Title != "" {
		p.Title = dto.Title
	}

	if dto.Visibility != "" {
		p.Visibility = dto.Visibility
	}

	if dto.Content != "" {
		p.Content = dto.Content
	}
}

func EntityToResponseDto(paste models.Paste) (response.Dto, error) {
	visibilityType, typeError := visibility.TypeValueOf(paste.Visibility)
	if typeError != nil {
		return response.Dto{}, typeError
	}

	return response.Dto{
		Hash:  paste.Hash,
		Title: paste.Title,
		Visibility: visibility.Dto{
			Type:  visibilityType.String(),
			Title: visibilityType.Title(),
		},
		CreatedAt: paste.CreatedAt.Format(time.RFC3339),
		ExpiredAt: paste.ExpiredAt.Format(time.RFC3339),
		Content:   paste.Content,
	}, nil
}

func RequestDtoToEntity(hash string, dto request.Dto) models.Paste {
	now := time.Now()
	return models.Paste{
		Hash:       hash,
		Title:      dto.Title,
		Visibility: dto.Visibility,
		CreatedAt:  now,
		ExpiredAt:  now.Add(time.Hour),
		Content:    dto.Content,
	}
}
