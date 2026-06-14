-- +goose Up
CREATE TABLE IF NOT EXISTS images (
    image_id   UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id    UUID REFERENCES users(user_id) ON DELETE SET NULL,
    filename   TEXT NOT NULL,
    file_path  TEXT NOT NULL,
    file_size  BIGINT NOT NULL,
    mime_type  TEXT NOT NULL,
    width      INT,
    height     INT,
    alt_text   TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS albums (
    album_id       UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id        UUID REFERENCES users(user_id) ON DELETE SET NULL,
    title          VARCHAR(255) NOT NULL,
    description    TEXT,
    cover_image_id UUID REFERENCES images(image_id) ON DELETE SET NULL,
    is_public      BOOLEAN NOT NULL DEFAULT false,
    created_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at     TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS album_images (
    album_id   UUID NOT NULL REFERENCES albums(album_id) ON DELETE CASCADE,
    image_id   UUID NOT NULL REFERENCES images(image_id) ON DELETE CASCADE,
    position   INT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (album_id, image_id)
);

CREATE INDEX idx_images_user_id ON images(user_id);
CREATE INDEX idx_images_deleted_at ON images(deleted_at);
CREATE INDEX idx_albums_user_id ON albums(user_id);
CREATE INDEX idx_albums_deleted_at ON albums(deleted_at);
CREATE INDEX idx_albums_cover ON albums(cover_image_id);
CREATE INDEX idx_album_images_image_id ON album_images(image_id);
CREATE INDEX idx_album_images_ordering ON album_images(album_id, position);

-- +goose Down
DROP TABLE IF EXISTS album_images;
DROP TABLE IF EXISTS albums;
DROP TABLE IF EXISTS images;
