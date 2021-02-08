CREATE TABLE user (
    id INT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR (100) NOT NULL UNIQUE,
    password VARCHAR (100) NOT NULL,
    email VARCHAR (100) NOT NULL,
    phone VARCHAR (100) NOT NULL,
    state INT NOT NULL
)

INSERT INTO user(username, password, email, phone, state) VALUES('user', '12345678', 'user@qq.com', '11223344556', 1);
INSERT INTO user(username, password, email, phone, state) VALUES('user2', '11223344', 'user2@163.com', '11122233344', 1);
INSERT INTO user(username, password, email, phone, state) VALUES('user3', 'qweasdzx', 'user3@qq.com', '11112222333', 0);