package repository

import (
	"context"
	"database/sql"
	"go-web-native/domain"
	"time"

	"github.com/doug-martin/goqu/v9"
)

type CustomerRepository struct {
    db *goqu.Database
}

func NewCustomerRepository(db *sql.DB) domain.CustomerRepository {
	return &CustomerRepository{
		db: goqu.New("default", db),
	}
}

// FindByID retrieves a customer by ID from the database.
func (cr CustomerRepository) FindByID(ctx context.Context, id string) (result domain.Customer, err error) {
	dataset := cr.db.From("customers").Where(goqu.C("id").Eq(id)).Where(goqu.C("deleted_at").IsNull())

	found, err := dataset.ScanStructContext(ctx, &result) 
	if err != nil {
		return domain.Customer{}, err 
	}
	if !found {
		return domain.Customer{}, nil 
	}

	return result, nil
}

func (cr CustomerRepository) FindAll(ctx context.Context) (result []domain.Customer, err error) {
	dataset := cr.db.From("customers").Where(goqu.C("deleted_at").IsNull())
	err = dataset.ScanStructsContext(ctx, &result)
	return
}

func (cr CustomerRepository) Save(ctx context.Context, c *domain.Customer) error {
	executor := cr.db.Insert("customers").Rows(c).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (cr CustomerRepository) Update(ctx context.Context, c *domain.Customer) error {
	executor := cr.db.Update("customers").Set(c).Where(goqu.C("id").Eq(c.ID)).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (cr CustomerRepository) Delete(ctx context.Context, id string) error {
	executor := cr.db.Update("customers").
		Where(goqu.C("id").Eq(id)).
		Set(goqu.Record{"deleted_at": sql.NullTime{Valid: true, Time: time.Now()}}).
		Executor()
	_, err := executor.ExecContext(ctx)

	return err
}