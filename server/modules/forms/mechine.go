package forms

import "github.com/nomango/bellex/server/models"

type MechineForm struct {
	Code       string `json:"code"`
	Secret     string `json:"secret"`
	ScheduleID int    `json:"schedule_id"`
}

func (m *MechineForm) Update(mechine *models.Mechine) {
	mechine.Code = m.Code
	mechine.Secret = m.Secret
	mechine.Schedule = &models.Schedule{Id: m.ScheduleID}
}
