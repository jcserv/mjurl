CREATE TABLE url (
    id SERIAL PRIMARY KEY,
    short_url TEXT NOT NULL,
    long_url TEXT NOT NULL
);

CREATE INDEX url_short_url_idx ON url (short_url);
CREATE INDEX url_long_url_idx ON url(long_url);