package group

import "fmt"

type cache struct {
	groups []*Group
}

var cacheInstance *cache

func Cache() *cache {
	if cacheInstance == nil {
		cacheInstance = &cache{groups: fromStorageByID()}
		for _, g := range cacheInstance.groups {
			fmt.Printf("%+v", g)
		}
	}
	return cacheInstance
}
