package task

import (
	"log"
	"time"

	"gorm.io/gorm"
	"pawprints-server/internal/model"
)

type ReminderChecker struct {
	DB *gorm.DB
}

func NewReminderChecker(db *gorm.DB) *ReminderChecker {
	return &ReminderChecker{DB: db}
}

func (c *ReminderChecker) Check() {
	now := time.Now()
	var count int64
	c.DB.Model(&model.Reminder{}).
		Where("completed = ? AND remind_at <= ? AND (snoozed_until IS NULL OR snoozed_until <= ?)", false, now, now).
		Count(&count)
	if count > 0 {
		log.Printf("[ReminderChecker] %d active reminders due", count)
	}
}

func (c *ReminderChecker) Start(interval time.Duration) {
	c.Check()
	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			c.Check()
		}
	}()
}
