-- ユーザーのテストデータ
INSERT INTO
    users (name, email, password)
VALUES
    ('ottotto','user1@example.com', 'password1'),
    ('yamato','user2@example.com', 'password2'),
    ('kizuku','user3@example.com', 'password3');

-- 投稿のテストデータ
INSERT INTO
    posts (content, user_id)
VALUES
    ('Hello, world!', 1),
    ('Learning Go and PostgreSQL', 2),
    ('Enjoying developing an SNS API', 3);

-- いいねのテストデータ
INSERT INTO likes (user_id, post_id) VALUES (1, 2), (2, 3), (3, 1);

-- フォローのテストデータ
INSERT INTO follows (follower_id, followee_id) VALUES (1, 2), (2, 3), (3, 1);