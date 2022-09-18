CREATE TABLE IF NOT EXISTS users(
    user_id serial PRIMARY KEY,
    username VARCHAR (50) UNIQUE NOT NULL,
    password VARCHAR (50) NOT NULL,
    email VARCHAR (300) UNIQUE NOT NULL
);


-- mysql
-- CREATE TABLE IF NOT EXISTS users(
--     user_id NOT NULL AUTO_INCREMENT PRIMARY KEY,
--     username VARCHAR (50) NOT NULL,
--     password VARCHAR (50) NOT NULL,
--     email VARCHAR (300) NOT NULL
-- );