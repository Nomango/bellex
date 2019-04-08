package forms

import "github.com/nomango/bellex/server/models"

type MechineForm struct {
	Code       string `json:"code"`
	Secret     string `json:"secret"`
	ScheduleID int    `json:"schedule_id"`
}

func (m *MechineForm) Assign(mechine *models.Mechine) error {
	schedule := &models.Schedule{Id: m.ScheduleID}
	if err := schedule.Read(); err != nil {
		return err
	}

	mechine.Code = m.Code
	mechine.Secret = m.Secret
	mechine.Schedule = schedule
	return nil
}
