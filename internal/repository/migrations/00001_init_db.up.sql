create table Users
(
    student  TEXT not null primary key,
    task_id  INTEGER,
    password TEXT not null,
    pending  INTEGER,
    msg      TEXT
);
create table PullRequests
(
    id      INTEGER PRIMARY KEY AUTOINCREMENT,
    link    TEXT not null,
    student TEXT references Users
);