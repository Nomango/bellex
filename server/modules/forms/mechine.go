package forms

import "github.com/nomango/bellex/server/models"

type MechineForm struct {
	Name       string `json:"name"`
	Code       string `json:"code"`
	ScheduleID int    `json:"schedule_id"`
}

func (m *MechineForm) Assign(mechine *models.Mechine) error {
	schedule := &models.Schedule{Id: m.ScheduleID}
	if err := schedule.Read(); err != nil {
		return err
	}

	mechine.Name = m.Name
	mechine.Code = m.Code
	mechine.SetNewSchedule(schedule)
	return nil
}
