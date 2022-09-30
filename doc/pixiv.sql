CREATE TABLE IF NOT EXISTS pixiv_account
(
    id         BIGINT AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP    DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP    DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
    is_deleted BOOL         DEFAULT FALSE COMMENT '是否删除',

    pid             BIGINT       NOT NULL    COMMENT 'pixiv user id',
    name            VARCHAR(256) NOT NULL    COMMENT 'pixiv user name',
    avatar          VARCHAR(512) DEFAULT ''  COMMENT 'pixiv user avatar',
    region          VARCHAR(128) DEFAULT ''  COMMENT 'pixiv user region',
    gender          VARCHAR(128) DEFAULT ''  COMMENT 'pixiv user gender',
    birth_day       VARCHAR(128) DEFAULT ''  COMMENT 'pixiv user birth_day',
    Job             VARCHAR(128) DEFAULT ''  COMMENT 'pixiv user Job',
    following       INT          DEFAULT 0   COMMENT '关注数量',
    art_updated     BIGINT       DEFAULT 0   COMMENT '最后一次更新时间戳',

    INDEX idx_pan (pid, art_updated, name)
) CHARACTER SET utf8mb4;


CREATE TABLE IF NOT EXISTS pixiv_follow
(
    id           BIGINT AUTO_INCREMENT PRIMARY KEY,
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
    is_deleted   BOOL      DEFAULT FALSE COMMENT '是否删除',

    pid          BIGINT    NOT NULL COMMENT 'user id',
    followed_pid BIGINT    NOT NULL COMMENT '被关注者的 uid id',

    INDEX idx_followed (pid, followed_pid)
) CHARACTER SET utf8mb4;


CREATE TABLE IF NOT EXISTS pixiv_tag
(
    id         BIGINT AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
    is_deleted BOOL      DEFAULT FALSE COMMENT '是否删除',

    tag_id      BIGINT       NOT NULl COMMENT 'tag id',
    jp          VARCHAR(512) NOT NULL COMMENT 'tage name, jp',
    en          VARCHAR(512) NOT NULL COMMENT 'tage name, en',
    ko          VARCHAR(512) NOT NULL COMMENT 'tage name, ko',
    zh          VARCHAR(512) NOT NULL COMMENT 'tage name, zh',
    romaji      VARCHAR(512) NOT NULL COMMENT 'tage romaji, jp',
    image       VARCHAR(512) DEFAULT ''  COMMENT 'tag image',

    UNIQUE idx_tag_id(tag_id),
    INDEX idx_tag (jp, tag_id)
) CHARACTER SET utf8mb4;


CREATE TABLE IF NOT EXISTS pixiv_tag_tree
(
    # 日语tag标签层级关系
    id         BIGINT AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
    is_deleted BOOL      DEFAULT FALSE  COMMENT '是否删除',

    parent_id  BIGINT          NOT NULL    COMMENT 'parent tag, id',
    parent_jp  VARCHAR(512)    NOT NULL    COMMENT 'parent tag, jp (name)',

    tag_id  BIGINT          NOT NULL    COMMENT 'curr tag, id',
    tag_jp  VARCHAR(512)    NOT NULL    COMMENT 'curr tag, jp (name)',

    UNIQUE INDEX idx_tree (parent_id, tag_id)
) CHARACTER SET utf8mb4;


CREATE TABLE IF  NOT EXISTS pixiv_artwork_tag (
    id               BIGINT AUTO_INCREMENT PRIMARY KEY,
    created_at       TIMESTAMP  DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at       TIMESTAMP  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
    is_deleted       BOOL       DEFAULT FALSE COMMENT '是否删除',

    pid             BIGINT      NOT NULL    COMMENT 'pixiv user id',
    art_id          BIGINT      NOT NULL    COMMENT 'artwork id',
    art_type        VARCHAR(64) NOT NULL    COMMENT 'artwork type',
    tag_id          BIGINT      NOT NULL    COMMENT 'tag id',

    INDEX idx_art (pid, tag_id),
    INDEX idx_art_tag (art_type, art_id, tag_id)
) CHARACTER SET utf8mb4;


CREATE TABLE IF NOT EXISTS pixiv_artwork
(
    id             BIGINT AUTO_INCREMENT PRIMARY KEY,
    created_at     TIMESTAMP    DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at     TIMESTAMP    DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
    is_deleted     BOOL         DEFAULT FALSE COMMENT '是否删除',

    # 作品信息
    pid            BIGINT       NOT NULL COMMENT 'pixiv user id',
    art_id         BIGINT       NOT NULL COMMENT 'pixiv art_id',
    art_type       VARCHAR(64)  NOT NULL  COMMENT 'pixiv art_type',

    title          TEXT                COMMENT 'pixiv title',
    description    TEXT                COMMENT 'pixiv illust desc',
    page_count     BIGINT    DEFAULT 0 COMMENT '页数',
    view_count     BIGINT    DEFAULT 0 COMMENT '浏览数量',
    like_count     BIGINT    DEFAULT 0 COMMENT '喜欢数',
    bookmark_count BIGINT    DEFAULT 0 COMMENT '收藏欢数',
    create_date    TIMESTAMP NOT NULL  COMMENT '创建时间',

    UNIQUE INDEX idx_illust (pid, art_type, art_id),
    INDEX idx_art (art_type, art_id)
) CHARACTER SET utf8mb4;
