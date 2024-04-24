package paste

import (
	"context"
	"strings"

	"github.com/Mark1708/go-pastebin/pkg/ip"

	"github.com/Mark1708/go-pastebin/internal/mapper"
	model "github.com/Mark1708/go-pastebin/internal/model/paste"
	repo "github.com/Mark1708/go-pastebin/internal/repository/paste"
	hasher "github.com/Mark1708/go-pastebin/pkg/hash"
	"github.com/Mark1708/go-pastebin/pkg/logger"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Service interface {
	GetPasteByHash(ctx context.Context, hash string) (model.ResponseDto, error)

	CreatePaste(ctx context.Context, reqDto model.RequestDto) (model.ResponseDto, error)

	UpdatePaste(ctx context.Context, hash string, reqDto model.RequestDto) (model.ResponseDto, error)

	DeletePaste(ctx context.Context, hash string) error
}

type service struct {
	repository repo.Repository
}

func NewService(repository repo.Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetPasteByHash(
	ctx context.Context,
	hash string,
) (model.ResponseDto, error) {
	if validateErr := validateHash(hash); validateErr != nil {
		return model.ResponseDto{}, errors.New("")
	}

	entity, getPasteError := s.repository.GetPasteByHash(ctx, hash)
	if getPasteError != nil {
		return model.ResponseDto{}, errors.New("")
	}

	respDto, typeError := mapper.EntityToResponseDto(entity)
	if typeError != nil {
		return model.ResponseDto{}, errors.New("")
	}

	return respDto, nil
}

func (s *service) CreatePaste(
	ctx context.Context,
	reqDto model.RequestDto,
) (model.ResponseDto, error) {
	remoteAddr, ok := ctx.Value(ip.UserIPKey{}).(string)
	if !ok {
		return model.ResponseDto{}, errors.New("remote ip is required")
	}
	if remoteAddr == "" {
		return model.ResponseDto{}, errors.New("")
	}
	hash := hasher.GenerateHash(remoteAddr)
	newEntity := mapper.PasteDtoToEntity(hash, reqDto)

	entity, createPasteErr := s.repository.CreatePaste(ctx, newEntity)
	if createPasteErr != nil {
		logger.Log.With(zap.Error(createPasteErr)).Error("error during save paste in db")
		return model.ResponseDto{}, errors.New("")
	}

	respDto, typeError := mapper.EntityToResponseDto(entity)
	if typeError != nil {
		return model.ResponseDto{}, errors.New("")
	}

	return respDto, nil
}

func (s *service) UpdatePaste(
	ctx context.Context,
	hash string,
	reqDto model.RequestDto,
) (model.ResponseDto, error) {
	if validateErr := validateHash(hash); validateErr != nil {
		return model.ResponseDto{}, errors.New("")
	}

	dbEntity, getErr := s.repository.GetPasteByHash(ctx, hash)
	if getErr != nil {
		return model.ResponseDto{}, errors.New("")
	}
	mapper.UpdateFromDTO(&dbEntity, reqDto)

	entity, updatePasteErr := s.repository.UpdatePaste(ctx, dbEntity)
	if updatePasteErr != nil {
		return model.ResponseDto{}, errors.New("")
	}

	respDto, typeError := mapper.EntityToResponseDto(entity)
	if typeError != nil {
		return model.ResponseDto{}, errors.New("")
	}

	return respDto, nil
}

func (s *service) DeletePaste(
	ctx context.Context,
	hash string,
) error {
	deleteErr := s.repository.DeletePaste(ctx, hash)
	if deleteErr != nil {
		return errors.New("")
	}
	return nil
}

func validateHash(hash string) error {
	// Игнорируем статичные файлы которые запрашивает браузер
	if strings.Contains(hash, ".") {
		return errors.New("hash not include point letter")
	}
	return nil
}
