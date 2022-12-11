package repository

import (
	"fmt"
	"math"
	"time"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) List(offset, limit int, from, to *time.Time, search string, by, fields []string, table, queryJoin string) (tx *gorm.DB, total, pages uint64) {
	tx = r.DB
	if queryJoin != "" {
		tx = tx.Joins(queryJoin)
	}

	if from != nil && !from.IsZero() {
		tx = tx.Where(table+".created_at >= ?", from)
	}

	if to != nil && !to.IsZero() {
		tx = tx.Where(table+".created_at <= ?", to)
	}

	m := make(map[string]bool)
	for _, field := range fields {
		m[field] = true
	}
	if search != "" {
		search = "%" + search + "%"
		conditions := r.DB
		if len(by) > 0 && by[0] != "" {
			for i, field := range by {
				if m[field] {
					if i > 0 {
						conditions = conditions.Or(fmt.Sprintf("%s LIKE ?", field), search)
					} else {
						conditions = conditions.Where(fmt.Sprintf("%s LIKE ?", field), search)
					}
				}
			}
		} else {
			for i, field := range fields {
				if i > 0 {
					conditions = conditions.Or(fmt.Sprintf("%s LIKE ?", field), search)
				} else {
					conditions = conditions.Where(fmt.Sprintf("%s LIKE ?", field), search)
				}
			}
		}
		tx = tx.Where(conditions)
	}

	total, pages = totalPages(tx, limit, table)

	return
}

func totalPages(tx *gorm.DB, limit int, table string) (total uint64, pages uint64) {
	var count int64
	tx.Debug().Table(table).Count(&count)

	total = uint64(count)
	pages = 1
	if count > 0 && limit > 0 {
		pages = uint64(math.Ceil(float64(count) / float64(limit)))
	}

	return
}

func New(db *gorm.DB) *Repository {
	return &Repository{db}
}
