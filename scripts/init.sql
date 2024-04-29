-- scripts/init.sql

CREATE TABLE IF NOT EXISTS posts (
                                     id SERIAL PRIMARY KEY,
                                     title VARCHAR(255) NOT NULL,
                                     content TEXT,
                                     allow_comments BOOLEAN DEFAULT TRUE
);

CREATE TABLE IF NOT EXISTS comments (

                          id SERIAL PRIMARY KEY,
                          post_id INT REFERENCES posts(id) ON DELETE CASCADE,
    content TEXT,
    parent_id INT,
    FOREIGN KEY (parent_id) REFERENCES comments(id));

CREATE INDEX idx_post_id ON posts(id);


-- Индекс для поиска по заголовку
CREATE INDEX idx_post_title ON posts(title);
