CREATE TABLE profiles (
       id SERIAL PRIMARY KEY,
       user_id BIGINT NOT NULL,
       price VARCHAR(255) NOT NULL,
       description VARCHAR(10000) NOT NULL,
       category_id BIGINT NOT NULL,
       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE reviews (
       id SERIAL PRIMARY KEY,
       user_id BIGINT NOT NULL,
       profile_id BIGINT NOT NULL,
       makr BIGINT NOT NULL,
       content VARCHAR(1000) NULL,
       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
