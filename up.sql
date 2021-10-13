CREATE TABLE IF NOT EXISTS links
(
    id bigserial PRIMARY KEY,
    short_url text,
    long_url text,
);

INSERT INTO links (short_url, long_url) VALUES ("YH-nYjDnR", "http://www.google.com")
INSERT INTO links (short_url, long_url) VALUES ("gHUCaCv7R", "http://www.amazon.co.uk")

