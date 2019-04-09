package forms

import "github.com/nomango/bellex/server/models"

type MechineForm struct {
	Name       string `json:"name"`
	Code       string `json:"code"`
	ScheduleID int    `json:"schedule_id"`
}

func (m *MechineForm) Assign(mechine *models.Mechine) error {
	mechine.Schedule = &models.Schedule{Id: m.ScheduleID}
	if err := mechine.Schedule.Read(); err != nil {
		return err
	}

	mechine.Name = m.Name
	mechine.Code = m.Code
	return nil
}
