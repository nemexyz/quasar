CREATE TABLE satellites (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) unique NOT NULL,
    x INTEGER NOT NULL,
    y INTEGER NOT NULL
);
CREATE TABLE messages (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    distance NUMERIC(20, 10) NOT NULL,
    date TIMESTAMP NOT NULL,
    satellite_id BIGINT NOT NULL,
    CONSTRAINT fk_satellite FOREIGN KEY(satellite_id) REFERENCES satellites(id)
);
CREATE TABLE words (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    word VARCHAR(255) NOT NULL,
    position INTEGER NOT NULL,
    message_id BIGINT NOT NULL,
    CONSTRAINT fk_message FOREIGN KEY(message_id) REFERENCES messages(id)
);
INSERT INTO satellites(name, x, y)
VALUES ('kenobi', -500, -200);
INSERT INTO satellites(name, x, y)
VALUES ('skywalker', 100, -100);
INSERT INTO satellites(name, x, y)
VALUES ('sato', 500, 100);