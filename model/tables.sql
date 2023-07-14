CREATE TABLE users
(
    uid        INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    name       VARCHAR(32) COLLATE utf8_general_ci  NOT NULL,
    password   VARCHAR(64) COLLATE utf8_general_ci  NOT NULL,
    mail       VARCHAR(200) COLLATE utf8_general_ci NOT NULL,
    url        VARCHAR(200) COLLATE utf8_general_ci          DEFAULT NULL,
    screenName VARCHAR(32) COLLATE utf8_general_ci           DEFAULT NULL,
    created    INT(10) UNSIGNED NOT NULL DEFAULT '0',
    activated  INT(10) UNSIGNED NOT NULL DEFAULT '0',
    logged     INT(10) UNSIGNED NOT NULL DEFAULT '0',
    userGroup  VARCHAR(16) COLLATE utf8_general_ci  NOT NULL DEFAULT 'visitor',
    authCode   VARCHAR(64) COLLATE utf8_general_ci           DEFAULT NULL,
    PRIMARY KEY (uid),
    KEY        idx_name ( name),
    KEY        idx_mail (mail)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

CREATE TABLE relationships
(
    cid INT(10) UNSIGNED NOT NULL,
    mid INT(10) UNSIGNED NOT NULL,
    PRIMARY KEY (cid),
    UNIQUE KEY idx_cid_mid (cid, mid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

CREATE TABLE options
(
    name  VARCHAR(32) COLLATE utf8_general_ci,
    user  INT(10) UNSIGNED NOT NULL DEFAULT 0 PRIMARY KEY,
    value TEXT COLLATE utf8_general_ci
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

CREATE TABLE metas
(
    mid         INT(10) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name        VARCHAR(200) COLLATE utf8_general_ci DEFAULT NULL,
    slug        VARCHAR(200) COLLATE utf8_general_ci DEFAULT NULL,
    type        VARCHAR(32) COLLATE utf8_general_ci NOT NULL,
    description VARCHAR(200) COLLATE utf8_general_ci DEFAULT NULL,
    count       INT(10) UNSIGNED NOT NULL DEFAULT 0,
    orders      INT(10) UNSIGNED NOT NULL DEFAULT 0,
    parent      INT(10) UNSIGNED NOT NULL DEFAULT 0,
    INDEX       idx_slug (slug)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

CREATE TABLE fields
(
    cid         INT(10) UNSIGNED NOT NULL,
    name        VARCHAR(200) COLLATE utf8_general_ci NOT NULL,
    type        VARCHAR(8) COLLATE utf8_general_ci            DEFAULT 'str',
    str_value   TEXT COLLATE utf8_general_ci                  DEFAULT NULL,
    int_value   INT(10) NOT NULL DEFAULT 0,
    float_value FLOAT                                NOT NULL DEFAULT 0,
    PRIMARY KEY (cid, name),
    INDEX       idx_int (int_value),
    INDEX       idx_float (float_value)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

CREATE TABLE comments
(
    `coid`     int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `cid`      int(10) UNSIGNED NOT NULL DEFAULT '0',
    `created`  int(10) UNSIGNED NOT NULL DEFAULT '0',
    `author`   varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci         DEFAULT NULL,
    `authorId` int(10) UNSIGNED NOT NULL DEFAULT '0',
    `ownerId`  int(10) UNSIGNED NOT NULL DEFAULT '0',
    `mail`     varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci         DEFAULT NULL,
    `url`      varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci         DEFAULT NULL,
    `ip`       varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci          DEFAULT NULL,
    `agent`    varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci         DEFAULT NULL,
    `text`     text CHARACTER SET utf8 COLLATE utf8_general_ci,
    `type`     varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'comment',
    `status`   varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'approved',
    `parent`   int(10) UNSIGNED NOT NULL DEFAULT '0',
    PRIMARY KEY (`coid`),
    KEY        `cid` (`cid`),
    KEY        `created` (`created`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


CREATE TABLE contents
(
    cid          int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    title        varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci         DEFAULT NULL,
    slug         varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci         DEFAULT NULL,
    created      int(10) UNSIGNED NOT NULL DEFAULT '0',
    modified     int(10) UNSIGNED NOT NULL DEFAULT '0',
    text         longtext CHARACTER SET utf8 COLLATE utf8_general_ci,
    orders        int(10) UNSIGNED NOT NULL DEFAULT '0',
    authorId     int(10) UNSIGNED NOT NULL DEFAULT '0',
    template     varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci          DEFAULT NULL,
    type         varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'post',
    status       varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'publish',
    password     varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci          DEFAULT NULL,
    commentsNum  int(10) UNSIGNED NOT NULL DEFAULT '0',
    allowComment char(1) CHARACTER SET utf8 COLLATE utf8_general_ci     NOT NULL DEFAULT '0',
    allowPing    char(1) CHARACTER SET utf8 COLLATE utf8_general_ci     NOT NULL DEFAULT '0',
    PRIMARY KEY (cid),
    KEY          slug (slug),
    KEY          created (created)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;