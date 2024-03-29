package activities

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"

	"github.com/LightAlykard/GoBackEnd-2/Hw5-1-0.1.0/HW5-1/manager"
	"github.com/LightAlykard/GoBackEnd-2/Hw5-1-0.1.0/HW5-1/pool"
)

type Activity struct {
	UserId int
	Date   time.Time
	Name   string
}

func (a *Activity) connection(m *manager.Manager, p *pool.Pool) (*sql.DB, error) {
	s, err := m.ShardById(a.UserId)
	if err != nil {
		return nil, err
	}
	return p.Connection(s.Address)
}

func (a *Activity) Create(m *manager.Manager, p *pool.Pool) error {
	c, err := a.connection(m, p)
	if err != nil {
		return err
	}
	_, err = c.Exec(`INSERT INTO "activities" VALUES ($1, $2, $3)`, a.UserId,
		a.Date, a.Name)
	return err
}
func (a *Activity) Read(m *manager.Manager, p *pool.Pool) error {
	c, err := a.connection(m, p)
	if err != nil {
		return err
	}
	r := c.QueryRow(`SELECT "date", "name" FROM "activities" WHERE "user_id" =
	$1`, a.UserId)
	return r.Scan(
		&a.Date,
		&a.Name,
	)
}
func (a *Activity) Update(m *manager.Manager, p *pool.Pool) error {
	c, err := a.connection(m, p)
	if err != nil {
		return err
	}
	_, err = c.Exec(`UPDATE "activities" SET "name" = $3, "date" = $4
	WHERE "user_id" = $1 AND "date" = $2`, a.UserId,
		a.Name, a.Date, a.Name)
	return err
}
func (a *Activity) Delete(m *manager.Manager, p *pool.Pool) error {
	c, err := a.connection(m, p)
	if err != nil {
		return err
	}
	_, err = c.Exec(`DELETE FROM "activities" WHERE "user_id" = $1 AND "date" = $2`,
		a.UserId, a.Date)
	return err
}
