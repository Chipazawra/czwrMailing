package receivers

type IDBctx interface {
	Create(usr, receiver string) (uint, error)
	Read(usr string) ([]string, error)
	Update(usr string, id uint, receiver string) error
	Delete(usr string, id uint) error
}

type Receivers struct {
	dbctx IDBctx
}

func New(dbctx IDBctx) *Receivers {
	return &Receivers{dbctx: dbctx}
}

func (r *Receivers) Create(usr, receiver string) (uint, error) {

	id, err := r.dbctx.Create(usr, receiver)
	if err != nil {
		return 0, err
	}
	return id, nil

}

func (r *Receivers) Read(usr string) ([]string, error) {

	lst, err := r.dbctx.Read(usr)
	if err != nil {
		return nil, err
	}
	return lst, nil

}

func (r *Receivers) Update(usr string, id uint, receiver string) error {

	err := r.dbctx.Update(usr, id, receiver)
	if err != nil {
		return err
	}
	return err

}

func (r *Receivers) Delete(usr string, id uint) error {

	err := r.dbctx.Delete(usr, id)
	if err != nil {
		return err
	}
	return nil
}
