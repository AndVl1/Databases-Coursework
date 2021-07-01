\c bugtrackcourse;

CREATE TYPE status_enum AS ENUM ('new', 'in progress', 'review', 'testing', 'ready', 'closed');

CREATE TYPE label_enum AS ENUM ('DB', 'Interface', 'Docs');

CREATE TYPE role_enum AS ENUM ('Admin', 'PM', 'Developer', 'User');

CREATE TABLE "User"
(
    userId   serial    NOT NULL,
    login    text      NOT NULL,
    password text      NOT NULL,
    name     text      NOT NULL,
    CONSTRAINT pk_User PRIMARY KEY (userId),
    UNIQUE (login),
    UNIQUE (name)
);

CREATE TABLE Project
(
    projectId           serial  NOT NULL ,
    projectName         text    NOT NULL ,
    projectDescription  text    NULL ,
    issuesCount         int     NOT NULL DEFAULT 0 ,
    CONSTRAINT pk_Project PRIMARY KEY (projectId)
);

CREATE TABLE ProjectUser
(
    userId      int NOT NULL,
    projectId   int NOT NULL,
    CONSTRAINT  pk_ProjectUser  PRIMARY KEY (userId, projectId),
    CONSTRAINT  fk_Project      FOREIGN KEY (projectId) REFERENCES Project (projectId)
        ON UPDATE CASCADE ON DELETE CASCADE ,
    CONSTRAINT  fk_User         FOREIGN KEY (userId)    REFERENCES "User" (userId)
        ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE VIEW ProjectUsersView AS
    SELECT U.userId, U.login, U.name, P.projectId,
           P.projectName, P.projectDescription, P.issuesCount FROM "User" as U
        LEFT JOIN ProjectUser PU on U.userId = PU.userId
        LEFT JOIN Project P on PU.projectId = P.projectId;

CREATE TABLE Status
(
    statusId    serial      NOT NULL ,
    statusName  status_enum NOT NULL ,
    CONSTRAINT pk_Status PRIMARY KEY (statusId)
);

INSERT INTO Status (statusName) VALUES
('new'),
('in progress'),
('review'),
('testing'),
('ready'),
('closed');

CREATE TABLE Label
(
    labelId     serial      NOT NULL ,
    labelName   label_enum  NOT NULL ,
    CONSTRAINT pk_Label PRIMARY KEY (labelId)
);

INSERT INTO Label (labelName) VALUES ('DB'),
                                     ('Interface'),
                                     ('Docs');

CREATE TABLE Issue
(
    issueId             serial      NOT NULL,
    name                text        NOT NULL,
    projectIssueNumber  int         NOT NULL,
    description         text        NOT NULL,
    creationDate        bigint      NOT NULL,
    deadline            bigint      NULL DEFAULT NULL,
    assigneeId          int         NULL DEFAULT NULL,
    authorId            int         NOT NULL,
    projectId           int         NOT NULL,
    releaseVersion      text        NOT NULL,
    statusId            int         NOT NULL,
    -- in 'new', 'in progress', 'review', 'testing', 'ready', 'closed'
    labelId             int         NOT NULL,
    CONSTRAINT pk_Bug       PRIMARY KEY (issueId),
    CONSTRAINT fk_Project   FOREIGN KEY (projectId)     REFERENCES Project (projectId) ON DELETE CASCADE ,
    CONSTRAINT fk_Author    FOREIGN KEY (authorId)      REFERENCES "User" (userId) ON DELETE RESTRICT ,
    CONSTRAINT fk_Developer FOREIGN KEY (assigneeId)    REFERENCES "User" (userId) ON DELETE SET NULL ,
    -- wanted to constraint them to ProjectUser, but is didn't work
    -- "there is no unique constraint matching given keys for referenced table "projectuser""
    -- in ProjectUser must be unique combination, but not unique fields
    CONSTRAINT fk_Status    FOREIGN KEY (statusId)      REFERENCES Status (statusId) ON DELETE RESTRICT,
    CONSTRAINT fk_Label     FOREIGN KEY (labelId)       REFERENCES Label (labelId) ON DELETE RESTRICT
);

CREATE TABLE Comment
(
    commentId   serial  NOT NULL ,
    commentText text    NOT NULL ,
    commentDate bigint  NOT NULL ,
    authorId    int     NOT NULL ,
    issueId     int     NOT NULL ,
    CONSTRAINT fk_Issue     FOREIGN KEY (issueId)   REFERENCES Issue(issueId) ON DELETE CASCADE ,
    CONSTRAINT fk_User      FOREIGN KEY (authorId)  REFERENCES "User"(userId) ON DELETE CASCADE
-- User можно сделать SET NULL, но заморачиваться не особо хочется, лучше просто всё удалить
);

CREATE VIEW CommentView AS
    SELECT C.commentId, C.commentText, C.commentDate, C.issueId, U.name, U.userId, U.login
    FROM Comment as C
        INNER JOIN "User" U on U.userId = C.authorId;
