DROP INDEX IF EXISTS idx_song_lyrics_song_id_verse_number;
DROP INDEX IF EXISTS idx_song_lyrics_verse_number;
DROP INDEX IF EXISTS idx_song_lyrics_song_id;

DROP INDEX IF EXISTS idx_songs_title;
DROP INDEX IF EXISTS idx_songs_created_at;
DROP INDEX IF EXISTS idx_songs_group_id;

DROP INDEX IF EXISTS idx_groups_name;

DROP TABLE IF EXISTS song_lyrics;

DROP TABLE IF EXISTS songs;

DROP TABLE IF EXISTS groups;
