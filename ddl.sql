CREATE DATABASE IF NOT EXISTS monitor DEFAULT CHAR SET utf8mb4 COLLATE utf8mb4_general_ci;


CREATE TABLE m_users
(
    id           INT UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
    user_name    VARCHAR(100) NOT NULL COMMENT '用户名',
    password     VARCHAR(100) NOT NULL COMMENT '密码',
    status       TINYINT(1)   NOT NULL DEFAULT 0 COMMENT '状态',
    created_time  DATETIME              DEFAULT NULL COMMENT '创建时间',
    updated_time DATETIME              DEFAULT NULL COMMENT '修改时间',
    deleted_time DATETIME              DEFAULT NULL COMMENT '删除时间',
    unique index udx_user_name (user_name),
    index idx_deleted_time (deleted_time)
) ENGINE INNODB
  default char set utf8mb4
  collate utf8mb4_general_ci COMMENT '后台用户表';