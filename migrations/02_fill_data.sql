-- TODO: remove when user auth service is ready
INSERT INTO users (login, password)
VALUES
  ('user1', 'lasjdfijw234023jd'),
  ('user2', 'oiwenb203isdfa'),
  ('AI_tictactoe', 'alsdjflkajsdfowksdjf');

INSERT INTO users (user_id, login, password)
VALUES
  ('00000000-0000-0000-0000-000000000000', 'empty_user', 'asjdlfj23j2ofjlasjd');

INSERT INTO game_types (name)
VALUES ('TicTacToe');
