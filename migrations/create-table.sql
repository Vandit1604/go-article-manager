DROP TABLE IF EXISTS "articles";
CREATE TABLE articles (
  id         SERIAL PRIMARY KEY,
  title      VARCHAR(128) NOT NULL,
  content    VARCHAR(255) NOT NULL,
  date       DATE 
);

INSERT INTO articles
  (title, content, date)
VALUES
('Effective Go', 'Go is a new language...', CURRENT_DATE),
('The Secret', 'The Secret’s principles for manifestation...', CURRENT_DATE),
('Rohan', 'Gandi Baatey Karta Hai', CURRENT_DATE),
('Effective Go', 'Go is a new language...', CURRENT_DATE),
('The Secret', 'The Secret’s principles for manifestation...', CURRENT_DATE),
('Rohan', 'Gandi Baatey Karta Hai', CURRENT_DATE);

