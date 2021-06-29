CREATE TABLE m_user
(
    id           INT UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
    username     VARCHAR(100)                        NOT NULL COMMENT '用户名',
    password     VARCHAR(100)                        NOT NULL COMMENT '密码',
    status       TINYINT(1)                          NOT NULL DEFAULT 0 COMMENT '状态',
    created_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '创建时间 ',
    updated_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    unique index udx_username (username)
) ENGINE INNODB
  default char set utf8mb4
  collate utf8mb4_general_ci COMMENT '后台用户表';

-- 插入默认用户
INSERT INTO m_user (username, password, status, created_time, updated_time) VALUES ('admin', '$2a$10$hKqAhWZqzMM7LeiFpLbWM.njlNcmLtHuO5HVv.R3/MGGYP4pp5CrC', 0, null, null);
