package service

import (
	"fmt"
	"log"

	"github.com/IrvanWijayaSardam/mynotesapp/dto"
	"github.com/IrvanWijayaSardam/mynotesapp/entity"
	"github.com/IrvanWijayaSardam/mynotesapp/repository"
	"github.com/mashingan/smapping"
)

type PagerService interface {
	Insert(b dto.PagerCreateDTO) entity.Pager
	Update(b dto.PagerUpdateDTO) entity.Pager
	Delete(b entity.Pager)
	All() []entity.Pager
	FindByID(bookID uint64) entity.Pager
	IsAllowedToEdit(userID string, bookID uint64) bool
}

type pagerService struct {
	pagerRepository repository.PagerRepository
}

func NewPagerService(pagerRepo repository.PagerRepository) PagerService {
	return &pagerService{
		pagerRepository: pagerRepo,
	}
}

func (service *pagerService) Insert(b dto.PagerCreateDTO) entity.Pager {
	pager := entity.Pager{}
	err := smapping.FillStruct(&pager, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.pagerRepository.InsertPager(&pager)
	return res
}

func (service *pagerService) Update(b dto.PagerUpdateDTO) entity.Pager {
	pager := entity.Pager{}
	err := smapping.FillStruct(&pager, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v : ", err)
	}
	res := service.pagerRepository.UpdatePager(&pager)
	return res
}

func (service *pagerService) Delete(b entity.Pager) {
	service.pagerRepository.DeletePager(&b)
}

func (service *pagerService) All() []entity.Pager {
	return service.pagerRepository.AllPager()
}

func (service *pagerService) FindByID(pagerID uint64) entity.Pager {
	return service.pagerRepository.FindPagerById(pagerID)
}

func (service *pagerService) IsAllowedToEdit(userID string, pagerID uint64) bool {
	b := service.pagerRepository.FindPagerById(pagerID)
	id := fmt.Sprintf("%v", b.User.ID)
	return userID == id
}
