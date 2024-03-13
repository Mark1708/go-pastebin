package visibility

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Type int

const (
	PUBLIC Type = iota
	PRIVATE
	UNLISTED
)

func Names() []string {
	return []string{"PUBLIC", "PRIVATE", "UNLISTED"}
}

func Titles() []string {
	return []string{"Public", "Private", "Unlisted"}
}

func visibilityTypesNameToValueMap() map[string]Type {
	return map[string]Type{
		"PUBLIC":   PUBLIC,
		"PRIVATE":  PRIVATE,
		"UNLISTED": UNLISTED,
	}
}

func TypeValueOf(s string) (Type, error) {
	name := strings.ToUpper(strings.TrimSpace(s))
	if val, ok := visibilityTypesNameToValueMap()[name]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to VisibilityTypes values", s)
}

func (d Type) String() string {
	return Names()[d]
}

func (d Type) Title() string {
	return Titles()[d]
}

func (d Type) EnumIndex() int {
	return int(d)
}

func (d Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d *Type) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("VisibilityTypes should be a string, got %s", data)
	}

	var err error
	*d, err = TypeValueOf(s)
	return err
}

type Dto struct {
	Type  string `json:"type"`
	Title string `json:"title"`
}

func (vtDTO Dto) Render(w http.ResponseWriter, _ *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return nil
}
