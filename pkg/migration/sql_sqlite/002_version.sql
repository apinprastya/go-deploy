-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE versions (
	id INTEGER PRIMARY KEY,
	created_at DATETIME DEFAULT (datetime('now', 'localtime')),
    project_id INTEGER NOT NULL,
    version INTEGER NOT NULL,
    version_number VARCHAR(255) NOT NULL,
    live BOOLEAN NOT NULL
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE versions;