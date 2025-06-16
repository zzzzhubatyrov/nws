package repository

import (
    "gorm.io/gorm"
)

// Generic repository abstraction to simplify future changes

type Repository[T any] struct {
    db *gorm.DB
}

func New[T any](db *gorm.DB) *Repository[T] {
    return &Repository[T]{db: db}
}

func (r *Repository[T]) All(out *[]T) error {
    return r.db.Find(out).Error
}

func (r *Repository[T]) GetByID(id uint, out *T) error {
    return r.db.First(out, id).Error
}

func (r *Repository[T]) Create(obj *T) error {
    return r.db.Create(obj).Error
}

func (r *Repository[T]) Update(obj *T) error {
    return r.db.Save(obj).Error
}

func (r *Repository[T]) Delete(id uint) error {
    return r.db.Delete(new(T), id).Error
}

// DB exposes the underlying *gorm.DB for advanced queries
func (r *Repository[T]) DB() *gorm.DB {
    return r.db
}
