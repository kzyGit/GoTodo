package service

import (
    "context"

    "../db"
    "../schema"
)
// This service has only simple functions that call functions with the same name in the db package

func Close(ctx context.Context) {
    db.Close(ctx)
}

func Insert(ctx context.Context, todo *schema.Todo) (int, error) {
    return db.Insert(ctx, todo)
}

func Delete(ctx context.Context, id int) error {
    return db.Delete(ctx, id)
}

func GetAll(ctx context.Context) ([]schema.Todo, error) {
    return db.GetAll(ctx)
}