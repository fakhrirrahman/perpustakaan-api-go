package repository

import (
	"context"
	"database/sql"
	"time"

	"go-web-native/domain"

	"github.com/doug-martin/goqu/v9"
)

type bookRepository struct {
	db *goqu.Database
}

func NewBook(db *sql.DB) domain.BookRepository {
	return &bookRepository{
		db: goqu.New("default", db),
	}
}

func (b bookRepository) FindAll(ctx context.Context) (books []domain.Book, err error) {
	dataset := b.db.From("books").Where(goqu.C("deleted_at").IsNull())
	err = dataset.ScanStructsContext(ctx, &books) 
	return
}

func (b bookRepository) FindById(ctx context.Context, id string) (book domain.Book, err error) {
	dataset := b.db.From("books").Where(
		goqu.C("id").Eq(id),
		goqu.C("deleted_at").IsNull(),
	)
	_, err = dataset.ScanStructContext(ctx, &book)
	return
}

func (b bookRepository) Save(ctx context.Context, book *domain.Book) error {
	executor := b.db.Insert("books").Rows(book).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (b bookRepository) Update(ctx context.Context, book *domain.Book) error {
	executor := b.db.Update("books").Set(book).Where(goqu.C("id").Eq(book.ID)).Executor() 
	_, err := executor.ExecContext(ctx)
	return err
}

func (b bookRepository) Delete(ctx context.Context, id string) error {
	executor := b.db.Update("books").
		Where(goqu.C("id").Eq(id)).
		Set(goqu.Record{"deleted_at": time.Now()}).Executor() 
	_, err := executor.ExecContext(ctx)
	return err
}
