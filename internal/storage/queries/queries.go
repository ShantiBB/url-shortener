package queries

const CreateURLTable string = `
	CREATE TABLE IF NOT EXISTS url (
		id INTEGER GENERATED ALWAYS AS IDENTITY,
		alias TEXT NOT NULL UNIQUE,
		url TEXT NOT NULL
	);
	CREATE INDEX IF NOT EXISTS idx_alias ON url (alias);
`

const CreateURL string = `INSERT INTO url(url, alias) VALUES($1, $2) RETURNING id`

const GetURLByAlias string = `SELECT url FROM url WHERE alias = $1`

const DeleteURLByAlias string = `DELETE FROM url WHERE alias  = $1`
