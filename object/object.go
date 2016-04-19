package object

import "strconv"

// StaticAddr returns an object's url given its id.
func StaticAddr(uuid string) string {
	id, err := strconv.ParseInt(uuid[5:], 10, 32)
	if err == nil {
		if a := fromStorageStaticAddr(uint32(id), true); a != "" {
			return a
		}
	}
	return uuid
}

// ID returns an object's id given its url.
func ID(url string) uint32 {
	if id := fromStorageID(url, true); id != 0 {
		return id
	}
	return 1
}
