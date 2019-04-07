package forms

import "github.com/nomango/bellex/server/models"

type ScheduleForm struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

func (m *ScheduleForm) Update(schedule *models.Schedule) {
	schedule.Name = m.Name
	schedule.Content = m.Content
}
