

CREATE TABLE todolists
(
    id NUMERIC PRIMARY KEY,
    title varchar(200) NOT NULL ,
    created_at date NOT NULL DEFAULT 0.00,
    updated_at date ,
    completed BOOLEAN 
)


