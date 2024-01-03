package sqlite

const table string = `
CREATE VIRTUAL TABLE IF NOT EXISTS commands_fts USING fts5(
  id,
  cmd,
  desc,
  updated_at,
  created_at,
  deleted_at,
  content='commands',
  content_rowid='id',
  tokenize="trigram case_sensitive 1",
);

CREATE TRIGGER IF NOT EXISTS commands_insert AFTER INSERT ON commands
  BEGIN
      INSERT INTO commands_fts (rowid, cmd, desc, updated_at, created_at, deleted_at)
      VALUES (new.id, new.cmd, new.desc, new.updated_at, new.created_at, new.deleted_at);
  END;
`
