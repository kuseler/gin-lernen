-- CREATE DATABASE epic_recipe_db;
--\c epic_recipe_db;
CREATE TABLE users (
                       username varchar(16) PRIMARY KEY NOT NULL,
                       password varchar(32)
);

CREATE TABLE recipes (
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    creator varchar(16) references users(username),
    title varchar(32),
    content varchar(4096)
);

-- insert

INSERT INTO users VALUES ('Dave', 'Accounting');
INSERT INTO users VALUES ('Ava', 'Sales');
INSERT INTO users VALUES ('Clark', 'Sales');

-- fetch
-- SELECT * FROM users WHERE username = 'Clark';


