CREATE TABLE `users`
(
    `id` CHAR(26) NOT NULL COMMENT 'ユーザー識別子',
    `name` VARCHAR(255) NOT NULL COMMENT 'ユーザー名',
    `email` VARCHAR(255) NOT NULL COMMENT 'メールアドレス',
    `password` VARCHAR(255) NOT NULL COMMENT 'パスワードハッシュ',
    `created_at` DATETIME NOT NULL COMMENT '作成日時',
    `updated_at` DATETIME NOT NULL COMMENT '更新日時',
    `deleted_at` DATETIME COMMENT '削除日時',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uix_email_01` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ユーザー';

CREATE TABLE `todos`
(
    `id` CHAR(26) NOT NULL COMMENT 'Todo識別子',
    `user_id` CHAR(26) NOT NULL COMMENT 'Todoを作成したユーザーの識別子',
    `title` VARCHAR(255) NOT NULL COMMENT 'Todoのタイトル',
    `content` TEXT COMMENT 'Todoの内容',
    `is_complete` INT NOT NULL COMMENT 'Todoの状態',
    `created_at` DATETIME NOT NULL COMMENT '作成日時',
    `updated_at` DATETIME NOT NULL COMMENT '更新日時',
    `deleted_at` DATETIME COMMENT '削除日時',
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_user_id`
        FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
            ON DELETE RESTRICT ON UPDATE RESTRICT
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Todo';
