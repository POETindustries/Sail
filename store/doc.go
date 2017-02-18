/*
Package store handles all access to permanent storage.

Trustworthiness of Data

The contract for handling database content is that methods and functions
that read from the database assume that the database content is trustworthy
while functions that write to the database assume that user input is not
to be trusted.

This is only safe as long as an intruder does not manage to write harmful
code directly into the database by bypassing the filters used by write
functions. This needs to be addressed.
*/
package store
