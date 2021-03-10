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
    CONSTRAINT pk_User PRIMARY KEY (userId),
    UNIQUE (login),
    UNIQUE (name)
);

CREATE TABLE Project
(
    projectId           serial  NOT NULL ,
    projectName         text    NOT NULL ,
    projectDescription  text    NULL ,
    CONSTRAINT pk_Project PRIMARY KEY (projectId)
);

CREATE TABLE ProjectUser
(
    userId      int NOT NULL,
    projectId   int NOT NULL,
    CONSTRAINT  pk_ProjectUser  PRIMARY KEY (userId, projectId),
    CONSTRAINT  fk_Project      FOREIGN KEY (projectId) REFERENCES Project (projectId),
    CONSTRAINT  fk_User         FOREIGN KEY (userId)    REFERENCES "User" (userId)

);

CREATE TABLE Issue
(
    issueId             serial      NOT NULL,
    name                text        NOT NULL,
    groupIssueNumber    int         NOT NULL,
    description         text        NOT NULL,
    -- in 'new', 'in progress', 'review', 'testing', 'ready', 'closed'
    status              status_enum NOT NULL,
    releaseVersion      int         NOT NULL,
    creationDate        date        NOT NULL,
    deadline            date        NOT NULL,
    assigneeId          int         NOT NULL,
    authorId            int         NOT NULL,
    projectId           int         NOT NULL,
    CONSTRAINT pk_Bug       PRIMARY KEY (issueId),
    CONSTRAINT fk_Project   FOREIGN KEY (projectId)     REFERENCES Project (projectId),
    CONSTRAINT fk_Author    FOREIGN KEY (authorId)      REFERENCES ProjectUser (userId),
    CONSTRAINT fk_Developer FOREIGN KEY (assigneeId)    REFERENCES ProjectUser (userId)
);

-- CREATE OR REPLACE FUNCTION createExampleProject() RETURNS TRIGGER AS $ExampleProject$
--     BEGIN
--         INSERT INTO Project
--     end;
--     $ExampleProject$ LANGUAGE plpgsql;
--
-- CREATE TRIGGER ExampleProject AFTER INSERT ON "User"
--     FOR EACH ROW EXECUTE PROCEDURE createExampleProject();