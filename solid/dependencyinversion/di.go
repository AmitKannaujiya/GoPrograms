package dependencyinversion

import "fmt"

type DB interface {
	connect() error
	query(string) (string, error)
}

type Mysql struct {
	Host string
	Port int
}

func (m *Mysql) connect() error {
	fmt.Printf("mysql connecting on host : %s, port : %d \n", m.Host, m.Port)
	return nil
}

func (m *Mysql) query(query string) (string, error) {
	var result string
	if err := m.connect(); err != nil {
		return result, fmt.Errorf("error connecting to db %s", err.Error())
	}
	result = fmt.Sprintf("mysql query : %s, result in row", query)
	return result, nil
}

type Postgres struct {
	Host string
	Port int
}

func (p *Postgres) connect() error {
	fmt.Printf("postgres connecting on host : %s, port : %d \n", p.Host, p.Port)
	return nil
}

func (p *Postgres) query(query string) (string, error) {
	var result string
	if err := p.connect(); err != nil {
		return result, fmt.Errorf("error connecting to db %s", err.Error())
	}
	result = fmt.Sprintf("postgres query : %s, result in row", query)
	return result, nil
}

type User struct{
	name string
	db DB
}

func NewUser(name string, db DB) *User {
	return &User{
		name: name,
		db: db,
	}
}

func(u *User) GetUser(query string) {
	u.db.query(query)
}