package repository

import (
	"context"
	"database/sql"

	"go-web-native/domain"

	"github.com/doug-martin/goqu/v9"
)

type BookStockRepository struct {
	db *goqu.Database
}

func NewStock(db *sql.DB) domain.BookStockRepository {
	return &BookStockRepository{
		db: goqu.New("default", db),
	}
}

func (b BookStockRepository) FindBookId(ctx context.Context, id string) (result []domain.BookStock, err error) {
	dataset := b.db.From("book_stocks").Where(
		goqu.C("book_id").Eq(id),
	)
	err = dataset.ScanStructsContext(ctx, &result)
	return
}

func (b BookStockRepository) FindByBookAndCode(ctx context.Context, id string, code string) error {
	dataset := b.db.From("book_stocks").Where(
		goqu.C("book_id").Eq(id),
		goqu.C("code").Eq(code),
		goqu.C("deleted_at").IsNull(),
	)
	var result domain.BookStock
	found, err := dataset.ScanStructContext(ctx, &result) 
	if err != nil {
		return err
	}
	if !found {
		return sql.ErrNoRows
	}
	return nil
}

func (b BookStockRepository) Save(ctx context.Context, data []domain.BookStock) error {
	executor := b.db.Insert("book_stocks").Rows(data).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (b BookStockRepository) Update(ctx context.Context, stock *domain.BookStock) error {
	executor := b.db.Update("book_stocks").
		Where(goqu.C("code").Eq(stock.Code)).
		Set(stock).
		Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (b BookStockRepository) DeleteByBookId(ctx context.Context, id string) error {
	executor := b.db.Update("book_stocks").
		Where(goqu.C("book_id").Eq(id)).
		Executor()

	_, err := executor.ExecContext(ctx)
	return err
}

func (b BookStockRepository) DeleteByCodes(ctx context.Context, codes []string) error {
	_, err := b.db.Delete("book_stocks").
		Where(goqu.C("code").In(codes)).
		Executor().
		ExecContext(ctx)
	return err
}
