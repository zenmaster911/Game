BEGIN;

CREATE TABLE users_characters
(
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    character_id INT REFERENCES characters(id) ON DELETE CASCADE,
    PRIMARY KEY(user_id,character_id)
);


COMMIT;