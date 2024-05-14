CREATE TABLE IF NOT EXISTS users (
    user_id uuid default gen_random_uuid() not null constraint users_pk primary key,
    nip varchar(20) not null,
    name varchar(50) not null,
    password varchar(255) default null,
    role varchar(100) not null,
    identity_card_scan_img text,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    deleted_at timestamp default null
);

CREATE UNIQUE INDEX IF NOT EXISTS user_nip_idx ON users (nip, role)
WHERE
    (deleted_at IS NULL);


CREATE TABLE IF NOT EXISTS patients (
    patient_id uuid default gen_random_uuid() not null constraint patients_pk primary key,
    identity_number varchar(16) not null,
    phone varchar(25) not null,
    name varchar(50) not null,
    birth_date timestamp not null,
    gender varchar(50) not null,
    identity_card_scan_img text not null,
    created_at timestamp default now()
);

CREATE UNIQUE INDEX IF NOT EXISTS patients_number_idx ON patients (identity_number);

CREATE TABLE IF NOT EXISTS medical_records (
    id uuid default gen_random_uuid() not null constraint medical_records_pk primary key,
    patient_id uuid not null,
    user_id uuid not null,
    symptoms text not null,
    medications text not null,
    created_at timestamp default now(),
    constraint mr_patient_fk foreign key (patient_id) references patients(patient_id),
    constraint mr_user_fk foreign key (user_id) references users(user_id)
);

CREATE INDEX mr_patient_idx ON medical_records (patient_id);
CREATE INDEX mr_user_idx ON medical_records (user_id);