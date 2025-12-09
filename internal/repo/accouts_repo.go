package repo

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type MgmtDB struct{ Pool *pgxpool.Pool }

func NewMgmtDB(url string) (*MgmtDB, error) {
	p, err := pgxpool.New(context.Background(), url)
	if err != nil {
		return nil, err
	}
	return &MgmtDB{Pool: p}, nil
}
func (d *MgmtDB) Close() { d.Pool.Close() }

type Account struct {
	ID        string
	Email     string
	FullName  string
	BirthDate string
	Country   string
	Status    string
}

func (d *MgmtDB) CreateAccount(ctx context.Context, a Account) (*Account, error) {
	row := d.Pool.QueryRow(ctx, `
        INSERT INTO accounts (email, full_name, birth_date, country)
        VALUES ($1,$2,$3,$4)
        RETURNING id, email, full_name, birth_date, country, status
    `, a.Email, a.FullName, a.BirthDate, a.Country)
	var out Account
	if err := row.Scan(&out.ID, &out.Email, &out.FullName, &out.BirthDate, &out.Country, &out.Status); err != nil {
		return nil, err
	}
	return &out, nil
}

func (d *MgmtDB) GetAccount(ctx context.Context, id string) (*Account, error) {
	row := d.Pool.QueryRow(ctx, `
        SELECT id, email, full_name, birth_date, country, status
        FROM accounts WHERE id=$1
    `, id)
	var a Account
	if err := row.Scan(&a.ID, &a.Email, &a.FullName, &a.BirthDate, &a.Country, &a.Status); err != nil {
		return nil, err
	}
	return &a, nil
}

func (d *MgmtDB) ListAccounts(ctx context.Context) ([]Account, error) {
	rows, err := d.Pool.Query(ctx, `
        SELECT id, email, full_name, birth_date, country, status
        FROM accounts ORDER BY created_at DESC
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []Account
	for rows.Next() {
		var a Account
		if err := rows.Scan(&a.ID, &a.Email, &a.FullName, &a.BirthDate, &a.Country, &a.Status); err != nil {
			return nil, err
		}
		res = append(res, a)
	}
	return res, nil
}
