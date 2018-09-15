package main_test

import (
    "os"
    "testing"

    "."
)

var a main.App

func TestMain(m *testing.M) {
    a = main.App{}
    a.Initialize( "postgres","appointy","dbname" )

    ensureTableExists()

    code := m.Run()

    clearTable()

    os.Exit(code)
}

func ensureTableExists() {
    if _, err := a.DB.Exec(tableCreationQuery); err != nil {
        log.Fatal(err)
    }
}

func clearTable() {
    a.DB.Exec("DELETE FROM todolists")
    a.DB.Exec("ALTER SEQUENCE products_id_seq RESTART WITH 1")
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS todolists
(
    id NUMERIC PRIMARY KEY,
    title varchar(200) NOT NULL ,
    created_at date NOT NULL DEFAULT 0.00,
    updated_at date ,
    completed BOOLEAN 
)`


