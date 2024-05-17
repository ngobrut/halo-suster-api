DROP INDEX IF EXISTS user_nip_idx;
DROP INDEX IF EXISTS user_name_idx;
DROP INDEX IF EXISTS user_role_idx;

DROP INDEX IF EXISTS patient_number_idx;
DROP INDEX IF EXISTS patient_phone_idx;
DROP INDEX IF EXISTS patient_name_idx;

DROP INDEX IF EXISTS mr_patient_idx;
DROP INDEX IF EXISTS mr_user_idx;

DROP TABLE IF EXISTS medical_records;
DROP TABLE IF EXISTS patients;
DROP TABLE IF EXISTS users;