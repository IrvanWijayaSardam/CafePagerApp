package repository

import (
	"github.com/IrvanWijayaSardam/mynotesapp/entity"
	"gorm.io/gorm"
)

type PagerRepository interface {
	InsertPager(b *entity.Pager) entity.Pager
	UpdatePager(b *entity.Pager) entity.Pager
	DeletePager(b *entity.Pager)
	AllPager() []entity.Pager
	FindPagerById(pagerUserID uint64) entity.Pager
}

type pagerConnection struct {
	connection *gorm.DB
}

// NewPagerRepository for create connection .asd
func NewPagerRepository(dbConn *gorm.DB) PagerRepository {
	return &pagerConnection{
		connection: dbConn,
	}
}

func (db *pagerConnection) InsertPager(b *entity.Pager) entity.Pager {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return *b
}

func (db *pagerConnection) UpdatePager(b *entity.Pager) entity.Pager {
	db.connection.Save(&b)
	db.connection.Preload("UserID").Find(&b)
	return *b
}

func (db *pagerConnection) DeletePager(b *entity.Pager) {
	db.connection.Delete(&b)
}

func (db *pagerConnection) FindPagerById(UserID uint64) entity.Pager {
	var pager entity.Pager
	db.connection.Preload("user_id").Find(&pager, UserID)
	return pager
}

func (db *pagerConnection) AllPager() []entity.Pager {
	var pager []entity.Pager
	db.connection.Preload("User").Find(&pager)
	return pager
}
