package forms

import (
	"errors"
	"regexp"

	"github.com/nomango/bellex/server/models"
)

type ScheduleForm struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

func (m *ScheduleForm) Assign(schedule *models.Schedule) error {
	matched, _ := regexp.MatchString(`^\d{2}:\d{2}( \d{2}:\d{2})*$`, m.Content)
	if !matched {
		return errors.New("数据格式有误")
	}

	schedule.Name = m.Name
	schedule.Content = m.Content
	return nil
}
