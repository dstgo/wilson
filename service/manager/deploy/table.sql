SET FOREIGN_KEY_CHECKS = 0;

CREATE TABLE IF NOT EXISTS `casbin_rule`
(
    `id`    BIGINT(0) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `ptype` VARCHAR(100)       NULL DEFAULT NULL,
    `v0`    VARCHAR(100)       NULL DEFAULT NULL,
    `v1`    VARCHAR(100)       NULL DEFAULT NULL,
    `v2`    VARCHAR(100)       NULL DEFAULT NULL,
    `v3`    VARCHAR(100)       NULL DEFAULT NULL,
    `v4`    VARCHAR(100)       NULL DEFAULT NULL,
    `v5`    VARCHAR(100)       NULL DEFAULT NULL,

    UNIQUE INDEX `idx_casbin_rule` (`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) USING BTREE
    ) ENGINE = InnoDB
    CHARACTER SET = `utf8mb4`;


CREATE TABLE `department`
(
    `id`          BIGINT UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `parent_id`   BIGINT UNSIGNED                NOT NULL,
    `keyword`     CHAR(32)                       NOT NULL,
    `name`        VARCHAR(64)                    NOT NULL,
    `description` VARCHAR(256)                   NOT NULL,
    `created_at`  BIGINT UNSIGNED DEFAULT 0      NOT NULL,
    `updated_at`  BIGINT UNSIGNED DEFAULT 0      NOT NULL,

    CONSTRAINT `keyword` UNIQUE (`keyword`),
    INDEX `idx_department_created_at` (`created_at`),
    INDEX `idx_department_updated_at` (`updated_at`)
) ENGINE = InnoDB
  CHARACTER SET = `utf8mb4`;

CREATE TABLE `department_closure`
(
    `id`       BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `parent`   BIGINT UNSIGNED NOT NULL,
    `children` BIGINT UNSIGNED NOT NULL,

    CONSTRAINT `department_closure_ibfk_1` FOREIGN KEY (`children`) REFERENCES `department` (`id`) ON DELETE CASCADE,
    CONSTRAINT `department_closure_ibfk_2` FOREIGN KEY (`parent`) REFERENCES `department` (`id`) ON DELETE CASCADE,
    INDEX `children` (`children`),
    INDEX `parent` (`parent`)
) ENGINE = InnoDB
  CHARACTER SET = `utf8mb4`;

CREATE TABLE `dictionary`
(
    `id`          BIGINT UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `keyword`     CHAR(32)                       NOT NULL,
    `name`        VARCHAR(64)                    NOT NULL,
    `description` VARCHAR(256)                   NOT NULL,
    `created_at`  BIGINT UNSIGNED DEFAULT 0      NOT NULL,
    `updated_at`  BIGINT UNSIGNED DEFAULT 0      NOT NULL,
    `deleted_at`  BIGINT UNSIGNED DEFAULT 0      NOT NULL,

    CONSTRAINT `keyword` UNIQUE (`keyword`, `deleted_at`),
    INDEX `idx_dictionary_created_at` (`created_at`),
    INDEX `idx_dictionary_deleted_at` (`deleted_at`),
    INDEX `idx_dictionary_updated_at` (`updated_at`)
) ENGINE = InnoDB
  CHARACTER SET = `utf8mb4`;

CREATE TABLE `dictionary_value`
(
    `id`            BIGINT UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `dictionary_id` BIGINT UNSIGNED                NOT NULL,
    `label`         VARCHAR(128)                   NOT NULL,
    `value`         VARCHAR(128)                   NOT NULL,
    `status`        TINYINT(1)      DEFAULT 1      NOT NULL,
    `weight`        INT UNSIGNED    DEFAULT 0      NOT NULL,
    `type`          CHAR(32)                       NOT NULL,
    `extra`         VARCHAR(512)                   NOT NULL,
    `description`   VARCHAR(256)                   NOT NULL,
    `created_at`    BIGINT UNSIGNED DEFAULT 0      NOT NULL,
    `updated_at`    BIGINT UNSIGNED DEFAULT 0      NOT NULL,

    CONSTRAINT `value` UNIQUE (`dictionary_id`, `value`),
    CONSTRAINT `fk_dictionary_value_dict` FOREIGN KEY (`dictionary_id`) REFERENCES `dictionary` (`id`) ON DELETE CASCADE,
    INDEX `idx_dictionary_value_created_at` (`created_at`),
    INDEX `idx_dictionary_value_updated_at` (`updated_at`),
    INDEX `idx_dictionary_value_weight` (`weight`)
) ENGINE = InnoDB
  CHARACTER SET = `utf8mb4`;

CREATE TABLE `menu`
(
    `id`         BIGINT UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `parent_id`  BIGINT UNSIGNED                NOT NULL,
    `title`      VARCHAR(128)                   NOT NULL,
    `type`       CHAR(32)                       NOT NULL,
    `keyword`    VARCHAR(64)                    NULL,
    `icon`       CHAR(32)                       NULL,
    `api`        VARCHAR(128)                   NULL,
    `method`     VARCHAR(12)                    NULL,
    `path`       VARCHAR(128)                   NULL,
    `permission` VARCHAR(128)                   NULL,
    `component`  VARCHAR(128)                   NULL,
    `redirect`   VARCHAR(128)                   NULL,
    `weight`     INT UNSIGNED    DEFAULT 0      NULL,
    `is_hidden`  TINYINT(1)                     NULL,
    `is_cache`   TINYINT(1)                     NULL,
    `is_home`    TINYINT(1)                     NULL,
    `is_affix`   TINYINT(1)                     NULL,
    `created_at` BIGINT UNSIGNED DEFAULT 0      NOT NULL,
    `updated_at` BIGINT UNSIGNED DEFAULT 0      NOT NULL,

    CONSTRAINT `api_method` UNIQUE (`api`, `method`),
    CONSTRAINT `keyword` UNIQUE (`keyword`),
    CONSTRAINT `path` UNIQUE (`path`),
    INDEX `idx_menu_created_at` (`created_at`),
    INDEX `idx_menu_updated_at` (`updated_at`),
    INDEX `idx_menu_weight` (`weight`)
) ENGINE = InnoDB
  CHARACTER SET = `utf8mb4`;

CREATE TABLE `menu_closure`
(
    `id`       BIGINT UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `parent`   BIGINT UNSIGNED                NOT NULL,
    `children` BIGINT UNSIGNED                NOT NULL,

    CONSTRAINT `menu_closure_ibfk_1` FOREIGN KEY (`children`) REFERENCES `menu` (`id`) ON DELETE CASCADE,
    CONSTRAINT `menu_closure_ibfk_2` FOREIGN KEY (`parent`) REFERENCES `menu` (`id`) ON DELETE CASCADE,
    INDEX `idx_menu_closure_children` (`children`),
    INDEX `idx_menu_closure_parent` (`parent`)
) ENGINE = InnoDB
  CHARACTER SET = `utf8mb4`;

CREATE TABLE `resource`
(
    `id`            BIGINT UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `keyword`       VARCHAR(32)                    NOT NULL,
    `department_id` BIGINT UNSIGNED                NOT NULL,
    `resource_id`   BIGINT UNSIGNED                NOT NULL,

    CONSTRAINT `department_id` UNIQUE (`keyword`, `department_id`, `resource_id`),
    CONSTRAINT `resource_ibfk_1` FOREIGN KEY (`department_id`) REFERENCES `department` (`id`) ON DELETE CASCADE,
    INDEX `idx_resource_department_id` (`department_id`)
) ENGINE = InnoDB
  CHARACTER SET = `utf8mb4`;

CREATE TABLE `role`
(
    `id`             BIGINT UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `parent_id`      BIGINT UNSIGNED                NOT NULL,
    `name`           VARCHAR(64)                    NOT NULL,
    `keyword`        CHAR(32)                       NOT NULL,
    `status`         TINYINT(1)      DEFAULT 0      NOT NULL,
    `description`    VARCHAR(128)                   NOT NULL,
    `department_ids` TINYTEXT                       NULL,
    `data_scope`     CHAR(32)                       NOT NULL,
    `created_at`     BIGINT UNSIGNED DEFAULT 0      NOT NULL,
    `updated_at`     BIGINT UNSIGNED DEFAULT 0      NOT NULL,
    `deleted_at`     BIGINT UNSIGNED DEFAULT 0      NOT NULL,

    CONSTRAINT `keyword` UNIQUE (`keyword`, `deleted_at`),
    INDEX `idx_role_created_at` (`created_at`),
    INDEX `idx_role_deleted_at` (`deleted_at`),
    INDEX `idx_role_updated_at` (`updated_at`)
) ENGINE = InnoDB
  CHARACTER SET = `utf8mb4`;

CREATE TABLE `role_closure`
(
    `id`       BIGINT UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `parent`   BIGINT UNSIGNED                NOT NULL,
    `children` BIGINT UNSIGNED                NOT NULL,

    CONSTRAINT `role_closure_ibfk_1` FOREIGN KEY (`children`) REFERENCES `role` (`id`) ON DELETE CASCADE,
    CONSTRAINT `role_closure_ibfk_2` FOREIGN KEY (`parent`) REFERENCES `role` (`id`) ON DELETE CASCADE,
    INDEX `idx_role_closure_children` (`children`),
    INDEX `idx_role_closure_parent` (`parent`)
) ENGINE = InnoDB
  CHARACTER SET = `utf8mb4`;

CREATE TABLE `role_menu`
(
    `id`      BIGINT UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `role_id` BIGINT UNSIGNED                NOT NULL,
    `menu_id` BIGINT UNSIGNED                NOT NULL,

    CONSTRAINT `role_id_2` UNIQUE (`role_id`, `menu_id`),
    CONSTRAINT `role_menu_ibfk_1` FOREIGN KEY (`menu_id`) REFERENCES `menu` (`id`) ON DELETE CASCADE,
    CONSTRAINT `role_menu_ibfk_2` FOREIGN KEY (`role_id`) REFERENCES `role` (`id`) ON DELETE CASCADE,
    INDEX `idx_role_menu_menu_id` (`menu_id`),
    INDEX `idx_role_menu_role_id` (`role_id`)
) ENGINE = InnoDB
  CHARACTER SET = `utf8mb4`;

CREATE TABLE `user`
(
    `id`            BIGINT UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `department_id` BIGINT UNSIGNED                NOT NULL,
    `role_id`       BIGINT UNSIGNED                NOT NULL,
    `name`          CHAR(32)                       NOT NULL,
    `nickname`      VARCHAR(64)                    NOT NULL,
    `gender`        CHAR(32)                       NOT NULL,
    `avatar`        VARCHAR(256)                   NULL,
    `email`         VARCHAR(64)                    NOT NULL,
    `phone`         CHAR(32)                       NOT NULL,
    `password`      VARCHAR(256)                   NOT NULL,
    `status`        TINYINT(1) DEFAULT 0           NULL,
    `setting`       TINYTEXT                       NULL,
    `token`         VARCHAR(512)                   NULL,
    `logged_at`     BIGINT     DEFAULT 0           NOT NULL,
    `created_at`    BIGINT     DEFAULT 0           NOT NULL,
    `updated_at`    BIGINT     DEFAULT 0           NOT NULL,

    CONSTRAINT `fk_user_department` FOREIGN KEY (`department_id`) REFERENCES `department` (`id`),
    CONSTRAINT `fk_user_role` FOREIGN KEY (`role_id`) REFERENCES `role` (`id`),
    INDEX `idx_user_created_at` (`created_at`),
    INDEX `idx_user_updated_at` (`updated_at`)
) ENGINE = InnoDB
  CHARACTER SET = `utf8mb4`;

CREATE TABLE `user_role`
(
    `id`      BIGINT UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `role_id` BIGINT UNSIGNED                NOT NULL,
    `user_id` BIGINT UNSIGNED                NOT NULL,

    CONSTRAINT `role_id` UNIQUE (`role_id`, `user_id`),
    CONSTRAINT `user_role_ibfk_1` FOREIGN KEY (`role_id`) REFERENCES `role` (`id`) ON DELETE CASCADE,
    CONSTRAINT `user_role_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE,
    INDEX `idx_user_role_user_id` (`user_id`)
) ENGINE = InnoDB
  CHARACTER SET = `utf8mb4`;

SET FOREIGN_KEY_CHECKS = 1;