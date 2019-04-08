package forms

import "github.com/nomango/bellex/server/models"

type ScheduleForm struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

func (m *ScheduleForm) Assign(schedule *models.Schedule) error {
	schedule.Name = m.Name
	schedule.Content = m.Content
	return nil
}
