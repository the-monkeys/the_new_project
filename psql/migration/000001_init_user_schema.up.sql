CREATE TABLE IF NOT EXISTS the_monkeys_user
(
    id SERIAL,
    unique_id text NOT NULL,
    first_name text NOT NULL,
    last_name text NOT NULL,
    email text NOT NULL,
    password text NOT NULL,
    profile_pic bytea,
    create_time text,
    update_time text,
    is_active boolean,
    role integer,
    last_login text,
    country_code text DEFAULT 'none',
    mobile_no text DEFAULT 'none',
    about text DEFAULT 'none',
    instagram text DEFAULT 'none',
    twitter text DEFAULT 'none',

    email_verified boolean,
    email_verification_token text,
    email_verification_timeout timestamp with time zone,
    mobile_verified boolean,
    mobile_verification_token text,
    mobile_verification_timeout timestamp with time zone,

    CONSTRAINT users_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS pw_reset
(
    id SERIAL ,
    user_id bigint NOT NULL,
    email text,
    recovery_hash text ,
    time_out timestamp with time zone,
    -- TODO: rename last_password_reset to password_reset_time
    last_password_reset timestamp with time zone,

    CONSTRAINT password_resets_pkey PRIMARY KEY (id),

    CONSTRAINT fk_user_id FOREIGN KEY (user_id)
        REFERENCES the_monkeys_user(id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION 
);

