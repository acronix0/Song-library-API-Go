CREATE TABLE groups (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE songs (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    group_id INT NOT NULL REFERENCES groups(id) ON DELETE CASCADE,
    link VARCHAR(500),
    release_date DATE,
    text TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE song_lyrics (
    id SERIAL PRIMARY KEY,
    song_id INT NOT NULL REFERENCES songs(id) ON DELETE CASCADE,
    verse_number INT NOT NULL,
    text TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    UNIQUE (song_id, verse_number)
);

CREATE UNIQUE INDEX idx_groups_name ON groups (name);

CREATE INDEX idx_songs_group_id ON songs (group_id);
CREATE INDEX idx_songs_created_at ON songs (created_at);
CREATE INDEX idx_songs_title ON songs (title);

CREATE INDEX idx_song_lyrics_song_id ON song_lyrics (song_id);
CREATE INDEX idx_song_lyrics_verse_number ON song_lyrics (verse_number);
CREATE UNIQUE INDEX idx_song_lyrics_song_id_verse_number ON song_lyrics (song_id, verse_number);
