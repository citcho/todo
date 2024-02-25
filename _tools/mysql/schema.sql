CREATE TABLE `users`
(
    `ulid` CHAR(26) NOT NULL COMMENT 'ユーザー識別子',
    `name` VARCHAR(255) NOT NULL COMMENT 'ユーザー名',
    `email` VARCHAR(255) NOT NULL COMMENT 'メールアドレス',
    `password` VARCHAR(255) NOT NULL COMMENT 'パスワードハッシュ',
    `created_at` DATETIME NOT NULL COMMENT '作成日時',
    `updated_at` DATETIME NOT NULL COMMENT '更新日時',
    `deleted_at` DATETIME COMMENT '削除日時',
    PRIMARY KEY (`ulid`),
    UNIQUE KEY `uix_email_01` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ユーザー';

CREATE TABLE `todos`
(
    `ulid` CHAR(26) NOT NULL COMMENT 'Todo識別子',
    `user_ulid` CHAR(26) NOT NULL COMMENT 'Todoを作成したユーザーの識別子',
    `title` VARCHAR(255) NOT NULL COMMENT 'Todoのタイトル',
    `content` TEXT COMMENT 'Todoの内容',
    `completed` INT NOT NULL COMMENT 'Todoの状態',
    `created_at` DATETIME NOT NULL COMMENT '作成日時',
    `updated_at` DATETIME NOT NULL COMMENT '更新日時',
    `deleted_at` DATETIME COMMENT '削除日時',
    PRIMARY KEY (`ulid`),
    CONSTRAINT `fk_user_ulid`
        FOREIGN KEY (`user_ulid`) REFERENCES `users` (`ulid`)
            ON DELETE RESTRICT ON UPDATE RESTRICT
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Todo';
