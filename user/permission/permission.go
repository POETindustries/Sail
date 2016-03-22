package permission

type Domain uint16
type Mode uint8

const (
	Users       Domain = 0x0000
	Maintenance Domain = 0x0001
)

const (
	Create Mode = 0x01
	Update Mode = 0x02
	Delete Mode = 0x04
)

// Permission is a lightweight datastructure for passing
// permission information around between packages.
type Permission struct {
	id   Domain
	mode Mode
}

func (p *Permission) Cr() bool {
	return p.mode|Create == p.mode
}

func (p *Permission) U() bool {
	return p.mode|Update == p.mode
}

func (p *Permission) D() bool {
	return p.mode|Delete == p.mode
}
