package response

import (
	"bytes"
	"net/http"
	"sail/page"
	"sail/page/cache"
	"sail/page/content"
	"sail/page/fallback"
	"sail/page/template"
	"sail/page/widget"
)

type Response struct {
	FallbackID  uint32
	FallbackURL string
	Message     string
	URL         string

	response http.ResponseWriter
	request  *http.Request
}

func New(wr http.ResponseWriter, req *http.Request) *Response {
	return &Response{
		FallbackID:  1,
		FallbackURL: "/home",
		URL:         req.URL.Path,
		response:    wr,
		request:     req}
}

func (r *Response) Serve() {
	buf := bytes.NewBufferString(fallback.NOTFOUND404)
	if markup := cache.DB().Markup(r.URL); markup != nil {
		buf = bytes.NewBuffer(markup)
	} else {
		cnt := content.ByURL(r.URL)
		if cnt == nil {
			if cnt = content.ByURL(r.FallbackURL); cnt == nil {
				cnt = content.ByID(r.FallbackID)
			}
		}
		tmpl := cache.DB().Template(cnt.TemplateID)
		if tmpl == nil {
			tmpl = template.ByID(cnt.TemplateID)
			for _, w := range widget.ByIDs(tmpl.WidgetIDs...) {
				tmpl.Widgets[w.RefName] = w
			}
		}
		presenter := page.New(cnt, tmpl, cnt.URL)
		presenter.Message = r.Message
		buf = presenter.Compile()
	}
	buf.WriteTo(r.response)
}
