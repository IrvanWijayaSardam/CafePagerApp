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
	db.connection.Preload("PagerID").Find(&b)
	return *b
}

func (db *pagerConnection) UpdatePager(b *entity.Pager) entity.Pager {
	db.connection.Save(&b)
	db.connection.Preload("PagerID").Find(&b)
	return *b
}

func (db *pagerConnection) DeletePager(b *entity.Pager) {
	db.connection.Delete(&b)
}

func (db *pagerConnection) FindPagerById(PagerID uint64) entity.Pager {
	var pager entity.Pager
	db.connection.Preload("PagerID").Find(&pager, PagerID)
	return pager
}

func (db *pagerConnection) AllPager() []entity.Pager {
	var pager []entity.Pager
	db.connection.Preload("PagerID").Find(&pager)
	return pager
}
