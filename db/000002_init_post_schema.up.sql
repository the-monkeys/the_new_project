CREATE TABLE IF NOT EXISTS the_monkeys_post
(
    id SERIAL,
    uuid text NOT NULL,
    title text NOT NULL,
    html_content text NOT NULL,
    raw_content text NOT NULL,
    author_name text NOT NULL,
    author_id bigint NOT NULL,
    published boolean,
    tags text [],
    create_time text,
    update_time text,
    can_edit boolean,
    content_ownership text,
    folder_path text,
    

    CONSTRAINT blog_pkey PRIMARY KEY (id),

    CONSTRAINT fk_author_id FOREIGN KEY (author_id)
        REFERENCES the_monkeys_user(id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION 
);