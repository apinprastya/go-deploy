-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE projects (
	id INTEGER PRIMARY KEY,
	created_at DATETIME DEFAULT (datetime('now', 'localtime')),
	updated_at DATETIME DEFAULT (datetime('now', 'localtime')),
	deleted_at DATETIME,
	name VARCHAR(255) NOT NULL
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE projects;