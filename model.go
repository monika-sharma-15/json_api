package main


import (
    "database/sql"
    "errors"
    "time"
)

type Todolist struct {
    ID  int  `json:"id"`
  Title string  `json:"title" binding:"required"`
  CreatedAt time.Time  `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
  Completed bool  `json:"completed"`
}

func (p *Todolist) gettodolist(db *sql.DB) error {
    return db.QueryRow("SELECT title, created_at ,updated_at FROM todolists WHERE id=$1",
        p.ID).Scan(&p.Title, &p.CreatedAt, &p.UpdatedAt)
}


func getalllist(db *sql.DB, start, count int) ([]Todolist, error) {
    rows, err := db.Query(
       "SELECT title, created_at ,updated_at FROM todolists LIMIT $1 OFFSET $2",
        count, start)

    if err != nil {
        return nil, err
    }

    defer rows.Close()

    todolists := []Todolist{}

    for rows.Next() {
        var p Todolist
        if err := rows.Scan(&p.ID, &p.Title, &p.CreatedAt, &p.UpdatedAt, &p.Completed); err != nil {
            return nil, err
        }
        todolists = append(todolists, p)
    }

    return todolists, nil
}

func (p *Todolist) Addtodolist(db *sql.DB) error {
    err := db.QueryRow(
        "INSERT INTO todolists(title, created_at, updated_at, Completed) VALUES( bson.NewObjectId(), $2, $3,false ) RETURNING id",
        p.Title, p.CreatedAt, p.UpdatedAt).Scan(&p.ID)

    if err != nil {
        return err
    }

    return nil
}


func (p *Todolist) markascomplete(db *sql.DB) error {
    _, err :=
        db.Exec("UPDATE todolists SET completed='true' WHERE id=$3",
            p.Completed, p.ID)

    return err
}

func (p *Todolist) deletetodolist(db *sql.DB) error {
    _, err := db.Exec("DELETE FROM todolists WHERE id=$1", p.ID)

    return err
}

func deletealllist(db *sql.DB, start, count int) ([]Todolist, error) {
  return nil, errors.New("Not implemented")
}
