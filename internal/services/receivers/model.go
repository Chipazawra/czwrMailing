package receivers

type dbcontext interface {
	create()
	read()
	update()
	delete()
}

type Receivers struct {
	dbctx *dbcontext
}

func New() *Receivers {
	return &Receivers{}
}
