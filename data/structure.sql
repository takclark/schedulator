CREATE TABLE IF NOT EXISTS rules (
   id INTEGER PRIMARY KEY,
   name TEXT NOT NULL,
   rule_type TEXT NOT NULL,
   expression TEXT NOT NULL,
   url TEXT NOT NULL,
   method TEXT NOT NULL,
   body TEXT NULL
);
