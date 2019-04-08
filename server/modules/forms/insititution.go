package forms

import "github.com/nomango/bellex/server/models"

type InsititutionForm struct {
	Name string `json:"name"`
}

func (i *InsititutionForm) Assign(ins *models.Insititution) error {
	ins.Name = i.Name
	return nil
}
