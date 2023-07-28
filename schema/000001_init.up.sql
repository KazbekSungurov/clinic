CREATE TYPE GENDER AS ENUM ('not decided yet', 'male', 'female');

CREATE TABLE personal_data (
                               id SERIAL PRIMARY KEY,
                               phone CHARACTER VARYING[20],
                               email CHARACTER VARYING[60],
                               address TEXT,
                               polyclinic TEXT
);

CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       fk_personal_data_id INTEGER UNIQUE,
                       first_name CHARACTER VARYING[60] NOT NULL,
                       last_name CHARACTER VARYING[60] NOT NULL,
                       middle_name CHARACTER VARYING[60] NOT NULL,
                       gender GENDER DEFAULT 'not decided yet',
                       birth_date DATE NOT NULL,
                       details TEXT DEFAULT '',
                       created_at TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
                       FOREIGN KEY (fk_personal_data_id) REFERENCES personal_data (id) ON DELETE RESTRICT
);

CREATE TABLE patients (
                          id SERIAL PRIMARY KEY,
                          fk_user_id INTEGER UNIQUE NOT NULL,
                          diagnosis TEXT NOT NULL,
                          FOREIGN KEY (fk_user_id) REFERENCES users(id) ON DELETE RESTRICT
);


CREATE TABLE schedule_details (
                                  id SERIAL PRIMARY KEY,
                                  week_days DATE NOT NULL,
                                  start_at TIME NOT NULL,
                                  end_at TIME NOT NULL,
                                  examination_duration TIME NOT NULL,
                                  break_start TIME NOT NULL,
                                  break_duration TIME NOT NULL,
                                  appointments_limit INTEGER NOT NULL
);

CREATE TABLE schedule_operations (
                                     id SERIAL PRIMARY KEY,
                                     fk_schedule_details_id INTEGER UNIQUE NOT NULL,
                                     FOREIGN KEY (fk_schedule_details_id) REFERENCES schedule_details(id) ON DELETE RESTRICT
);

CREATE TABLE schedule_consults (
                                   id SERIAL PRIMARY KEY,
                                   fk_schedule_details_id INTEGER UNIQUE NOT NULL,
                                   FOREIGN KEY (fk_schedule_details_id) REFERENCES schedule_details(id) ON DELETE RESTRICT
);

CREATE TABLE appointments (
                              id SERIAL PRIMARY KEY,
                              fk_schedule_operations_id INTEGER,
                              fk_schedule_consults_id INTEGER,
                              day DATE NOT NULL,
                              patients INTEGER[],
                              FOREIGN KEY (fk_schedule_operations_id) REFERENCES schedule_operations(id) ON DELETE RESTRICT,
                              FOREIGN KEY (fk_schedule_consults_id) REFERENCES schedule_consults(id) ON DELETE RESTRICT

);

CREATE TABLE departments
(
    id                       SERIAL PRIMARY KEY,
    fk_schedule_consult_id   INTEGER UNIQUE NOT NULL,
    fk_schedule_operation_id INTEGER UNIQUE NOT NULL,
    name                     CHARACTER VARYING[60] NOT NULL,
    FOREIGN KEY (fk_schedule_consult_id) REFERENCES schedule_consults (id) ON DELETE RESTRICT,
    FOREIGN KEY (fk_schedule_operation_id) REFERENCES schedule_operations (id) ON DELETE RESTRICT
);

CREATE TABLE employees (
                           id SERIAL PRIMARY KEY,
                           fk_user_id INTEGER UNIQUE NOT NULL,
                           fk_departments_id INTEGER UNIQUE NOT NULL,
                           FOREIGN KEY (fk_user_id) REFERENCES users(id) ON DELETE RESTRICT,
                           FOREIGN KEY (fk_departments_id) REFERENCES departments(id) ON DELETE RESTRICT
);