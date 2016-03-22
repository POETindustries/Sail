package group

type PermissionDomain uint32
type PermissionMode uint8

const (
	PermUsers       PermissionDomain = 0x00000001
	PermMaintenance PermissionDomain = 0x00000002
)

const (
	ModeCreate PermissionMode = 0x01
	ModeUpdate PermissionMode = 0x02
	ModeDelete PermissionMode = 0x04
)

type Permission struct {
	id   PermissionDomain
	mode PermissionMode
}

func (p *Permission) Cr() bool {
	return p.mode|ModeCreate == p.mode
}

func (p *Permission) U() bool {
	return p.mode|ModeUpdate == p.mode
}

func (p *Permission) D() bool {
	return p.mode|ModeDelete == p.mode
}
