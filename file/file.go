package file

const (
	Private = 0
	Public  = 1
)

func StaticAddr(uuid string) []byte {
	if a := fromStorageGetAddr(uuid, true); a != "" {
		return []byte(a)
	}
	return []byte(uuid)
}
