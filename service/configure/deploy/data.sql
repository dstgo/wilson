-- ----------------------------
-- Table structure for gorm_init
-- ----------------------------
DROP TABLE IF EXISTS `gorm_init`;
CREATE TABLE `gorm_init`
(
    `id`   INT(0) UNSIGNED NOT NULL AUTO_INCREMENT,
    `init` TINYINT(1)      NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = `utf8mb4`
  COLLATE = `utf8mb4_0900_ai_ci`
  ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for server
-- ----------------------------
DROP TABLE IF EXISTS `server`;
CREATE TABLE `server`
(
    `id`          INT(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `created_at`  BIGINT(0)       NULL DEFAULT NULL COMMENT '创建时间',
    `updated_at`  BIGINT(0)       NULL DEFAULT NULL COMMENT '修改时间',
    `keyword`     CHAR(32)        NOT NULL COMMENT '服务标识',
    `name`        VARCHAR(64)     NOT NULL COMMENT '服务名称',
    `description` VARCHAR(128)    NOT NULL COMMENT '服务描述',
    `status`      TINYINT(1)      NULL DEFAULT 0,
    PRIMARY KEY (`id`) USING BTREE,
    INDEX `idx_server_created_at` (`created_at`) USING BTREE,
    INDEX `idx_server_updated_at` (`updated_at`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = `utf8mb4`
  COLLATE = `utf8mb4_0900_ai_ci` COMMENT = '服务信息';

-- ----------------------------
-- Table structure for env
-- ----------------------------
DROP TABLE IF EXISTS `env`;
CREATE TABLE `env`
(
    `id`          INT(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `created_at`  BIGINT(0)       NULL DEFAULT NULL COMMENT '创建时间',
    `updated_at`  BIGINT(0)       NULL DEFAULT NULL COMMENT '修改时间',
    `keyword`     CHAR(32)        NOT NULL COMMENT '环境标识',
    `name`        VARCHAR(64)     NOT NULL COMMENT '环境名称',
    `description` VARCHAR(128)    NOT NULL COMMENT '环境描述',
    `token`       VARCHAR(128)    NOT NULL COMMENT '环境token',
    `status`      TINYINT(1)      NULL DEFAULT 0 COMMENT '环境状态',
    PRIMARY KEY (`id`) USING BTREE,
    INDEX `idx_env_created_at` (`created_at`) USING BTREE,
    INDEX `idx_env_updated_at` (`updated_at`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = `utf8mb4`
  COLLATE = `utf8mb4_0900_ai_ci` COMMENT = '环境信息'
  ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for business
-- ----------------------------
DROP TABLE IF EXISTS `business`;
CREATE TABLE `business`
(
    `id`          INT(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `created_at`  BIGINT(0)       NULL DEFAULT NULL COMMENT '创建时间',
    `updated_at`  BIGINT(0)       NULL DEFAULT NULL COMMENT '修改时间',
    `server_id`   INT(0) UNSIGNED NOT NULL COMMENT '服务id',
    `keyword`     CHAR(32)        NOT NULL COMMENT '变量标识',
    `type`        VARCHAR(32)     NOT NULL COMMENT '变量类型',
    `description` VARCHAR(128)    NOT NULL COMMENT '变量描述',
    PRIMARY KEY (`id`) USING BTREE,
    INDEX `idx_business_created_at` (`created_at`) USING BTREE,
    INDEX `idx_business_updated_at` (`updated_at`) USING BTREE,
    INDEX `fk_business_server` (`server_id`) USING BTREE,
    CONSTRAINT `fk_business_server` FOREIGN KEY (`server_id`) REFERENCES `server` (`id`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE = InnoDB
  CHARACTER SET = `utf8mb4`
  COLLATE = `utf8mb4_0900_ai_ci` COMMENT = '业务变量'
  ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for business_value
-- ----------------------------
DROP TABLE IF EXISTS `business_value`;
CREATE TABLE `business_value`
(
    `id`          INT(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `created_at`  BIGINT(0)       NULL DEFAULT NULL COMMENT '创建时间',
    `updated_at`  BIGINT(0)       NULL DEFAULT NULL COMMENT '修改时间',
    `env_id`      INT(0) UNSIGNED NOT NULL COMMENT '环境id',
    `business_id` INT(0) UNSIGNED NOT NULL COMMENT '业务变量id',
    `value`       TEXT            NOT NULL COMMENT '业务变量id',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `env_id` (`env_id`, `business_id`) USING BTREE,
    INDEX `idx_business_value_created_at` (`created_at`) USING BTREE,
    INDEX `idx_business_value_updated_at` (`updated_at`) USING BTREE,
    INDEX `fk_business_value_env` (`env_id`) USING BTREE,
    INDEX `fk_business_values` (`business_id`) USING BTREE,
    CONSTRAINT `fk_business_value_env` FOREIGN KEY (`env_id`) REFERENCES `env` (`id`) ON DELETE CASCADE ON UPDATE RESTRICT,
    CONSTRAINT `fk_business_values` FOREIGN KEY (`business_id`) REFERENCES `business` (`id`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE = InnoDB
  CHARACTER SET = `utf8mb4`
  COLLATE = `utf8mb4_0900_ai_ci` COMMENT = '业务变量值'
  ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for configure
-- ----------------------------
DROP TABLE IF EXISTS `configure`;
CREATE TABLE `configure`
(
    `id`          INT(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `created_at`  BIGINT(0)       NULL DEFAULT NULL COMMENT '创建时间',
    `updated_at`  BIGINT(0)       NULL DEFAULT NULL COMMENT '修改时间',
    `server_id`   INT(0) UNSIGNED NOT NULL COMMENT '服务id',
    `env_id`      INT(0) UNSIGNED NOT NULL COMMENT '环境id',
    `content`     TEXT            NOT NULL COMMENT '配置内容',
    `version`     VARCHAR(32)     NOT NULL COMMENT '配置版本',
    `format`      VARCHAR(32)     NOT NULL COMMENT '配置格式',
    `description` VARCHAR(128)    NULL DEFAULT NULL COMMENT '配置描述',
    PRIMARY KEY (`id`) USING BTREE,
    INDEX `idx_configure_updated_at` (`updated_at`) USING BTREE,
    INDEX `idx_configure_created_at` (`created_at`) USING BTREE,
    INDEX `fk_configure_server` (`server_id`) USING BTREE,
    INDEX `fk_configure_env` (`env_id`) USING BTREE,
    CONSTRAINT `fk_configure_env` FOREIGN KEY (`env_id`) REFERENCES `env` (`id`) ON DELETE CASCADE ON UPDATE RESTRICT,
    CONSTRAINT `fk_configure_server` FOREIGN KEY (`server_id`) REFERENCES `server` (`id`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE = InnoDB
  CHARACTER SET = `utf8mb4`
  COLLATE = `utf8mb4_0900_ai_ci` COMMENT = '配置内容'
  ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for resource
-- ----------------------------
DROP TABLE IF EXISTS `resource`;
CREATE TABLE `resource`
(
    `id`          INT(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `created_at`  BIGINT(0)       NULL DEFAULT NULL COMMENT '创建时间',
    `updated_at`  BIGINT(0)       NULL DEFAULT NULL COMMENT '修改时间',
    `keyword`     CHAR(32)        NOT NULL COMMENT '变量标识',
    `description` VARCHAR(128)    NOT NULL COMMENT '变量描述',
    `fields`      VARCHAR(256)    NOT NULL COMMENT '变量字段集合',
    `tag`         VARCHAR(32)     NOT NULL COMMENT '变量标签',
    `private`     TINYINT(1)      NULL DEFAULT 0 COMMENT '是否私有',
    PRIMARY KEY (`id`) USING BTREE,
    INDEX `idx_resource_created_at` (`created_at`) USING BTREE,
    INDEX `idx_resource_updated_at` (`updated_at`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = `utf8mb4`
  COLLATE = `utf8mb4_0900_ai_ci` COMMENT = '资源变量'
  ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for resource_server
-- ----------------------------
DROP TABLE IF EXISTS `resource_server`;
CREATE TABLE `resource_server`
(
    `id`          INT(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `created_at`  BIGINT(0)       NULL DEFAULT NULL COMMENT '创建时间',
    `server_id`   INT(0) UNSIGNED NOT NULL COMMENT '服务id',
    `resource_id` INT(0) UNSIGNED NOT NULL COMMENT '资源id',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `sr` (`server_id`, `resource_id`) USING BTREE,
    INDEX `idx_resource_server_created_at` (`created_at`) USING BTREE,
    INDEX `fk_resource_resource_server` (`resource_id`) USING BTREE,
    CONSTRAINT `fk_resource_resource_server` FOREIGN KEY (`resource_id`) REFERENCES `resource` (`id`) ON DELETE CASCADE ON UPDATE RESTRICT,
    CONSTRAINT `fk_resource_server_server` FOREIGN KEY (`server_id`) REFERENCES `server` (`id`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE = InnoDB
  CHARACTER SET = `utf8mb4`
  COLLATE = `utf8mb4_0900_ai_ci` COMMENT = '资源服务信息'
  ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for resource_value
-- ----------------------------
DROP TABLE IF EXISTS `resource_value`;
CREATE TABLE `resource_value`
(
    `id`          INT(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `created_at`  BIGINT(0)       NULL DEFAULT NULL COMMENT '创建时间',
    `updated_at`  BIGINT(0)       NULL DEFAULT NULL COMMENT '修改时间',
    `env_id`      INT(0) UNSIGNED NOT NULL COMMENT '环境id',
    `resource_id` INT(0) UNSIGNED NOT NULL COMMENT '资源id',
    `value`       TEXT            NOT NULL COMMENT '资源id',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `er` (`env_id`, `resource_id`) USING BTREE,
    INDEX `idx_resource_value_updated_at` (`updated_at`) USING BTREE,
    INDEX `idx_resource_value_created_at` (`created_at`) USING BTREE,
    INDEX `fk_resource_resource_value` (`resource_id`) USING BTREE,
    CONSTRAINT `fk_resource_resource_value` FOREIGN KEY (`resource_id`) REFERENCES `resource` (`id`) ON DELETE CASCADE ON UPDATE RESTRICT,
    CONSTRAINT `fk_resource_value_env` FOREIGN KEY (`env_id`) REFERENCES `env` (`id`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE = InnoDB
  CHARACTER SET = `utf8mb4`
  COLLATE = `utf8mb4_0900_ai_ci` COMMENT = '资源变量值'
  ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for template
-- ----------------------------
DROP TABLE IF EXISTS `template`;
CREATE TABLE `template`
(
    `id`          INT(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `created_at`  BIGINT(0)       NULL DEFAULT NULL COMMENT '创建时间',
    `updated_at`  BIGINT(0)       NULL DEFAULT NULL COMMENT '修改时间',
    `server_id`   INT(0) UNSIGNED NOT NULL COMMENT '模板id',
    `content`     TEXT            NOT NULL COMMENT '模板内容',
    `version`     VARCHAR(32)     NOT NULL COMMENT '模板版本',
    `is_use`      TINYINT(1)      NULL DEFAULT 0 COMMENT '是否使用',
    `format`      VARCHAR(32)     NOT NULL COMMENT '模板格式',
    `description` VARCHAR(128)    NOT NULL COMMENT '模板描述',
    `compare`     TEXT            NOT NULL COMMENT '变更详情',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `sv` (`server_id`, `version`) USING BTREE,
    INDEX `idx_template_created_at` (`created_at`) USING BTREE,
    INDEX `idx_template_updated_at` (`updated_at`) USING BTREE,
    CONSTRAINT `fk_template_server` FOREIGN KEY (`server_id`) REFERENCES `server` (`id`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE = InnoDB
  CHARACTER SET = `utf8mb4`
  COLLATE = `utf8mb4_0900_ai_ci` COMMENT = '配置模板';