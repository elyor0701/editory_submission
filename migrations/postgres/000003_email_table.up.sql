create table email_verification (
    email varchar(255) not null,
    token varchar not null,
    sent bool default false,
    expires_at timestamp default CURRENT_TIMESTAMP + INTERVAL '1 day',
    created_at timestamp default CURRENT_TIMESTAMP,
    UNIQUE(email, token)
);