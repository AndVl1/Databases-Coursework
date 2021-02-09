\c bugtrackcourse;
--
-- CREATE TABLE students
-- (
--     id bigint NOT NULL,
--     name text COLLATE pg_catalog."default",
--     CONSTRAINT student_pkey PRIMARY KEY (id)
--
-- );
--
-- INSERT INTO students(id, name) VALUES
-- (1, 'A'),
-- (2, 'B'),
-- (3, 'C');

CREATE TYPE status_enum AS ENUM ('new', 'in progress', 'review', 'testing', 'ready', 'closed');

CREATE TYPE label_enum AS ENUM ('DB', 'Interface', 'Docs');

CREATE TYPE role_enum AS ENUM ('Admin', 'PM', 'Developer', 'User');

CREATE TABLE "User"
(
    userId   serial    NOT NULL,
    login    text      NOT NULL,
    password text      NOT NULL,
    name     text      NOT NULL,
    -- in 'Admin', 'PM', 'Developer', 'User'
    role     role_enum NOT NULL,
    CONSTRAINT pk_User PRIMARY KEY (userId)
);

INSERT INTO "User" (login, password, name, role)
VALUES ('abc', 'anc', 'anc', 'PM');

CREATE TABLE Bug
(
    bugId              serial      NOT NULL,
    name               text        NOT NULL,
    description        text        NOT NULL,
    -- in 'new', 'in progress', 'review', 'testing', 'ready', 'closed'
    status             status_enum NOT NULL,
    authorId           int         NOT NULL,
    attachmentId       int         NOT NULL,
    commentId          int         NOT NULL,
    givenToDeveloperId int         NOT NULL,
    CONSTRAINT pk_Bug PRIMARY KEY (
                                   bugId
        )
);

CREATE TABLE Label
(
    labelId   serial     NOT NULL,
    -- 'DB', 'Interface', 'Docs'
    labelName label_enum NOT NULL,
    CONSTRAINT pk_Label PRIMARY KEY (
                                     labelId
        )
);

CREATE TABLE Comment
(
    commentId      serial NOT NULL,
    text           text   NOT NULL,
    authorId       int    NOT NULL,
    attachmentId   int    NOT NULL,
    CONSTRAINT pk_Comment PRIMARY KEY (commentId)
);

CREATE TABLE Attachment
(
    attachmentId   serial NOT NULL,
    authorId       int    NOT NULL,
    attachmentPath text   NOT NULL,
    CONSTRAINT pk_Attachment PRIMARY KEY (attachmentId)
);

CREATE TABLE Bug_Label_Intersection
(
    bugId   int NOT NULL,
    labelId int NOT NULL,
    CONSTRAINT pk_Bug_Label_Intersection PRIMARY KEY (bugId, labelId)
);

ALTER TABLE Bug
    ADD CONSTRAINT fk_Bug_AuthorId FOREIGN KEY (authorId)
        REFERENCES "User" (userId);

ALTER TABLE Bug
    ADD CONSTRAINT fk_Bug_AttachmentId FOREIGN KEY (attachmentId)
        REFERENCES Attachment (attachmentId);

ALTER TABLE Bug
    ADD CONSTRAINT fk_Bug_CommentId FOREIGN KEY (commentId)
        REFERENCES Comment (commentId);

ALTER TABLE Bug
    ADD CONSTRAINT fk_Bug_GivenToDeveloperId FOREIGN KEY (givenToDeveloperId)
        REFERENCES "User" (userId);

ALTER TABLE Comment
    ADD CONSTRAINT fk_Comment_AuthorId FOREIGN KEY (authorId)
        REFERENCES "User" (userId);

ALTER TABLE Comment
    ADD CONSTRAINT fk_Comment_AttachmentId FOREIGN KEY (attachmentId)
        REFERENCES Attachment (attachmentId);

ALTER TABLE Attachment
    ADD CONSTRAINT fk_Attachment_AuthorId FOREIGN KEY (authorId)
        REFERENCES "User" (userId);

ALTER TABLE Bug_Label_Intersection
    ADD CONSTRAINT fk_Bug_Label_Intersection_BugId FOREIGN KEY (bugId)
        REFERENCES Bug (bugId);

ALTER TABLE Bug_Label_Intersection
    ADD CONSTRAINT fk_Bug_Label_Intersection_LabelId FOREIGN KEY (labelId)
        REFERENCES Label (labelId);

