package repository

import (
	"database/sql"
	"go-web-native/domain"

	"github.com/doug-martin/goqu/v9"
)

type CustomerRepository struct {
    db *goqu.Database
}

func NewCustomerRepository(db *sql.DB) *CustomerRepository {
	return &CustomerRepository{
		db: goqu.New("mysql", db),
	}
}


func (cr *CustomerRepository) FindByID(string) (*domain.Customer, error) {
panic("implement me")
}

func (cr *CustomerRepository) FindAll() ([]*domain.Customer, error) {
panic("implement me")
}

func (cr *CustomerRepository) Save(customer *domain.Customer) error {
panic("implement me")
}

func (cr *CustomerRepository) Update(customer *domain.Customer) error {
panic("implement me")
}

func (cr *CustomerRepository) Delete(id int) error {
panic("implement me")
}



