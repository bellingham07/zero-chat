// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"

	"time"

	"gorm.io/gorm"
)

type (
	messageModel interface {
		Insert(ctx context.Context, tx *gorm.DB, data *Message) error

		FindOne(ctx context.Context, id int64) (*Message, error)
		Update(ctx context.Context, tx *gorm.DB, data *Message) error

		Delete(ctx context.Context, tx *gorm.DB, id int64) error
		Transaction(ctx context.Context, fn func(db *gorm.DB) error) error
	}

	defaultMessageModel struct {
		conn  *gorm.DB
		table string
	}

	Message struct {
		Id        int64          `gorm:"column:id"`
		SendId    int64          `gorm:"column:send_id"`
		ReceiveId int64          `gorm:"column:receive_id"`
		Msg       string         `gorm:"column:msg"`
		CreatedAt time.Time      `gorm:"column:created_at"`
		UpdatedAt sql.NullTime   `gorm:"column:updated_at"`
		DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
	}
)

func (Message) TableName() string {
	return "`message`"
}

func newMessageModel(conn *gorm.DB) *defaultMessageModel {
	return &defaultMessageModel{
		conn:  conn,
		table: "`message`",
	}
}

func (m *defaultMessageModel) Insert(ctx context.Context, tx *gorm.DB, data *Message) error {
	db := m.conn
	if tx != nil {
		db = tx
	}
	err := db.WithContext(ctx).Create(&data).Error
	return err
}

func (m *defaultMessageModel) FindOne(ctx context.Context, id int64) (*Message, error) {
	var resp Message
	err := m.conn.WithContext(ctx).Model(&Message{}).Where("`id` = ?", id).Take(&resp).Error
	switch err {
	case nil:
		return &resp, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultMessageModel) Update(ctx context.Context, tx *gorm.DB, data *Message) error {
	db := m.conn
	if tx != nil {
		db = tx
	}
	err := db.WithContext(ctx).Save(data).Error
	return err
}

func (m *defaultMessageModel) Delete(ctx context.Context, tx *gorm.DB, id int64) error {
	db := m.conn
	if tx != nil {
		db = tx
	}
	err := db.WithContext(ctx).Delete(&Message{}, id).Error

	return err
}

func (m *defaultMessageModel) Transaction(ctx context.Context, fn func(db *gorm.DB) error) error {
	return m.conn.WithContext(ctx).Transaction(fn)
}