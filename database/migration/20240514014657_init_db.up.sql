CREATE TABLE IF NOT EXISTS users (
    user_id uuid default gen_random_uuid() not null constraint users_pk primary key,
    nip varchar(20) not null,
    name varchar(50) not null,
    password varchar(255) default null,
    role varchar(100) not null,
    identity_card_scan_img varchar(255),
    created_at timestamp default now(),
    updated_at timestamp default now(),
    deleted_at timestamp default null
);

CREATE UNIQUE INDEX IF NOT EXISTS user_nip_idx ON users (nip, role)
WHERE
    (deleted_at IS NULL);