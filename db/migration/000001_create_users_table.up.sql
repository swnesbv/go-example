
CREATE TABLE users (
   user_id      serial PRIMARY KEY,
   username     VARCHAR (50) UNIQUE,
   email        TEXT UNIQUE,
   password     TEXT,
   img          TEXT,
   created_at   TIMESTAMPTZ,
   updated_at   TIMESTAMPTZ
);
CREATE TABLE article (
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    img         TEXT,
    owner       INTEGER,
    completed   BOOLEAN DEFAULT false,
    created_at  TIMESTAMPTZ,
    updated_at  TIMESTAMPTZ,
    FOREIGN KEY (owner) REFERENCES users(user_id) ON DELETE CASCADE
);
CREATE TABLE msguser (
    id          SERIAL PRIMARY KEY,
    coming      TEXT,
    img         TEXT,
    owner       INTEGER,
    to_user     INTEGER,
    completed   BOOLEAN DEFAULT false,
    created_at  TIMESTAMPTZ,
    updated_at  TIMESTAMPTZ,
    FOREIGN KEY (owner) REFERENCES users(user_id),
    FOREIGN KEY (to_user) REFERENCES users(user_id)
);
CREATE TABLE groups (
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    img         TEXT,
    owner       INTEGER,
    completed   BOOLEAN DEFAULT false,
    created_at  TIMESTAMPTZ,
    updated_at  TIMESTAMPTZ,
    FOREIGN KEY (owner) REFERENCES users(user_id) ON DELETE CASCADE
);
CREATE TABLE msggroups (
    id          SERIAL PRIMARY KEY,
    coming      TEXT,
    img         TEXT,
    owner       INTEGER,
    to_group    INTEGER,
    completed   BOOLEAN DEFAULT false,
    created_at  TIMESTAMPTZ,
    updated_at  TIMESTAMPTZ,
    FOREIGN KEY (owner) REFERENCES users(user_id),
    FOREIGN KEY (to_group) REFERENCES groups(id) ON DELETE CASCADE
);
CREATE TABLE subscription (
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    owner       INTEGER,
    to_user     INTEGER,
    to_group    INTEGER,
    completed   BOOLEAN DEFAULT false,
    created_at  TIMESTAMPTZ,
    updated_at  TIMESTAMPTZ,
    FOREIGN KEY (owner) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (to_user) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (to_group) REFERENCES groups(id) ON DELETE CASCADE
);
CREATE TABLE provision_d (
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(255),
    description TEXT,
    owner       INTEGER,
    st_date     Date,
    en_date     Date,
    s_dates     Date[],
    e_dates     Date[],
    dates       Date[],
    completed   BOOLEAN DEFAULT false,
    created_at  TIMESTAMPTZ,
    updated_at  TIMESTAMPTZ,
    FOREIGN KEY (owner) REFERENCES users(user_id) ON DELETE CASCADE
);
CREATE TABLE provision_h (
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(255),
    description TEXT,
    owner       INTEGER,
    st_hour     TIMESTAMP,
    en_hour     TIMESTAMP,
    s_hours     TIMESTAMP[],
    e_hours     TIMESTAMP[],
    hours       TIMESTAMP[],
    completed   BOOLEAN DEFAULT false,
    created_at  TIMESTAMPTZ,
    updated_at  TIMESTAMPTZ,
    FOREIGN KEY (owner) REFERENCES users(user_id) ON DELETE CASCADE
);
CREATE TABLE booking (
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    owner       INTEGER,
    to_prv_d    INTEGER,
    to_prv_h    INTEGER,
    st_date     Date,
    en_date     Date,
    st_hour     TIMESTAMP,
    en_hour     TIMESTAMP,
    completed   BOOLEAN DEFAULT false,
    created_at  TIMESTAMPTZ,
    updated_at  TIMESTAMPTZ,
    FOREIGN KEY (owner) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (to_prv_d) REFERENCES provision_d(id) ON DELETE CASCADE,
    FOREIGN KEY (to_prv_h) REFERENCES provision_h(id) ON DELETE CASCADE
);
CREATE TABLE schedule (
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(255),
    description TEXT,
    owner       INTEGER,
    st_hour     TIMESTAMP,
    en_hour     TIMESTAMP,
    hours       TIMESTAMP[],
    occupied    TIMESTAMP[],
    completed   BOOLEAN DEFAULT false,
    created_at  TIMESTAMPTZ,
    updated_at  TIMESTAMPTZ,
    FOREIGN KEY (owner) REFERENCES users(user_id) ON DELETE CASCADE
);
CREATE TABLE recording (
    id          SERIAL PRIMARY KEY,
    owner       INTEGER,
    to_schedule INTEGER,
    record_d    Date,
    record_h    TIMESTAMP,
    completed   BOOLEAN DEFAULT false,
    created_at  TIMESTAMPTZ,
    updated_at  TIMESTAMPTZ,
    FOREIGN KEY (owner) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (to_schedule) REFERENCES schedule(id) ON DELETE CASCADE
);
CREATE TABLE collection (
    id            SERIAL PRIMARY KEY,
    collection_id VARCHAR(8),
    owner         INTEGER,
    pfile         TEXT[],
    completed     BOOLEAN DEFAULT false,
    created_at    TIMESTAMPTZ,
    updated_at    TIMESTAMPTZ,
    FOREIGN KEY (owner) REFERENCES users(user_id) ON DELETE CASCADE
);
CREATE TABLE slider (
    id          SERIAL PRIMARY KEY,
    collection_id VARCHAR(8),
    title       VARCHAR(255),
    description TEXT,
    owner       INTEGER,
    to_art      INTEGER,
    to_sch      INTEGER,
    to_prv_d    INTEGER,
    to_prv_h    INTEGER,
    pfile       TEXT[],
    completed   BOOLEAN DEFAULT false,
    created_at  TIMESTAMPTZ,
    updated_at  TIMESTAMPTZ,
    FOREIGN KEY (owner) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (to_art) REFERENCES article(id) ON DELETE CASCADE,
    FOREIGN KEY (to_sch) REFERENCES schedule(id) ON DELETE CASCADE,
    FOREIGN KEY (to_prv_d) REFERENCES provision_d(id) ON DELETE CASCADE,
    FOREIGN KEY (to_prv_h) REFERENCES provision_h(id) ON DELETE CASCADE
);
