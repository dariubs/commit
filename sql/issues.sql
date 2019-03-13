CREATE TABLE issues(
    id serial primary key,

    title text,
    context text,

    url text,
    repo text,
    issue_id int,

    user_id int,
    username text,
    state text,
    locked boolean,
    assignee text,
    
    user text,
    repo text,

    lang text,

    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);