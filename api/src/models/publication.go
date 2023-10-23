package models

import (
	"api/db/database"
	"errors"
	"strings"
)

type Publication database.Publication

func (p *Publication) Prepare() error {

	if err := p.validatePublication(); err != nil {
		return err
	}

	p.format()

	return nil
}

func (p *Publication) format() {
	p.Title = strings.TrimSpace(p.Title)
	p.Content = strings.TrimSpace(p.Content)
}

func (p *Publication) validatePublication() error {
	if p.Title == "" {
		return errors.New("title cant be null")
	}

	if p.Content == "" {
		return errors.New("content cant be null")
	}

	return nil
}
