package forms

import "github.com/nomango/bellex/server/models"

type InstitutionForm struct {
	Name string `json:"name"`
}

func (i *InstitutionForm) Assign(ins *models.Institution) error {
	ins.Name = i.Name
	return nil
}
