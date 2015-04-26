package page

import ()

// DBKeys holds an enumeration of database table column names.
// This is for insertion into select queries that would become very long
// if they were written out ervery time. Another perk of defining them
// here is that there is only one place within the code base where these
// values had to be changed if the database layout itself changed.
const DBMETAKEYS = "title,keywords,description,language,page_topic,revisit_after,robots"

// Meta holds the meta information of a web page. It is used to store values
// for display in an html page's <head> block. This struct holds values that
// are used foremost for SEO purposes. Some meta information that is not page
// specific and doesn't really change across websites is omitted here in favor
// of being embedded directly into the templates. (The charset directive is a
// good example. It is and should be set to utf-8, always, so there's no reason
// to store and process it on a per-page basis.)
type Meta struct {
	Title        string
	Keywords     string
	Description  string
	Language     string
	PageTopic    string
	RevisitAfter string
	Robots       string
}
