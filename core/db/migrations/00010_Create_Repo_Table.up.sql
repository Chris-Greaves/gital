CREATE TABLE Repositories (
  name TEXT,
  path TEXT UNIQUE,
  current_branch TEXT,
  last_updated INTEGER
);