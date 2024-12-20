CREATE TABLE song_unique (
    id smallserial PRIMARY KEY NOT NULL,
    group_id integer,
    song_id integer,
    FOREIGN KEY (group_id) REFERENCES groups (id),
    FOREIGN KEY (song_id) REFERENCES songs (id)
);