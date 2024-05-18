DROP TABLE IF EXISTS "articles";
CREATE TABLE articles (
  id         SERIAL PRIMARY KEY,
  title      VARCHAR NOT NULL,
  content    TEXT NOT NULL,
  date       DATE 
);


