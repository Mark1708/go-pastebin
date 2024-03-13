package visibility_test

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/Mark1708/go-pastebin/internal/paste/dto/visibility"
)

func TestEnumIndex(t *testing.T) {
	visibilityType := visibility.PUBLIC
	result := visibilityType.EnumIndex()

	if result != 0 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 0)
	}
}

func TestString(t *testing.T) {
	visibilityType := visibility.PUBLIC
	result := visibilityType.String()

	if result != "PUBLIC" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "PUBLIC")
	}
}

func TestTitle(t *testing.T) {
	visibilityType := visibility.PUBLIC
	result := visibilityType.Title()

	if result != "Public" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "Public")
	}
}

func TestTypeValueOf(t *testing.T) {
	result, err := visibility.TypeValueOf("PUBLIC")

	if err != nil {
		t.Error(err)
	}

	if result != visibility.PUBLIC {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, visibility.PUBLIC.String())
	}
}

func TestMarshalJSON(t *testing.T) {
	blob := `["PUBLIC","UNLISTED","PRIVATE","PUBLIC","PRIVATE","UNLISTED"]`
	var visibilityTypes []visibility.Type
	if err := json.Unmarshal([]byte(blob), &visibilityTypes); err != nil {
		log.Fatal(err)
	}

	census := make(map[visibility.Type]int)
	for _, visibilityType := range visibilityTypes {
		census[visibilityType]++
	}

	t.Logf("Type Census:\n* Public: %d\n* Private:  %d\n* Unlisted: %d\n",
		census[visibility.PUBLIC], census[visibility.PRIVATE], census[visibility.UNLISTED])

	// Output:
	// Type Census:
	// * Public: 2
	// * Private:  2
	// * Unlisted: 2
}

func TestUnmarshalJSON(t *testing.T) {
	visibilityTypes := []visibility.Type{
		visibility.PUBLIC, visibility.UNLISTED, visibility.PRIVATE,
		visibility.PUBLIC, visibility.PRIVATE, visibility.UNLISTED,
	}
	bytes, err := json.Marshal(visibilityTypes)
	if err != nil {
		log.Fatal(err)
	}
	jsonRes := string(bytes)
	t.Log(jsonRes)

	// Output:
	// ["PUBLIC","UNLISTED","PRIVATE","PUBLIC","PRIVATE","UNLISTED"]
}
