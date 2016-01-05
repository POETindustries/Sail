package domains

import (
	"fmt"
	"sail/cache"
	"sail/conf"
	"sail/domain"
	"sail/errors"
	"sail/storage/domainstore"
	"sail/templates"
)

// BuildWithID returns domains that match the given id(s).
//
// It should be used to prepare one or more domains for rendering
// and is guaranteed to contain at least one correctly set up domain
// object at the first position of the returned slice.
func BuildWithID(ids ...uint32) []*domain.Domain {
	domains, err := fetchByID(ids...)
	if len(domains) == 0 || err != nil {
		errors.Log(err, conf.Instance().DevMode)
		return []*domain.Domain{domain.New()}
	}
	for _, d := range domains {
		d.Template = templates.BuildWithID(d.Template.ID)[0]
		cache.Domains[d.ID] = d
		fmt.Printf("domain added to cache: %d", d.ID)
	}
	return domains
}

func FromCache(id uint32) *domain.Domain {
	if domain, ok := cache.Domains[id].(*domain.Domain); ok {
		fmt.Printf("found domain in cache: %d\n", id)
		return domain
	}
	return BuildWithID(id)[0]
}

func fetchByID(ids ...uint32) ([]*domain.Domain, error) {
	return domainstore.Get().ByID(ids...).Domains()
}
