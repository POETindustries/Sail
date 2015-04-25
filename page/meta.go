package page

import ()

// DBKeys holds an enumeration of database table column names.
// This is for insertion into select queries that would become very long
// if they were written out ervery time. Another perk of defining them
// here is that there is only one place within the code base where these
// values had to be changed if the database layout itself changed.
const DBMETAKEYS = "title,keywords,description,language,page_topic,revisit_after,robots"

type Meta struct {
	Title        string
	Keywords     string
	Description  string
	Language     string
	PageTopic    string
	RevisitAfter string
	Robots       string
}
