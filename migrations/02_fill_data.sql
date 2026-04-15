-- TODO: remove when user auth service is ready
INSERT INTO users (username, password)
VALUES
  ('user1', 'lasjdfijw234023jd'),
  ('user2', 'oiwenb203isdfa'),
  ('AI_tictactoe', 'alsdjflkajsdfowksdjf');

INSERT INTO users (user_id, username, password)
VALUES
  ('00000000-0000-0000-0000-000000000000', 'empty_user', 'asjdlfj23j2ofjlasjd'),
  ('11111111-1111-1111-1111-111111111111', 'draw_user', 'asjdlfj23j2ofjlasjd');

INSERT INTO game_types (name)
VALUES ('TicTacToe');
