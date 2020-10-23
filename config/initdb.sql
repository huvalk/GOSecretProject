CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    login TEXT NOT NULL DEFAULT '',
    password TEXT NOT NULL DEFAULT '',
    phone TEXT NOT NULL DEFAULT ''
);

CREATE TABLE IF NOT EXISTS session (
    user_id INT NOT NULL,
    session_id TEXT NOT NULL DEFAULT '' UNIQUE
);

CREATE TABLE IF NOT EXISTS recipe (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    title TEXT NOT NULL,
    cooking_time INTEGER,
    ingredients TEXT[],
    steps TEXT[],
    CONSTRAINT user_fkey FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS rating (
    user_id INTEGER NOT NULL,
    recipe_id INTEGER NOT NULL,
    stars INTEGER NOT NULL,
    CONSTRAINT user_fkey FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT recipe_fkey FOREIGN KEY (recipe_id) REFERENCES recipe(id)
);

CREATE TABLE IF NOT EXISTS favorites (
    user_id INTEGER NOT NULL,
    recipe_id INTEGER NOT NULL,
    CONSTRAINT user_fkey FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT recipe_fkey FOEIGN KEY (recipe_id) REFERENCES recipe(id)
);