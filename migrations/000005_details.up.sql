CREATE TABLE details (
    uniq_id integer PRIMARY KEY NOT NULL,
    release_date timestamp,
    text varchar,
    link varchar,
    FOREIGN KEY (uniq_id) REFERENCES song_unique (id)
)