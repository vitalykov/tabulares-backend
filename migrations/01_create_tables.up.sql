CREATE TABLE users (
	user_id    UUID PRIMARY KEY          DEFAULT uuidv7(),
	login      VARCHAR(50)      NOT NULL CHECK (octet_length(login) >= 3),
	password   TEXT             NOT NULL CHECK (octet_length(password) >= 8),
	created_at TIMESTAMP        NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE game_types (
	game_type_id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	name         VARCHAR(50) NOT NULL
);

CREATE TABLE games (
	game_id         UUID PRIMARY KEY          DEFAULT uuidv7(),
	type            INTEGER          NOT NULL REFERENCES game_types(game_type_id) ON DELETE RESTRICT,
	board_width     INTEGER          NOT NULL,
	board_height    INTEGER          NOT NULL,
	winner          UUID                      REFERENCES users(user_id) ON DELETE SET NULL,
	additional_info TEXT,
	created_at      TIMESTAMP        NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE players (
	game_id   UUID    NOT NULL REFERENCES games(game_id) ON DELETE CASCADE,
	player_id UUID    NOT NULL REFERENCES users(user_id) ON DELETE SET NULL,
	position  INTEGER NOT NULL,
	PRIMARY KEY (game_id, player_id),
	UNIQUE (game_id, position)
);

CREATE TABLE moves (
	game_id   UUID    REFERENCES games(game_id) ON DELETE CASCADE,
	player_id UUID    REFERENCES users(user_id) ON DELETE SET NULL,
	move      TEXT,
	position  INTEGER,
	PRIMARY KEY (game_id, player_id, move),
	UNIQUE (game_id, position)
);
