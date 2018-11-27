package main

import (
	"./providers"
)

type SessionJar interface {
	ClearSession(key string)

	LoadSession(key string) (*providers.SessionState, error)

	SaveSession(key string, s *providers.SessionState) error
}

type InMemSessionJar struct {
	SessionJar
	sessions map[string]*providers.SessionState
}

func NewInMemSessionJar() *InMemSessionJar {
	jar := InMemSessionJar{}
	jar.sessions = make(map[string]*providers.SessionState)
	return &jar
}

func (jar *InMemSessionJar) ClearSession(key string) {
	delete(jar.sessions, key)
}

func (jar *InMemSessionJar) LoadSession(key string) (*providers.SessionState, error) {
	s := jar.sessions[key]
	return s, nil
}

func (jar *InMemSessionJar) SaveSession(key string, s *providers.SessionState) error {
	jar.sessions[key] = s
	return nil
}
