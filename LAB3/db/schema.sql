# noinspection SqlNoDataSourceInspectionForFile

-- Stores identities and links between them
CREATE TABLE IF NOT EXISTS `events`
(
    `id`          int                                             NOT NULL,
    `entity`      ENUM ('TASK', 'TASK_COMMENT', 'USER')           NOT NULL,
    `payload`     VARBINARY(255)                                  NOT NULL
)
    DEFAULT CHARSET = latin1
    DEFAULT COLLATE = latin1_general_ci
;