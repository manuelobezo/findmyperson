/*DROP TABLE IF EXISTS persons;*/

CREATE TABLE IF NOT EXISTS persons (
    id serial NOT NULL,
    first_name VARCHAR(150) NOT NULL,
    last_name VARCHAR(150) NOT NULL,
    curp VARCHAR(150) UNIQUE,
    found bool NOT NULL DEFAULT 'false',
    missing bool DEFAULT 'false',
    birthdate date,
    created_at timestamp DEFAULT now(),
    updated_at timestamp NOT NULL,
    CONSTRAINT pk_persons PRIMARY KEY(id)
);
