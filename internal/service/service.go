package service

import (
	"fmt"
	"log"
	"net/url"

	"github.com/takclark/schedulator/api"
	"github.com/takclark/schedulator/internal/engine"
	"github.com/takclark/schedulator/internal/jobs/requester"
)

type Service struct {
	e    *engine.Engine
	repo Repository
	l    *log.Logger
}

type Repository interface {
	Rules() ([]api.Rule, error)
	Rule(int64) (api.Rule, error)
	CreateRule(api.CreateRule) (api.Rule, error)
	UpdateRule(api.UpdateRule) (api.Rule, error)
}

func NewService(e *engine.Engine, repo Repository, l *log.Logger) *Service {
	return &Service{
		e:    e,
		repo: repo,
		l:    l,
	}
}

func (s *Service) Rules() ([]api.Rule, error) {
	rs, err := s.repo.Rules()
	if err != nil {
		s.l.Println("error loading rules:", err)
		return []api.Rule{}, err
	}

	return rs, nil
}

func (s *Service) Rule(id int64) (api.Rule, error) {
	return s.repo.Rule(id)
}

func (s *Service) CreateRule(data api.CreateRule) (api.Rule, error) {
	if data.Type != api.RuleTypeRequester {
		s.l.Println("disallowed type:", data.Type)
		return api.Rule{}, fmt.Errorf("unknown type: %v", data.Type)
	}

	converted, err := convertRequester(data)
	if err != nil {
		s.l.Println("error creating rule", err)
		return api.Rule{}, fmt.Errorf("converting requester: %w", err)
	}

	_, err = s.e.Register(converted)
	if err != nil {
		s.l.Println("error creating rule", err)
		return api.Rule{}, fmt.Errorf("registering requester: %w", err)
	}

	r, err := s.repo.CreateRule(data)
	if err != nil {
		s.l.Println("error creating rule", err)
		return api.Rule{}, err
	}

	return r, nil
}

func (s *Service) UpdateRule(data api.UpdateRule) (api.Rule, error) {
	return s.repo.UpdateRule(data)
}

func convertRequester(data api.CreateRule) (*requester.Requester, error) {
	u, err := url.Parse(data.URL)
	if err != nil {
		return &requester.Requester{}, nil
	}

	r := &requester.Requester{
		Exp:    data.Expression,
		URL:    *u,
		Method: data.Method,
		Body:   data.Body,
	}

	return r, nil
}
