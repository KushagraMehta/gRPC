-- Select Blog database

    -- \c CRUD;

-- Create User Table
    CREATE TABLE IF NOT EXISTS USERS(
        ID          SERIAL PRIMARY KEY,
        FNAME       VARCHAR(20) ,
        CITY        VARCHAR(20) ,
        PHONE       BIGINT NOT NULl UNIQUE,
        HEIGHT      FLOAT,
        MARRIED     BOOLEAN
    );
