package paste

import "net/http"

type API struct{}

// Get godoc
//
//	@summary        Get paste
//	@description    Get paste by ID
//	@tags           Pastes API
//	@accept         json
//	@produce        json
//	@param          id	path        string  true    "Paste ID"
//	@success        200 {object}    ResponseDTO
//	@failure        400 {object}    err.Error
//	@failure        404
//	@failure        500 {object}    err.Error
//	@router         /v1/pastes/{id} [get]
func (a *API) Get(w http.ResponseWriter, r *http.Request) {}

// Create godoc
//
//	@summary		Create paste
//	@description	Create paste
//	@tags			Pastes API
//	@accept			json
//	@produce		json
//	@param			body	body	RequestDTO	true	"Create Paste DTO"
//	@success		201 {object}	ResponseDTO
//	@failure		400	{object}	err.Error
//	@failure		422	{object}	err.Errors
//	@failure		500	{object}	err.Error
//	@router			/v1/pastes [post]
func (a *API) Create(w http.ResponseWriter, r *http.Request) {}

// Update godoc
//
//	@summary        Update paste
//	@description    Update paste
//	@tags           Pastes API
//	@accept         json
//	@produce        json
//	@param          id      path    string  true    "Paste ID"
//	@param          body    body    RequestDTO    true    "Update Paste DTO"
//	@success        200 {object}    ResponseDTO
//	@failure        400 {object}    err.Error
//	@failure        404
//	@failure        422 {object}    err.Errors
//	@failure        500 {object}    err.Error
//	@router         /v1/pastes/{id} [put]
func (a *API) Update(w http.ResponseWriter, r *http.Request) {}

// Delete godoc
//
//	@summary        Delete paste
//	@description    Delete paste
//	@tags           Pastes API
//	@accept         json
//	@produce        json
//	@param          id  path    string  true    "Paste ID"
//	@success        200
//	@failure        400 {object}    err.Error
//	@failure        404
//	@failure        500 {object}    err.Error
//	@router         /v1/pastes/{id} [delete]
func (a *API) Delete(w http.ResponseWriter, r *http.Request) {}
