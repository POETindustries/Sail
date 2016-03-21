package group

const (
	PermUsers       = 0x00000001
	PermMaintenance = 0x00000002
)

type Permission struct {
	id     uint32
	create bool
	update bool
	delete bool
}

func (p *Permission) Cr() bool {
	return p.create
}

func (p *Permission) U() bool {
	return p.update
}

func (p *Permission) D() bool {
	return p.delete
}
