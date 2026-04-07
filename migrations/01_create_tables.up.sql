DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS game_types;
DROP TABLE IF EXISTS games;
DROP TABLE IF EXISTS players; 
DROP TABLE IF EXISTS moves;

CREATE TABLE users (
	user_id 		UUID PRIMARY KEY 						DEFAULT uuidv7(),
	login 			TEXT 							NOT NULL 	CHECK (login <> ''),
	password 		TEXT 							NOT NULL 	CHECK (octet_length(password) >= 8),
	created_at 	TIMESTAMP 									DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE game_types (
	game_type_id 	INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	name 					VARCHAR(50) NOT NULL
);

CREATE TABLE games (
	game_id		 			UUID PRIMARY KEY 		DEFAULT uuidv7(),
	type 						INTEGER 						REFERENCES game_types(game_type_id),
	board_width 		INTEGER,
	board_height 		INTEGER,
	winner 					UUID 								REFERENCES users(user_id),
	additional_info TEXT,
	created_at		 	TIMESTAMP 					DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE players (
	game_id 	UUID 		REFERENCES games(game_id),
	player_id UUID 		REFERENCES users(user_id),
	position 	INTEGER,
	PRIMARY KEY (game_id, player_id),
	UNIQUE (game_id, position)
);

CREATE TABLE moves (
	game_id 	UUID 		REFERENCES games(game_id),
	player_id UUID 		REFERENCES users(user_id),
	move 			TEXT,
	position 	INTEGER,
	PRIMARY KEY (game_id, player_id, move),
	UNIQUE (game_id, position)
);
