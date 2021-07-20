package service

import (
	"errors"
	"time"
)

type Service interface {
	GetAll() ([]ShortedLink, error)
	Create(link string) (ShortedLink, error)
	Update(shortLink string) (ShortedLink, error)
	FindLongLink(shortLink string) (ShortedLink, error)
}

type service struct {
	Memmory []ShortedLink
}

func NewService(m []ShortedLink) *service {
	return &service{m}
}

func (s *service) GetAll() ([]ShortedLink, error) {
	return s.Memmory, nil
}

func (s *service) Create(link string) (ShortedLink, error) {
	if link == "" {
		return ShortedLink{}, errors.New("error link not input")
	}

	// createShortLink := s.RandString(8)

	// for _, m := range s.Memmory {
	// if m.ShortLink == createShortLink {
	// 	createShortLink = s.RandString(8)
	// }
	// }

	newShortLinkData := ShortedLink{
		// Id:        Count,
		LongLink: link,
		// ShortLink: createShortLink,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	s.Memmory = append(s.Memmory, newShortLinkData)
	return newShortLinkData, nil
}

func (s *service) Update(shortLink string) (ShortedLink, error) {
	return ShortedLink{}, nil
}

func (s *service) FindLongLink(shortLink string) (ShortedLink, error) {
	var find ShortedLink

	for _, m := range s.Memmory {
		if m.ShortLink == shortLink {
			find = m
		}
	}

	if find.LongLink == "" && find.ShortLink == "" {
		return find, errors.New("error shorted link not found")
	}

	return find, nil
}
