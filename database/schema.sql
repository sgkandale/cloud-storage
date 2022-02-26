-- postgres

CREATE DATABASE cloud_storage WITH OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.UTF-8'
    LC_CTYPE = 'en_US.UTF-8'
    TEMPLATE = template0
    CONNECTION LIMIT = -1;

-- users sequence
CREATE SEQUENCE users_id_seq
    START WITH 100000000
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

-- users table
CREATE TABLE users (
    id BIGINT NOT NULL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at BIGINT NOT NULL
);

-- default user
INSERT INTO users (id, username, password, created_at)
    VALUES (100000000, 'admin', 'admin', 0);

-- objects sequence
CREATE SEQUENCE objects_id_seq
    START WITH 1000000000
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

-- objects table
CREATE TABLE objects (
    id BIGINT NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    parent_folder_id BIGINT NOT NULL,
    object_type VARCHAR(10) NOT NULL,
    created_at BIGINT NOT NULL,
    owner_id BIGINT NOT NULL,
    file_link VARCHAR(255)
);

-- insert default object
INSERT INTO objects (id, name, parent_folder_id, object_type, created_at, owner_id)
    VALUES (1000000000, 'system_root', 1000000000, 'DISK', 0, 100000000);

-- create foreign key
ALTER TABLE objects
    ADD CONSTRAINT fk_objects_owner_id
    FOREIGN KEY (owner_id)
    REFERENCES users (id)
    ON DELETE CASCADE;
ALTER TABLE objects
    ADD CONSTRAINT parent_folder_id_fk
    FOREIGN KEY (parent_folder_id)
    REFERENCES objects (id)
    ON DELETE CASCADE;
