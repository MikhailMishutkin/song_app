CREATE INDEX idx_songs_song ON songs(song);
CREATE INDEX idx_groups_group ON groups(group_name);
CREATE INDEX idx_details_release_date ON details(release_date);
CREATE INDEX idx_details_text ON details(text);
CREATE INDEX idx_details_link ON details(link);
CREATE INDEX idx_details_release_date_text_link ON details(release_date, text, link);
CREATE INDEX idx_details_text_link ON details(text, link);