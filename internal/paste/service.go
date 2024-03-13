package paste

import (
	"github.com/Mark1708/go-pastebin/internal/paste/dto/request"
	"github.com/Mark1708/go-pastebin/internal/paste/dto/response"
	"github.com/Mark1708/go-pastebin/internal/pkg/common"
)

type Service interface {
	GetByHash(hash string) (response.Dto, common.ErrResponseDto)

	Create(remoteAddr string, requestDTO request.Dto) (response.Dto, common.ErrResponseDto)

	Update(hash string, requestDTO request.Dto) (response.Dto, common.ErrResponseDto)

	Delete(hash string) (common.OperationResponseDto, common.ErrResponseDto)
}
