package queries

const CreateURLTable string = `
	CREATE TABLE IF NOT EXISTS url (
		id INTEGER PRIMARY KEY,
		alias TEXT NOT NULL UNIQUE,
		url TEXT NOT NULL
	);
	CREATE INDEX IF NOT EXISTS idx_alias ON url (alias);
`

const CreateURL string = `INSERT INTO url(url, alias) VALUES(?, ?)`

const GetURLByAlias string = `SELECT url FROM url WHERE alias = ?`
