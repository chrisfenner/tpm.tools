package jsonproto

import (
	"encoding/json"
	"errors"
	"fmt"

	ttpb "github.com/chrisfenner/tpm.tools/proto"
)

var (
	ErrJSONParsingFailure = errors.New("failed to parse JSON")
)

func LoadCommandProtos(jsonData []byte) (map[string]*ttpb.CommandDescription, error) {
	list := make(map[string]*ttpb.CommandDescription)

	// Unmarshal the whole file
	if err := json.Unmarshal(jsonData, &list); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrJSONParsingFailure, err)
	}

	return list, nil
}
