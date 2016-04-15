package object

import "strconv"

func StaticAddr(uuid string) string {
	id, err := strconv.ParseInt(uuid[5:], 10, 32)
	if err == nil {
		if a := fromStorageStaticAddr(uint32(id), true); a != "" {
			return a
		}
	}
	return uuid
}
