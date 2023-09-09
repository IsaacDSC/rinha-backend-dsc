
CREATE TABLE IF NOT EXISTS person (
    id UUID PRIMARY KEY,
    username VARCHAR(32) UNIQUE NOT NULL, 
    name VARCHAR(100) NOT NULL,
    birthday VARCHAR(11) NOT NULL,
    updated_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS stack (
    stack_name VARCHAR(255) NOT NULL,
	person_id UUID NOT NULL,
	FOREIGN KEY(person_id) REFERENCES person(id)
);


CREATE INDEX index_username ON person(username);
CREATE INDEX index_name ON person(name);
CREATE INDEX index_birthday ON person(birthday);
CREATE INDEX index_created_at ON person(created_at);
