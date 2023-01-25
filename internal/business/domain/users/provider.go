package users

import (
	"encoding/json"
	"io"
)

type Provider struct {
	dataSource  io.Reader
	loadedUsers []User
}

func NewProvider(source io.Reader) *Provider {
	return &Provider{dataSource: source}
}

func (p *Provider) FindAll() ([]User, error) {
	if p.loadedUsers == nil {
		p.loadedUsers = make([]User, 0)
		if err := json.NewDecoder(p.dataSource).Decode(&p.loadedUsers); err != nil {
			return []User{}, err
		}
	}

	return p.loadedUsers, nil
}
