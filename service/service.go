package service

import (
	"errors"
	"shortedLink/shortener"
	"time"

	"github.com/google/uuid"
)

type Service interface {
	Create(link string, userId string) (ShortedLink, error)
	Update(id string, shortLink string) (ShortedLink, error)
	FindLongLink(shortLink string) (ShortedLink, error)
}

type service struct {
	repository Repository
}

// func to intialize new service
func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(link string, userId string) (ShortedLink, error) {
	newId := uuid.New()

	shortedLink := shortener.GenerateShortLink(link, userId)

	var newShortLink = ShortedLink{
		Id:        newId.String(),
		LongLink:  link,
		ShortLink: shortedLink,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserId:    userId,
	}

	createShortLink, err := s.repository.Add(newShortLink)

	if err != nil {
		return createShortLink, err
	}

	return createShortLink, nil
}

func (s *service) Update(id string, shortLink string) (ShortedLink, error) {
	checkIdShortLink, _ := s.repository.FindById(id)

	if checkIdShortLink.Id == "" || checkIdShortLink.LongLink == "" {
		return checkIdShortLink, errors.New("error id not found")
	}

	checkShortLink, _ := s.repository.FindLink(shortLink)

	if checkShortLink.Id != "" || checkIdShortLink.ShortLink != "" {
		return checkIdShortLink, errors.New("short link created by other user")
	}

	var dataUpdate = map[string]interface{}{}

	dataUpdate["short_link"] = shortLink
	dataUpdate["updated_at"] = time.Now()

	updateShortLink, err := s.repository.UpdateShortLink(id, dataUpdate)

	if err != nil {
		return updateShortLink, err
	}

	return updateShortLink, nil
}

func (s *service) FindLongLink(shortLink string) (ShortedLink, error) {
	findShortLink, err := s.repository.FindLink(shortLink)

	if err != nil {
		return findShortLink, err
	}

	if findShortLink.Id == "" || findShortLink.LongLink == "" {
		return findShortLink, errors.New("error data not found")
	}

	return findShortLink, nil
}
