package issuereferencechaining_test

import (
	"testing"

	"github.com/deepmap/oapi-codegen/internal/test/issues/issue-1223/api"
)

func TestIssue(t *testing.T) {
	x := api.GetItemRequestObject{}
	// no-op bc it will pass if ./api and ./item compile without errors
}
