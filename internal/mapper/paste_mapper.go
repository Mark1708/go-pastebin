package mapper

import (
	"time"

	"github.com/Mark1708/go-pastebin/internal/model/paste"
	"github.com/Mark1708/go-pastebin/internal/model/visibility"
)

func UpdateFromDTO(p *paste.Paste, dto paste.RequestDto) {
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

func EntityToResponseDto(entity paste.Paste) (paste.ResponseDto, error) {
	visibilityType, typeError := visibility.TypeValueOf(entity.Visibility)
	if typeError != nil {
		return paste.ResponseDto{}, typeError
	}

	return paste.ResponseDto{
		Hash:  entity.Hash,
		Title: entity.Title,
		Visibility: visibility.Dto{
			Type:  visibilityType.String(),
			Title: visibilityType.Title(),
		},
		CreatedAt: entity.CreatedAt.Format(time.RFC3339),
		ExpiredAt: entity.ExpiredAt.Format(time.RFC3339),
		Content:   entity.Content,
	}, nil
}

func PasteDtoToEntity(hash string, dto paste.RequestDto) paste.Paste {
	now := time.Now()
	return paste.Paste{
		Hash:       hash,
		Title:      dto.Title,
		Visibility: dto.Visibility,
		CreatedAt:  now,
		ExpiredAt:  now.Add(time.Hour),
		Content:    dto.Content,
	}
}
