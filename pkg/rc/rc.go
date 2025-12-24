package rc

import (
	"errors"

	ttpb "github.com/chrisfenner/tpm.tools/proto"
	"github.com/google/go-tpm/tpm2"
)

var (
	// ErrNotFound indicates that the requested RC couldn't be found.
	ErrNotFound = errors.New("RC not found")
)

// LookupResponseCodeByValue looks up the response code and returns it in proto
// form ready to put on the wire.
// Currently, only 0 or 1 results are returned.
func LookupResponseCodeByValue(value int32) ([]*ttpb.ReturnCodeLookupResult, error) {

	desc := tpm2.TPMRC(value).Error()

	return []*ttpb.ReturnCodeLookupResult{
		&ttpb.ReturnCodeLookupResult{
			Description: desc,
		},
	}, nil
}
