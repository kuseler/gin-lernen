<<<<<<< HEAD
CREATE DATABASE recipes;
\c recipes;
CREATE TABLE users (
                       username varchar(16) PRIMARY KEY NOT NULL,
                       password varchar(32)
);

CREATE TABLE recipes (
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    creator varchar(16) references users(username),
    title varchar(32),
    content varchar(4096)
)

-- insert
INSERT INTO users VALUES ('Clark', 'Sales');
INSERT INTO users VALUES ('Dave', 'Accounting');
INSERT INTO users VALUES ('Ava', 'Sales');

-- fetch
-- SELECT * FROM users WHERE username = 'Clark';


=======
CREATE DATABASE recipes;
\c recipes;
CREATE TABLE users (
                       username varchar(16) PRIMARY KEY NOT NULL,
                       password varchar(32)
);

CREATE TABLE recipes (
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    creator varchar(16) references users(username),
    title varchar(32),
    content varchar(4096)
)

-- insert
INSERT INTO users VALUES ('Clark', 'Sales');
INSERT INTO users VALUES ('Dave', 'Accounting');
INSERT INTO users VALUES ('Ava', 'Sales');

-- fetch
-- SELECT * FROM users WHERE username = 'Clark';


>>>>>>> 43dfc5aefbd0f2da92f9237fb13c126e96d4f5c1
