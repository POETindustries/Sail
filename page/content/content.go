package content

// NOTFOUND404 is a very basic web page signaling a 404 error.
// It contails the bare minimum necessary for a syntactically correct html web
// page and is used in those cases when not even basic database connections
// and templates work. The cms cannot be considered functional should that
// happen, and this markup at least tells the user as much. The markup is as
// generic as possible while still being somewhat good looking.
const NOTFOUND404 = `
<!doctype html>
<html>
	<head><title>Sorry About That</title><meta charset="utf-8"></head>
	<body style="background:black;text-align:center;color:white;padding:72px;font-size:1.5em;">
		<p style="font-size:2em;">Sorry About That!</p>
		<p>PAGE NOT FOUND</p>
	</body>
</html>`

// Page contains the information needed to generate a web page for display.
// This is the basic struct that contains all information needed to generate
// a correct and complete html page. It is the responsibility of the other
// functions and methods in package page to make sure its fields are
// properly initialized.
type Page struct {
	ID       uint32
	Title    string
	URL      string
	Content  string
	Meta     *Meta
	Template *Template

	Status int8
	Owner  string
	CDate  string
	EDate  string
}

// NewPage creates a new Page object with usable defaults.
func NewPage() *Page {
	return &Page{
		Meta:     NewMeta(),
		Template: NewTemplate()}
}

func FetchPageByURL(urls ...string) ([]*data.Page, error) {
	return Get().ByURL(urls...).Pages()
}

func FetchPageByID(ids ...uint32) ([]*data.Page, error) {
	return Get().ByID(ids...).Pages()
}

func Load404() *data.Page {
	p := NewPage()
	p.ID, p.Title = 0, "Sorry about that"
	return p
}
