package response

import (
	"net/http"
	"sail/object/cache"
	"sail/object/content"
	"sail/object/template"
	"sail/object/widget"
)

type Response struct {
	FallbackID  uint32
	FallbackURL string
	Message     string
	URL         string
	Presenter   Presenter

	response http.ResponseWriter
	request  *http.Request
	content  *content.Content
	template *template.Template
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
	r.Presenter.SetURL(r.URL)
	r.Presenter.SetMessage(r.Message)
	m := r.Presenter.Compile()
	//cache.DB().PushMarkup(r.URL, m.Bytes())
	m.WriteTo(r.response)
}

func (r *Response) Content() *content.Content {
	if r.content == nil {
		cnt := cache.DB().Content(r.URL)
		if cnt == nil {
			if cnt = content.ByURL(r.URL); cnt == nil {
				if cnt = content.ByURL(r.FallbackURL); cnt == nil {
					cnt = content.ByID(r.FallbackID)
				}
			}
		}
		//cache.DB().PushContent(cnt)
		r.content = cnt
	}
	return r.content
}

func (r *Response) Template() *template.Template {
	if r.template == nil {
		tmpl := cache.DB().Template(r.Content().TemplateID)
		if tmpl == nil {
			tmpl = template.ByID(r.Content().TemplateID)
			for _, w := range widget.ByIDs(tmpl.WidgetIDs...) {
				tmpl.Widgets[w.RefName] = w
			}
		}
		//cache.DB().PushTemplate(tmpl)
		r.template = tmpl
	}
	return r.template
}
