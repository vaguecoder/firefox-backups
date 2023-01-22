package filters

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/vaguecoder/firefox-backups/pkg/bookmark"
)

type Filter interface {
	Apply(context.Context, []bookmark.Bookmark) ([]bookmark.Bookmark, error)
	fmt.Stringer
}

// FilterName holds filter's name
type FilterName string

// String converts FilterName to string,
// making FilterName type implement fmt.Stringer
func (e FilterName) String() string {
	return string(e)
}

// filterNames is an unexported collection of filter names
type filterNames []FilterName

// String converts filters to string,
// making FilterName type implement fmt.Stringer
func (e filterNames) String() string {
	var filters []string

	for _, filter := range e {
		filters = append(filters, filter.String())
	}

	sort.Strings(filters)

	return strings.Join(filters, ", ")
}

// AllFilterNames holds list of filter names.
// All the filter names in echo of the filter packages should
// be appended to AllFilterNames during respective init()
var AllFilterNames filterNames

// ToFilterName converts stringer to FilterName type
func ToFilterName(s fmt.Stringer) FilterName {
	return FilterName(s.String())
}
