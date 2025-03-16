SET FOREIGN_KEY_CHECKS = 0;

CREATE TABLE `chunk`
(
    `id`         BIGINT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `upload_id`  VARCHAR(128)          NOT NULL,
    `index`      BIGINT                NOT NULL,
    `size`       BIGINT                NOT NULL,
    `sha`        VARCHAR(128)          NOT NULL,
    `data`       MEDIUMBLOB            NOT NULL,
    `created_at` BIGINT UNSIGNED       NOT NULL,

    CONSTRAINT `ui` UNIQUE (`upload_id`, `index`),
    CONSTRAINT `upload_id` UNIQUE (`upload_id`, `index`),
    INDEX `idx_chunk_created_at` (`created_at`)
) ENGINE = InnoDB
  CHARACTER SET = `utf8mb4`;

CREATE TABLE `directory`
(
    `id`         BIGINT UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `parent_id`  BIGINT UNSIGNED                NOT NULL,
    `name`       VARCHAR(64)                    NOT NULL,
    `accept`     TINYTEXT                       NOT NULL,
    `max_size`   BIGINT UNSIGNED                NOT NULL,
    `created_at` BIGINT                         NOT NULL,
    `updated_at` BIGINT                         NOT NULL,

    CONSTRAINT `pna` UNIQUE (`parent_id`, `name`),
    INDEX `idx_directory_created_at` (`created_at`),
    INDEX `idx_directory_updated_at` (`updated_at`)
) ENGINE = InnoDB
  CHARACTER SET = `utf8mb4`;

CREATE TABLE `directory_closure`
(
    `id`       BIGINT UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `parent`   BIGINT UNSIGNED                NOT NULL,
    `children` BIGINT UNSIGNED                NOT NULL,

    CONSTRAINT `directory_closure_ibfk_1` FOREIGN KEY (`children`) REFERENCES `directory` (`id`) ON DELETE CASCADE,
    CONSTRAINT `directory_closure_ibfk_2` FOREIGN KEY (`parent`) REFERENCES `directory` (`id`) ON DELETE CASCADE,
    INDEX `idx_directory_closure_children` (`children`),
    INDEX `idx_directory_closure_parent` (`parent`)
) ENGINE = InnoDB
  CHARACTER SET = `utf8mb4`;

CREATE TABLE `export`
(
    `id`            BIGINT AUTO_INCREMENT  NOT NULL PRIMARY KEY,
    `scene`         CHAR(32)               NOT NULL,
    `name`          VARCHAR(128)           NOT NULL,
    `size`          BIGINT      DEFAULT 0  NOT NULL,
    `sha`           VARCHAR(64)            NOT NULL,
    `src`           VARCHAR(128)           NULL,
    `reason`        VARCHAR(512)           NULL,
    `status`        VARCHAR(32) DEFAULT '' NOT NULL,
    `user_id`       BIGINT UNSIGNED        NULL,
    `department_id` BIGINT UNSIGNED        NULL,
    `expired_at`    BIGINT UNSIGNED        NULL,
    `created_at`    BIGINT UNSIGNED        NOT NULL,
    `updated_at`    BIGINT UNSIGNED        NOT NULL,

    CONSTRAINT `sha` UNIQUE (`sha`, `user_id`)
) ENGINE = InnoDB
  CHARACTER SET = `utf8mb4`;

CREATE TABLE `file`
(
    `id`           BIGINT UNSIGNED AUTO_INCREMENT                    NOT NULL PRIMARY KEY,
    `directory_id` BIGINT UNSIGNED                                   NOT NULL,
    `name`         VARCHAR(128)                                      NOT NULL,
    `type`         VARCHAR(64)                                       NOT NULL,
    `size`         BIGINT                                            NOT NULL,
    `sha`          VARCHAR(128)                                      NOT NULL,
    `key`          VARCHAR(128)                                      NOT NULL,
    `src`          VARCHAR(256)                                      NULL,
    `status`       ENUM ('PROGRESS', 'COMPLETED') DEFAULT 'PROGRESS' NULL,
    `upload_id`    VARCHAR(128)                                      NULL,
    `chunk_count`  INT                            DEFAULT 1          NULL,
    `created_at`   BIGINT UNSIGNED                                   NOT NULL,
    `updated_at`   BIGINT UNSIGNED                                   NOT NULL,
    `deleted_at`   BIGINT UNSIGNED                                   NULL,

    CONSTRAINT `sha` UNIQUE (`sha`, `directory_id`),
    CONSTRAINT `upload_id` UNIQUE (`upload_id`),
    CONSTRAINT `file_ibfk_1` FOREIGN KEY (`directory_id`) REFERENCES `directory` (`id`),
    INDEX `idx_file_deleted_at` (`deleted_at`),
    INDEX `idx_file_directory_id` (`directory_id`)
) ENGINE = InnoDB
  CHARACTER SET = `utf8mb4`;

SET FOREIGN_KEY_CHECKS = 1;