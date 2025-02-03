SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for chunk
-- ----------------------------
DROP TABLE IF EXISTS `chunk`;
CREATE TABLE `chunk`
(
    `id`         BIGINT(0)                                                         NOT NULL AUTO_INCREMENT COMMENT '分片id',
    `upload_id`  VARCHAR(128) CHARACTER SET `utf8mb4` COLLATE `utf8mb4_bin`        NOT NULL COMMENT '上传id',
    `index`      BIGINT(0)                                                         NOT NULL COMMENT '切片下标',
    `size`       BIGINT(0)                                                         NOT NULL COMMENT '切片大小',
    `sha`        VARCHAR(128) CHARACTER SET `utf8mb4` COLLATE `utf8mb4_0900_ai_ci` NOT NULL COMMENT '切片sha',
    `data`       MEDIUMBLOB                                                        NOT NULL COMMENT '切片数据',
    `created_at` BIGINT(0) UNSIGNED                                                NULL DEFAULT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `upload_id` (`upload_id`, `index`) USING BTREE,
    UNIQUE INDEX `ui` (`upload_id`, `index`) USING BTREE,
    INDEX `idx_chunk_created_at` (`created_at`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = `utf8mb4`
  COLLATE = `utf8mb4_0900_ai_ci`;

-- ----------------------------
-- Table structure for directory
-- ----------------------------
DROP TABLE IF EXISTS `directory`;
CREATE TABLE `directory`
(
    `id`         BIGINT(0) UNSIGNED                                               NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `parent_id`  BIGINT(0) UNSIGNED                                               NOT NULL COMMENT '父id',
    `name`       VARCHAR(64) CHARACTER SET `utf8mb4` COLLATE `utf8mb4_0900_ai_ci` NOT NULL COMMENT '目录名称',
    `accept`     TINYTEXT CHARACTER SET `utf8mb4` COLLATE `utf8mb4_0900_ai_ci`    NOT NULL COMMENT '允许后缀',
    `max_size`   BIGINT(0) UNSIGNED                                               NOT NULL COMMENT '最大大小',
    `created_at` BIGINT(0)                                                        NULL DEFAULT NULL COMMENT '创建时间',
    `updated_at` BIGINT(0)                                                        NULL DEFAULT NULL COMMENT '修改时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `pna` (`parent_id`, `name`) USING BTREE,
    INDEX `idx_directory_created_at` (`created_at`) USING BTREE,
    INDEX `idx_directory_updated_at` (`updated_at`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = `utf8mb4`
  COLLATE = `utf8mb4_0900_ai_ci` COMMENT = '目录信息';

-- ----------------------------
-- Table structure for directory_closure
-- ----------------------------
DROP TABLE IF EXISTS `directory_closure`;
CREATE TABLE `directory_closure`
(
    `id`       BIGINT(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `parent`   BIGINT(0) UNSIGNED NOT NULL COMMENT '目录id',
    `children` BIGINT(0) UNSIGNED NOT NULL COMMENT '目录id',
    PRIMARY KEY (`id`) USING BTREE,
    INDEX `parent` (`parent`) USING BTREE,
    INDEX `children` (`children`) USING BTREE,
    CONSTRAINT `directory_closure_ibfk_1` FOREIGN KEY (`children`) REFERENCES `directory` (`id`) ON DELETE CASCADE ON UPDATE RESTRICT,
    CONSTRAINT `directory_closure_ibfk_2` FOREIGN KEY (`parent`) REFERENCES `directory` (`id`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE = InnoDB
  CHARACTER SET = `utf8mb4`
  COLLATE = `utf8mb4_0900_ai_ci` COMMENT = '目录层级信息';

-- ----------------------------
-- Table structure for export
-- ----------------------------
DROP TABLE IF EXISTS `export`;
CREATE TABLE `export`
(
    `id`            BIGINT(0)                                                         NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `scene`         CHAR(32) CHARACTER SET `utf8mb4` COLLATE `utf8mb4_bin`            NOT NULL COMMENT '场景',
    `name`          VARCHAR(128) CHARACTER SET `utf8mb4` COLLATE `utf8mb4_0900_ai_ci` NOT NULL COMMENT '名称',
    `size`          BIGINT(0)                                                         NOT NULL DEFAULT 0 COMMENT '大小',
    `sha`           VARCHAR(64) CHARACTER SET `utf8mb4` COLLATE `utf8mb4_bin`         NOT NULL COMMENT '版本',
    `src`           VARCHAR(128) CHARACTER SET `utf8mb4` COLLATE `utf8mb4_bin`        NULL     DEFAULT NULL COMMENT '路径',
    `reason`        VARCHAR(512) CHARACTER SET `utf8mb4` COLLATE `utf8mb4_0900_ai_ci` NULL     DEFAULT NULL COMMENT '错误原因',
    `status`        VARCHAR(32) CHARACTER SET `utf8mb4` COLLATE `utf8mb4_0900_ai_ci`  NOT NULL DEFAULT '' COMMENT '状态',
    `user_id`       BIGINT(0) UNSIGNED                                                NULL     DEFAULT NULL COMMENT '创建人',
    `department_id` BIGINT(0) UNSIGNED                                                NULL     DEFAULT NULL COMMENT '创建部门',
    `expired_at`    BIGINT(0) UNSIGNED                                                NULL     DEFAULT NULL COMMENT '过期时间',
    `created_at`    BIGINT(0) UNSIGNED                                                NULL     DEFAULT NULL COMMENT '创建时间',
    `updated_at`    BIGINT(0) UNSIGNED                                                NULL     DEFAULT NULL COMMENT '修改时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `sha` (`sha`, `user_id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = `utf8mb4`
  COLLATE = `utf8mb4_0900_ai_ci` COMMENT = '导出任务';

-- ----------------------------
-- Table structure for file
-- ----------------------------
DROP TABLE IF EXISTS `file`;
CREATE TABLE `file`
(
    `id`           BIGINT(0) UNSIGNED                                                                 NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `directory_id` BIGINT(0) UNSIGNED                                                                 NOT NULL COMMENT '目录id',
    `name`         VARCHAR(128) CHARACTER SET `utf8mb4` COLLATE `utf8mb4_bin`                         NOT NULL COMMENT '文件名',
    `type`         VARCHAR(64) CHARACTER SET `utf8mb4` COLLATE `utf8mb4_0900_ai_ci`                   NOT NULL COMMENT '文件类型',
    `size`         BIGINT(0)                                                                          NOT NULL COMMENT '文件大小',
    `sha`          VARCHAR(128) CHARACTER SET `utf8mb4` COLLATE `utf8mb4_bin`                         NOT NULL COMMENT 'sha值',
    `key`          VARCHAR(128) CHARACTER SET `utf8mb4` COLLATE `utf8mb4_bin`                         NOT NULL COMMENT 'key值',
    `src`          VARCHAR(256) CHARACTER SET `utf8mb4` COLLATE `utf8mb4_bin`                         NULL DEFAULT NULL COMMENT '文件路径',
    `status`       ENUM ('PROGRESS','COMPLETED') CHARACTER SET `utf8mb4` COLLATE `utf8mb4_0900_ai_ci` NULL DEFAULT 'PROGRESS' COMMENT '上传状态',
    `upload_id`    VARCHAR(128) CHARACTER SET `utf8mb4` COLLATE `utf8mb4_bin`                         NULL DEFAULT NULL COMMENT '上传id',
    `chunk_count`  INT(0)                                                                             NULL DEFAULT 1 COMMENT '切片数量',
    `created_at`   BIGINT(0) UNSIGNED                                                                 NULL DEFAULT NULL COMMENT '创建时间',
    `updated_at`   BIGINT(0) UNSIGNED                                                                 NULL DEFAULT NULL COMMENT '修改时间',
    `deleted_at`   BIGINT(0) UNSIGNED                                                                 NULL DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `sha` (`sha`, `directory_id`) USING BTREE,
    UNIQUE INDEX `upload_id` (`upload_id`) USING BTREE,
    INDEX `deleted_at` (`deleted_at`) USING BTREE,
    INDEX `directory_id` (`directory_id`) USING BTREE,
    CONSTRAINT `file_ibfk_1` FOREIGN KEY (`directory_id`) REFERENCES `directory` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB
  CHARACTER SET = `utf8mb4`
  COLLATE = `utf8mb4_0900_ai_ci`;

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
  COLLATE = `utf8mb4_0900_ai_ci`;

SET FOREIGN_KEY_CHECKS = 1;