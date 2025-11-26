package dao

import (
	"context"
	"github.com/itering/subscan/model"
)

func (d *Dao) CreateNewSession(ctx context.Context, sessionId uint, validators []string) error {
	return d.db.WithContext(ctx).Scopes(model.IgnoreDuplicate).Create(&model.Session{
		SessionId:  sessionId,
		Validators: validators,
	}).Error
}

func (d *Dao) GetSessionValidatorsById(ctx context.Context, sessionId uint) []string {
	var session model.Session
	if err := d.db.WithContext(ctx).Where("session_id = ?", sessionId).First(&session).Error; err != nil {
		return nil
	}
	return session.Validators
}
