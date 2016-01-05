package cache

// Pages holds all cached pages, indexed by url
var Pages = make(map[string]interface{})

// Domains holds the cached domains, indexed by id
var Domains = make(map[uint32]interface{})

// Markup stores the markup that was last compiled for the given url
var Markup = make(map[string]interface{})
