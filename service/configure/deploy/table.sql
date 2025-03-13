SET FOREIGN_KEY_CHECKS = 0;

CREATE TABLE `business`
(
    `id`          INT UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `created_at`  BIGINT                      NULL,
    `updated_at`  BIGINT                      NULL,
    `server_id`   INT UNSIGNED                NOT NULL,
    `keyword`     CHAR(32)                    NOT NULL,
    `type`        VARCHAR(32)                 NOT NULL,
    `description` VARCHAR(128)                NOT NULL,

    CONSTRAINT `fk_business_server` FOREIGN KEY (`server_id`) REFERENCES `server` (`id`) ON DELETE CASCADE,
    INDEX `idx_business_created_at` (`created_at`),
    INDEX `idx_business_updated_at` (`updated_at`)
) ENGINE = InnoDB
  CHARACTER SET = `utf8mb4`;

CREATE TABLE `business_value`
(
    `id`          INT UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `created_at`  BIGINT                      NULL,
    `updated_at`  BIGINT                      NULL,
    `env_id`      INT UNSIGNED                NOT NULL,
    `business_id` INT UNSIGNED                NOT NULL,
    `value`       TEXT                        NOT NULL,

    CONSTRAINT `env_id` UNIQUE (`env_id`, `business_id`),
    CONSTRAINT `fk_business_value_env` FOREIGN KEY (`env_id`) REFERENCES `env` (`id`) ON DELETE CASCADE,
    CONSTRAINT `fk_business_values` FOREIGN KEY (`business_id`) REFERENCES `business` (`id`) ON DELETE CASCADE,
    INDEX `idx_business_value_created_at` (`created_at`),
    INDEX `idx_business_value_updated_at` (`updated_at`)
) ENGINE = InnoDB
  CHARACTER SET = `utf8mb4`;

CREATE TABLE `configure`
(
    `id`          INT UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `created_at`  BIGINT                      NULL,
    `updated_at`  BIGINT                      NULL,
    `server_id`   INT UNSIGNED                NOT NULL,
    `env_id`      INT UNSIGNED                NOT NULL,
    `content`     TEXT                        NOT NULL,
    `version`     VARCHAR(32)                 NOT NULL,
    `format`      VARCHAR(32)                 NOT NULL,
    `description` VARCHAR(128)                NULL,

    CONSTRAINT `fk_configure_env` FOREIGN KEY (`env_id`) REFERENCES `env` (`id`) ON DELETE CASCADE,
    CONSTRAINT `fk_configure_server` FOREIGN KEY (`server_id`) REFERENCES `server` (`id`) ON DELETE CASCADE,
    INDEX `idx_configure_created_at` (`created_at`),
    INDEX `idx_configure_updated_at` (`updated_at`)
) ENGINE = InnoDB
  CHARACTER SET = `utf8mb4`;

CREATE TABLE `env`
(
    `id`          INT UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `created_at`  BIGINT                      NULL,
    `updated_at`  BIGINT                      NULL,
    `keyword`     CHAR(32)                    NOT NULL,
    `name`        VARCHAR(64)                 NOT NULL,
    `description` VARCHAR(128)                NOT NULL,
    `token`       VARCHAR(128)                NOT NULL,
    `status`      TINYINT(1) DEFAULT 0        NULL,

    INDEX `idx_env_created_at` (`created_at`),
    INDEX `idx_env_updated_at` (`updated_at`)
) ENGINE = InnoDB
  CHARACTER SET = `utf8mb4`;

CREATE TABLE `resource`
(
    `id`          INT UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `created_at`  BIGINT                      NULL,
    `updated_at`  BIGINT                      NULL,
    `keyword`     CHAR(32)                    NOT NULL,
    `description` VARCHAR(128)                NOT NULL,
    `fields`      VARCHAR(256)                NOT NULL,
    `tag`         VARCHAR(32)                 NOT NULL,
    `private`     TINYINT(1) DEFAULT 0        NULL,

    INDEX `idx_resource_created_at` (`created_at`),
    INDEX `idx_resource_updated_at` (`updated_at`)
) ENGINE = InnoDB
  CHARACTER SET = `utf8mb4`;

CREATE TABLE `resource_server`
(
    `id`          INT UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `created_at`  BIGINT                      NULL,
    `server_id`   INT UNSIGNED                NOT NULL,
    `resource_id` INT UNSIGNED                NOT NULL,

    CONSTRAINT `sr` UNIQUE (`server_id`, `resource_id`),
    CONSTRAINT `fk_resource_resource_server` FOREIGN KEY (`resource_id`) REFERENCES `resource` (`id`) ON DELETE CASCADE,
    CONSTRAINT `fk_resource_server_server` FOREIGN KEY (`server_id`) REFERENCES `server` (`id`) ON DELETE CASCADE,
    INDEX `idx_resource_server_created_at` (`created_at`)
) ENGINE = InnoDB
  CHARACTER SET = `utf8mb4`;

CREATE TABLE `resource_value`
(
    `id`          INT UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `created_at`  BIGINT                      NULL,
    `updated_at`  BIGINT                      NULL,
    `env_id`      INT UNSIGNED                NOT NULL,
    `resource_id` INT UNSIGNED                NOT NULL,
    `value`       TEXT                        NOT NULL,

    CONSTRAINT `er` UNIQUE (`env_id`, `resource_id`),
    CONSTRAINT `fk_resource_resource_value` FOREIGN KEY (`resource_id`) REFERENCES `resource` (`id`) ON DELETE CASCADE,
    CONSTRAINT `fk_resource_value_env` FOREIGN KEY (`env_id`) REFERENCES `env` (`id`) ON DELETE CASCADE,
    INDEX `idx_resource_value_created_at` (`created_at`),
    INDEX `idx_resource_value_updated_at` (`updated_at`)
) ENGINE = InnoDB
  CHARACTER SET = `utf8mb4`;

CREATE TABLE `server`
(
    `id`          INT UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `created_at`  BIGINT                      NULL,
    `updated_at`  BIGINT                      NULL,
    `keyword`     CHAR(32)                    NOT NULL,
    `name`        VARCHAR(64)                 NOT NULL,
    `description` VARCHAR(128)                NOT NULL,
    `status`      TINYINT(1) DEFAULT 0        NULL,

    INDEX `idx_server_created_at` (`created_at`),
    INDEX `idx_server_updated_at` (`updated_at`)
) ENGINE = InnoDB
  CHARACTER SET = `utf8mb4`;

CREATE TABLE `template`
(
    `id`          INT UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `created_at`  BIGINT                      NULL,
    `updated_at`  BIGINT                      NULL,
    `server_id`   INT UNSIGNED                NOT NULL,
    `content`     TEXT                        NOT NULL,
    `version`     VARCHAR(32)                 NOT NULL,
    `is_use`      TINYINT(1) DEFAULT 0        NULL,
    `format`      VARCHAR(32)                 NOT NULL,
    `description` VARCHAR(128)                NOT NULL,
    `compare`     TEXT                        NOT NULL,

    CONSTRAINT `sv` UNIQUE (`server_id`, `version`),
    CONSTRAINT `fk_template_server` FOREIGN KEY (`server_id`) REFERENCES `server` (`id`) ON DELETE CASCADE,
    INDEX `idx_template_created_at` (`created_at`),
    INDEX `idx_template_updated_at` (`updated_at`)
) ENGINE = InnoDB
  CHARACTER SET = `utf8mb4`;

SET FOREIGN_KEY_CHECKS = 1;