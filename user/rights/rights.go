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

func New(domain Domain, mode Mode) *Permission {
	return &Permission{
		id:    domain,
		mode:  mode,
		paths: paths[domain]}
}

func (p *Permission) C() bool {
	return p.mode|Create == p.mode
}

func (p *Permission) R() bool {
	return p.mode != 0
}

func (p *Permission) U() bool {
	return p.mode|Update == p.mode
}

func (p *Permission) D() bool {
	return p.mode|Delete == p.mode
}

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
