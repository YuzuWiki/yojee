CREATE TABLE IF NOT EXISTS pixiv_user (
    id          BIGINT       AUTO_INCREMENT     PRIMARY KEY,
    create_at   TIMESTAMP    DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    update_at   TIMESTAMP    DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT  '创建时间',
    is_delete   BOOL         DEFAULT FALSE       COMMENT '是否删除',

    uuid        VARCHAR(127) DEFAULT ''         COMMENT 'pixiv uuid',
    pid         BIGINT       NOT NULL           COMMENT 'user id',
    name        VARCHAR(256) NOT NULL           COMMENT 'user name',
    nick_name   VARCHAR(256) DEFAULT ''         COMMENT 'nick name',
    avatar      VARCHAR(512) DEFAULT ''         COMMENT 'user avatar',
    following   INT          DEFAULT 0          COMMENT '关注数量',
    followers   INT          DEFAULT 0          COMMENT '粉丝数量',

    INDEX idx_pid (pid)
    ) CHARACTER SET utf8;


CREATE TABLE IF NOT EXISTS pixiv_follow (
    id  BIGINT  AUTO_INCREMENT  PRIMARY KEY ,
    create_at   TIMESTAMP    DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    update_at   TIMESTAMP    DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT  '创建时间',
    is_delete   BOOL         DEFAULT FALSE       COMMENT '是否删除',

    pid             BIGINT       NOT NULL           COMMENT 'user id',
    followed_pid    BIGINT       NOT NULL           COMMENT 'user id',
    followed_at     TIMESTAMP    NOT NULL            COMMENT '关注时间',

    INDEX idx_followed (pid, followed_pid)
)  CHARACTER SET utf8;

CREATE TABLE IF NOT EXISTS pixiv_tag_group (
    # 标签组, tage name 为日语
    id  BIGINT  AUTO_INCREMENT  PRIMARY KEY ,
    create_at   TIMESTAMP    DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    update_at   TIMESTAMP    DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT  '创建时间',
    is_delete   BOOL         DEFAULT FALSE       COMMENT '是否删除',

    tag_name        VARBINARY(512)  NOT NULL COMMENT 'tage name',

    INDEX idx_tag (tag_name)
)  CHARACTER SET utf8;

CREATE TABLE IF NOT EXISTS pixiv_tag (
    # 标签组, tage name 为日语
    id  BIGINT  AUTO_INCREMENT  PRIMARY KEY ,
    create_at   TIMESTAMP    DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    update_at   TIMESTAMP    DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT  '创建时间',
    is_delete   BOOL         DEFAULT FALSE       COMMENT '是否删除',

    tag_group   BIGINT          NOT NULL COMMENT 'tag group id',
    language    VARBINARY(16)   NOT NULL COMMENT 'tag标签语种',
    tag_name    VARBINARY(512)  NOT NULL COMMENT 'tage name',

    INDEX idx_tag_group (tag_group, tag_name)
)  CHARACTER SET utf8;
