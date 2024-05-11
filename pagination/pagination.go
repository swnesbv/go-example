package pagination

import (
	"errors"
)

type Args struct {
	Max     int
	Pos     int
	Page    int
	Records int
	Total   int
	Size    int
}

const (
	ErrPageNo     = errBase + "Page > Pages"
	errBase       = "Error in pagination: "
	errInvSize    = errBase + "Size <= 0"
	errResultSize = errBase + "Results > Size"
	errRecordSize = errBase + "Results > Records"
)

func (a *Args) pages() (p int, err error) {
	p = (a.Total-1)/a.Size + 1
	if a.Page > p {
		err = errors.New(ErrPageNo)
	}
	return
}

func (a *Args) check() (err error) {
	if a.Size <= 0 {
		err = errors.New(errInvSize)
		return
	}
	if a.Records > a.Size {
		err = errors.New(errResultSize)
		return
	}
	if a.Records > a.Total {
		err = errors.New(errRecordSize)
	}
	return
}

type Pagination struct {
	pages int
	args  *Args
}

func NewPn(a Args) (pgn *Pagination, err error) {
	if err = a.check(); err != nil {
		return
	}
	p, err := a.pages()
	if err != nil {
		return
	}
	if p < a.Max {
		a.Max = p
	}

	pgn = &Pagination{
		pages: p,
		args:  &a,
	}
	return
}

func (p *Pagination) Prev() int {
	if p.args.Page <= 1 {
		return 0
	}
	return p.args.Page - 1
}

func (p *Pagination) Page() int {
	return p.args.Page
}

func (p *Pagination) Next() int {
	if p.args.Page == p.pages {
		return 0
	}
	return p.args.Page + 1
}

func (p *Pagination) Records() int {
	return p.args.Records
}

func (p *Pagination) Total() int {
	return p.args.Total
}

func (p *Pagination) Size() int {
	return p.args.Size
}

func (p *Pagination) Pages() int {
	return p.pages
}

type Entry struct {
	Active bool
	Number int
}

func (p *Pagination) Entries() (r []Entry) {
	sn := p.args.Page - p.args.Pos
	switch {
	case sn < 0:
		sn = 0
		break
	case p.args.Max-p.args.Pos+p.args.Page > p.pages:
		sn = p.pages - p.args.Max
	}
	sn++

	r = make([]Entry, p.args.Max)
	for i := 0; i < p.args.Max; i++ {
		var e Entry
		e.Number = sn + i
		e.Active = e.Number == p.args.Page
		r[i] = e
	}
	return
}
