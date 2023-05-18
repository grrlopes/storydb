package sqlite

const table string = `
	CREATE TABLE IF NOT EXISTS command(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL UNIQUE,
		description TEXT NOT NULL
	);
	`
