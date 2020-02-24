// Code generated by qtc from "find.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line find.html:1
package templates

//line find.html:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line find.html:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line find.html:2
type FindPage struct {
    Results []Result
  }

//line find.html:5
func (p *FindPage) StreamName(qw422016 *qt422016.Writer) {
//line find.html:5
	qw422016.N().S(`find`)
//line find.html:5
}

//line find.html:5
func (p *FindPage) WriteName(qq422016 qtio422016.Writer) {
//line find.html:5
	qw422016 := qt422016.AcquireWriter(qq422016)
//line find.html:5
	p.StreamName(qw422016)
//line find.html:5
	qt422016.ReleaseWriter(qw422016)
//line find.html:5
}

//line find.html:5
func (p *FindPage) Name() string {
//line find.html:5
	qb422016 := qt422016.AcquireByteBuffer()
//line find.html:5
	p.WriteName(qb422016)
//line find.html:5
	qs422016 := string(qb422016.B)
//line find.html:5
	qt422016.ReleaseByteBuffer(qb422016)
//line find.html:5
	return qs422016
//line find.html:5
}

//line find.html:5
func (p *FindPage) StreamTitle(qw422016 *qt422016.Writer) {
//line find.html:5
	qw422016.N().S(`Search IMDb`)
//line find.html:5
}

//line find.html:5
func (p *FindPage) WriteTitle(qq422016 qtio422016.Writer) {
//line find.html:5
	qw422016 := qt422016.AcquireWriter(qq422016)
//line find.html:5
	p.StreamTitle(qw422016)
//line find.html:5
	qt422016.ReleaseWriter(qw422016)
//line find.html:5
}

//line find.html:5
func (p *FindPage) Title() string {
//line find.html:5
	qb422016 := qt422016.AcquireByteBuffer()
//line find.html:5
	p.WriteTitle(qb422016)
//line find.html:5
	qs422016 := string(qb422016.B)
//line find.html:5
	qt422016.ReleaseByteBuffer(qb422016)
//line find.html:5
	return qs422016
//line find.html:5
}

//line find.html:5
func (p *FindPage) StreamContent(qw422016 *qt422016.Writer) {
//line find.html:5
	qw422016.N().S(`<div class=jumbotron><h1 class=display-3>IMBb Results (`)
//line find.html:5
	qw422016.N().D(len(p.Results))
//line find.html:5
	qw422016.N().S(`)</h1><p></p><form action=/find><label for=q>Find</label> <input name=q placeholder=Something...> <button type=submit>Search!</button></form><p></p></div><div class="">`)
//line find.html:5
	for _, r := range p.Results {
//line find.html:5
		qw422016.N().S(`<div><p>`)
//line find.html:5
		qw422016.N().S(r.PageID)
//line find.html:5
		qw422016.N().S(`<a href="`)
//line find.html:5
		qw422016.N().S(r.Location)
//line find.html:5
		qw422016.N().S(`">`)
//line find.html:5
		qw422016.N().S(r.Title)
//line find.html:5
		qw422016.N().S(`</a>(`)
//line find.html:5
		qw422016.N().F(r.Rank)
//line find.html:5
		qw422016.N().S(`)</p><span>`)
//line find.html:5
		qw422016.N().S(r.Summary)
//line find.html:5
		qw422016.N().S(`</span></div>`)
//line find.html:5
	}
//line find.html:5
	qw422016.N().S(`</div>`)
//line find.html:5
}

//line find.html:5
func (p *FindPage) WriteContent(qq422016 qtio422016.Writer) {
//line find.html:5
	qw422016 := qt422016.AcquireWriter(qq422016)
//line find.html:5
	p.StreamContent(qw422016)
//line find.html:5
	qt422016.ReleaseWriter(qw422016)
//line find.html:5
}

//line find.html:5
func (p *FindPage) Content() string {
//line find.html:5
	qb422016 := qt422016.AcquireByteBuffer()
//line find.html:5
	p.WriteContent(qb422016)
//line find.html:5
	qs422016 := string(qb422016.B)
//line find.html:5
	qt422016.ReleaseByteBuffer(qb422016)
//line find.html:5
	return qs422016
//line find.html:5
}
