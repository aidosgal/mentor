CREATE TABLE users (
       id SERIAL PRIMARY KEY,
       first_name VARCHAR(255) NOT NULL,
       last_name VARCHAR(255) NOT NULL,
       username VARCHAR(255) NOT NULL,
       chat_id VARCHAR(255) NOT NULL,
       role VARCHAR(255) NOT NULL,
       description VARCHAR(10000) NULL,
       phone VARCHAR(255) NULL,
       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
