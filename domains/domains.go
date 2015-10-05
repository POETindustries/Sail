package domains

import (
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
	}
	return domains
}

func fetchByID(ids ...uint32) ([]*domain.Domain, error) {
	return domainstore.Get().ByID(ids...).Domains()
}
