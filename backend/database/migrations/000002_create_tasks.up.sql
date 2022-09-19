CREATE TABLE IF NOT EXISTS tasks(
    id serial PRIMARY KEY,
    user_id integer NOT NULL,
    title VARCHAR (50) UNIQUE NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
    FOREIGN KEY user_id references users(user_id)
);