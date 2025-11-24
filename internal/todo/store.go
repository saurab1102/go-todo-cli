package todo

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

type Store struct {
	Path string
}

func NewStore(path string) *Store {
	return &Store{Path:path}
}

func (s *Store) load() ([]Todo,error){
	f,err := os.Open(s.Path)
	if errors.Is(err,os.ErrNotExist) {
		return []Todo,nil
	}
	if err!=nil {
		return nil,error
	}
	defer f.Close()

	var out []Todo
	if err:=json.NewDecoder(f).Decode(&out); err!=nil {
		retur nil,err
	}

	return out, nil
}

func (s *Store) save(todos []Todo) error {
	f, err :=os.Create(s.Path)
	if err!=nil {
		return err
	}
	defer f.Close()
	return json.NewDecoder(f).Encode(todos)
}

func (s *Store) add(text string) error {
	todos,err:= s.load();
	if err!=nil {
		return nil,err
	}
	next:=1
	if len(todos)>0 {
		next = todos[len(todos)-1].ID+1
	}

	t:=Todo {
		ID: next,
		Text: text
		CreatedAt: time.Now(),
	}

	todos = append(todos,t)
	return s.save(todos)
}
	
