CREATE TABLE `users`
(
    `ulid` CHAR(26) NOT NULL COMMENT 'ユーザー識別子',
    `name` VARCHAR(255) NOT NULL COMMENT 'ユーザー名',
    `password` VARCHAR(255) NOT NULL COMMENT 'パスワードハッシュ',
    `email` VARCHAR(255) NOT NULL COMMENT 'メールアドレス',
    `created_at` DATETIME NOT NULL COMMENT '作成日時',
    `updated_at` DATETIME NOT NULL COMMENT '更新日時',
    `deleted_at` DATETIME COMMENT '削除日時',
    PRIMARY KEY (`ulid`),
    UNIQUE KEY `uix_email_01` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ユーザー';
