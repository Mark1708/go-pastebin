package service

import (
	"errors"
	"strings"

	"github.com/Mark1708/go-pastebin/internal/paste/mapper"

	"github.com/Mark1708/go-pastebin/internal/paste"
	"github.com/Mark1708/go-pastebin/internal/paste/dto/request"
	"github.com/Mark1708/go-pastebin/internal/paste/dto/response"

	"github.com/Mark1708/go-pastebin/internal/pkg/common"
	hasher "github.com/Mark1708/go-pastebin/internal/pkg/hash"
)

type Service struct {
	Repo paste.Repository
}

func (s *Service) GetByHash(hash string) (response.Dto, common.ErrResponseDto) {
	if validateErr := validateHash(hash); !validateErr.IsEmpty() {
		return response.Dto{}, validateErr
	}

	entity, getPasteError := s.Repo.GetByHash(hash)
	if getPasteError != nil {
		return response.Dto{}, common.ErrNotFound()
	}

	respDto, typeError := mapper.EntityToResponseDto(entity)
	if typeError != nil {
		return response.Dto{}, common.ErrInternalServerError(typeError)
	}

	return respDto, common.ErrResponseDto{}
}

func (s *Service) Create(
	remoteAddr string,
	requestDTO request.Dto,
) (response.Dto, common.ErrResponseDto) {
	hash := hasher.GenerateHash(remoteAddr)
	newEntity := mapper.RequestDtoToEntity(hash, requestDTO)

	entity, createPasteErr := s.Repo.Create(newEntity)
	if createPasteErr != nil {
		return response.Dto{}, common.ErrInternalServerError(createPasteErr)
	}

	respDto, typeError := mapper.EntityToResponseDto(entity)
	if typeError != nil {
		return response.Dto{}, common.ErrInternalServerError(typeError)
	}

	return respDto, common.ErrResponseDto{}
}

func (s *Service) Update(hash string, requestDTO request.Dto) (response.Dto, common.ErrResponseDto) {
	if validateErr := validateHash(hash); !validateErr.IsEmpty() {
		return response.Dto{}, validateErr
	}

	dbEntity, getErr := s.Repo.GetByHash(hash)
	if getErr != nil {
		return response.Dto{}, common.ErrNotFound()
	}
	mapper.UpdateFromDTO(&dbEntity, requestDTO)

	entity, createPasteErr := s.Repo.Update(dbEntity)
	if createPasteErr != nil {
		return response.Dto{}, common.ErrInternalServerError(createPasteErr)
	}

	respDto, typeError := mapper.EntityToResponseDto(entity)
	if typeError != nil {
		return response.Dto{}, common.ErrInternalServerError(typeError)
	}

	return respDto, common.ErrResponseDto{}
}

func (s *Service) Delete(hash string) (common.OperationResponseDto, common.ErrResponseDto) {
	deleteErr := s.Repo.Delete(hash)
	if deleteErr != nil {
		return common.OperationResponseDto{}, common.ErrNotFound()
	}
	return common.OperationResponseDto{Message: "Success"}, common.ErrResponseDto{}
}

func validateHash(hash string) common.ErrResponseDto {
	// Игнорируем статичные файлы которые запрашивает браузер
	if strings.Contains(hash, ".") {
		return common.ErrInvalidRequest(errors.New("hash not include point letter"))
	}
	return common.ErrResponseDto{}
}
