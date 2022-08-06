CREATE TABLE IF NOT EXISTS pixiv_user
(
    id         BIGINT AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP    DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP    DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
    is_deleted BOOL         DEFAULT FALSE COMMENT '是否删除',

    uuid       VARCHAR(127) DEFAULT '' COMMENT 'pixiv uuid',
    pid        BIGINT       NOT NULL COMMENT 'pixiv user id',
    name       VARCHAR(256) NOT NULL COMMENT 'pixiv user name',
    nick_name  VARCHAR(256) DEFAULT '' COMMENT 'pixiv nick name',
    avatar     VARCHAR(512) DEFAULT '' COMMENT 'pixiv user avatar',
    following  INT          DEFAULT 0 COMMENT '关注数量',
    followers  INT          DEFAULT 0 COMMENT '粉丝数量',

    INDEX idx_pid (pid)
) CHARACTER SET utf8;


CREATE TABLE IF NOT EXISTS pixiv_follow
(
    id           BIGINT AUTO_INCREMENT PRIMARY KEY,
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
    is_deleted   BOOL      DEFAULT FALSE COMMENT '是否删除',

    pid          BIGINT    NOT NULL COMMENT 'user id',
    followed_pid BIGINT    NOT NULL COMMENT 'user id',
    followed_at  TIMESTAMP NOT NULL COMMENT '关注时间',

    INDEX idx_followed (pid, followed_pid)
) CHARACTER SET utf8;


CREATE TABLE IF NOT EXISTS pixiv_tag
(
    # 日语tag标签(多语言暂不可靠)
    id         BIGINT AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
    is_deleted BOOL      DEFAULT FALSE COMMENT '是否删除',

    name   TEXT NOT NULL COMMENT 'tage name, jp',

    INDEX idx_tag (name(10))
) CHARACTER SET utf8;


CREATE TABLE IF NOT EXISTS pixiv_illust
(
    id             BIGINT AUTO_INCREMENT PRIMARY KEY,
    created_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
    is_deleted     BOOL      DEFAULT FALSE COMMENT '是否删除',

    # 作者信息
    uid            BIGINT    NOT NULL COMMENT 'pixiv_user.id',
    pid            BIGINT    NOT NULL COMMENT 'pixiv user id',

    # 作品信息
    illust_id      BIGINT    NOT NULL  COMMENT 'pixiv illustId',
    title          TEXT                COMMENT 'pixiv title',
    description    TEXT                COMMENT 'pixiv illust desc',
    view_count     BIGINT    DEFAULT 0 COMMENT '浏览数量',
    like_count     BIGINT    DEFAULT 0 COMMENT '喜欢数',
    bookmark_count BIGINT    DEFAULT 0 COMMENT '收藏欢数',
    create_date    TIMESTAMP NOT NULL COMMENT '创建时间',

    INDEX idx_illust (pid, illust_id)
) CHARACTER SET utf8;


CREATE TABLE IF NOT EXISTS pixiv_manga
(
    id             BIGINT AUTO_INCREMENT PRIMARY KEY,
    created_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
    is_deleted     BOOL      DEFAULT FALSE COMMENT '是否删除',

    # 作者信息
    uid            BIGINT    NOT NULL COMMENT 'pixiv_user.id',
    pid            BIGINT    NOT NULL COMMENT 'pixiv user id',

    # 作品信息
    manga_id       BIGINT    NOT NULL COMMENT 'pixiv mangaID',
    title          TEXT               COMMENT 'pixiv title',
    description    TEXT               COMMENT 'pixiv manga desc',
    page_count     INT       DEFAULT 0 COMMENT '漫画作品页数',
    view_count     BIGINT    DEFAULT 0 COMMENT '浏览数量',
    like_count     BIGINT    DEFAULT 0 COMMENT '喜欢数',
    bookmark_count BIGINT    DEFAULT 0 COMMENT '收藏欢数',
    create_date    TIMESTAMP NOT NULL COMMENT '创建时间',

    INDEX idx_illust (pid, manga_id)
) CHARACTER SET utf8;


CREATE TABLE IF NOT EXISTS pixiv_novel
(
    id               BIGINT AUTO_INCREMENT PRIMARY KEY,
    created_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
    is_deleted       BOOL      DEFAULT FALSE COMMENT '是否删除',

    # 作者信息
    uid              BIGINT    NOT NULL COMMENT 'pixiv_user.id',
    pid              BIGINT    NOT NULL COMMENT 'pixiv user id',

    # 作品信息
    novel_id         BIGINT    NOT NULL COMMENT 'pixiv mangaID',
    title            TEXT               COMMENT 'pixiv title',
    description      TEXT               COMMENT 'pixiv novel caption',
    chapter_count    INT       DEFAULT 0 COMMENT '章节数量, total',
    wordage_count    BIGINT    DEFAULT 0 COMMENT '字数, publishedTotalCharacterCount',
    create_date      TIMESTAMP NOT NULL COMMENT '创建时间',
    last_update_date TIMESTAMP NOT NULL COMMENT '最近更新时间',

    INDEX idx_illust (pid, novel_id)
) CHARACTER SET utf8;


CREATE TABLE IF  NOT EXISTS pixiv_illust_tag (
    id               BIGINT AUTO_INCREMENT PRIMARY KEY,
    created_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
    is_deleted       BOOL      DEFAULT FALSE COMMENT '是否删除',

    illust_id      BIGINT    NOT NULL COMMENT 'pixiv illustId',
    tag_id      BIGINT    NOT NULL COMMENT 'tag id',

    INDEX idx_illust_tag (illust_id, tag_id)
) CHARACTER SET utf8;


CREATE TABLE IF  NOT EXISTS pixiv_manga_tag (
    id               BIGINT AUTO_INCREMENT PRIMARY KEY,
    created_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
    is_deleted       BOOL      DEFAULT FALSE COMMENT '是否删除',

    manga_id      BIGINT    NOT NULL COMMENT 'pixiv mangaId',
    tag_id      BIGINT    NOT NULL COMMENT 'tag id',

    INDEX idx_illust_tag (manga_id, tag_id)
) CHARACTER SET utf8;


CREATE TABLE IF  NOT EXISTS pixiv_novel_tag (
    id               BIGINT AUTO_INCREMENT PRIMARY KEY,
    created_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
    is_deleted       BOOL      DEFAULT FALSE COMMENT '是否删除',

    novel_id      BIGINT    NOT NULL COMMENT 'pixiv novel id',
    tag_id      BIGINT    NOT NULL COMMENT 'tag id',

    INDEX idx_illust_tag (novel_id, tag_id)
) CHARACTER SET utf8;