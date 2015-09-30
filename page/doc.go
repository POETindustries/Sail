// Package page creates html markup from database data and template files.
//
// Trustworthiness of Data
//
// The contract for handling database content is that methods and functions
// that read from the database assume that the database content is trustworthy
// while functions that write to the database assume that user input is not
// to be trusted.
//
// This is only safe as long as an intruder does not manage to write harmful
// code directly into the database by bypassing the filters used by write
// functions. This needs to be addressed.
//
// On Separation Logic From Views
//
// It is considered good practice to separate data, logic and outside views
// as much as possible (hence, MVC patterns and such). This package is
// something of an exception to this otherwise reasonable rule. There are
// some string constant and functions that contain html markup directly
// embedded into the code.
//
// The reason is simple: If all else fails, we still want to be able to let
// the user know that there is a problem with the website and that they should
// consider coming back later. We cannot load templates if something with
// loading templates is wrong and we cannot load data from databases if
// loading from databases is broken, so we have to assume that in the worst
// case scenario nothing else works other than simplest code.
//
// This is why there is some hardcoded html in this package, acting as some
// kind of failsafe.
package page
