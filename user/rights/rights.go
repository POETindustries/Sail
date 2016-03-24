package rights

import "errors"

type Domain uint16
type Mode uint8

const (
	Users       Domain = 0x0000
	Maintenance Domain = 0x0001
)
const DomainCount = 2

const (
	Create Mode = 0x01
	Update Mode = 0x02
	Delete Mode = 0x04
)

var paths = [][]string{
	[]string{"/office/users"},
	[]string{"/office/settings"}}

// Permission is a lightweight datastructure for passing
// rights information around between packages.
type Permission struct {
	id    Domain
	mode  Mode
	paths []string
}

// New returns a fresh Permission object, initialized with the values
// of the given domain and mode.
func New(domain Domain, mode Mode) *Permission {
	return &Permission{
		id:    domain,
		mode:  mode,
		paths: paths[domain]}
}

// C returns true if the Permission includes create actions.
func (p *Permission) C() bool {
	return p.mode|Create == p.mode
}

// R returns true if the Permission includes read actions.
func (p *Permission) R() bool {
	return p.mode != 0
}

// U returns true if the Permission includes update actions.
func (p *Permission) U() bool {
	return p.mode|Update == p.mode
}

// D returns true if the Permission includes delete actions.
func (p *Permission) D() bool {
	return p.mode|Delete == p.mode
}

// Dom returns the domain that the given path belongs to and an
// error that is not nil if the path does not belong in any known
// permission domain.
func Dom(path string) (d Domain, err error) {
	for ; d < DomainCount; d++ {
		for _, p := range paths[d] {
			if p == path {
				return
			}
		}
	}
	return 0, errors.New("No domain found")
}
