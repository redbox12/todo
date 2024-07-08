CREATE TABLE IF NOT EXISTS users
(
    id serial CONSTRAINT users_id_pk PRIMARY KEY,
    name varchar(255) not null,
    username varchar(255) not null CONSTRAINT username_key UNIQUE,
    password_hash varchar(255) not null
    );

CREATE TABLE IF NOT EXISTS todo_lists
(
    id serial CONSTRAINT todo_lists_id PRIMARY KEY,
    title varchar(255) not null,
    description varchar(255)
    );

CREATE TABLE IF NOT EXISTS users_lists
(
    id serial CONSTRAINT users_lists_id PRIMARY KEY,
    user_id INTEGER not null,
    list_id INTEGER not null,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (list_id) REFERENCES todo_lists (id) ON DELETE CASCADE
    );

CREATE TABLE IF NOT EXISTS todo_items
(
    id serial CONSTRAINT todo_items_id_pk PRIMARY KEY,
    title varchar(255) not null,
    description varchar(255),
    done boolean not null default false
    );

CREATE TABLE IF NOT EXISTS lists_items
(
    id serial CONSTRAINT lists_items_id_pk PRIMARY KEY,
    item_id INTEGER not null,
    list_id INTEGER not null,
    FOREIGN KEY (item_id) REFERENCES todo_items (id) ON DELETE CASCADE,
    FOREIGN KEY (list_id) REFERENCES todo_lists (id) ON DELETE CASCADE
    );