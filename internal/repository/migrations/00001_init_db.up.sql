create table Users
(
    student  TEXT not null primary key,
    task_id  INTEGER,
    password TEXT not null,
    pending  INTEGER,
    msg      INTEGER
);
create table PullRequests
(
    id      SERIAL primary key,
    link    TEXT not null,
    student TEXT references Users
);