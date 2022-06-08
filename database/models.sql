DROP TABLE IF EXISTS founds;
DROP TABLE IF EXISTS persons;


CREATE TABLE IF NOT EXISTS persons (
    id serial NOT NULL,
    first_name VARCHAR(150) NOT NULL,
    last_name VARCHAR(150) NOT NULL,
    curp VARCHAR(150) UNIQUE,
    birthdate date,
    last_seen DATE,
    missing bool DEFAULT 'true',
    created_at timestamp DEFAULT now(),
    updated_at timestamp NOT NULL,
    CONSTRAINT pk_persons PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS founds (
    id serial NOT NULL,
    person_id smallint references persons(id),
    found_date DATE,
    address VARCHAR(150),
    hospitalized bool DEFAULT 'false',
    condition VARCHAR(150),
    more_info VARCHAR(150),
    created_at timestamp DEFAULT now(),
    updated_at timestamp NOT NULL,
    CONSTRAINT pk_founds PRIMARY KEY(id)
);