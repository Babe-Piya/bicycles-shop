CREATE TABLE bicycles (
                          id serial PRIMARY KEY,
                          brand VARCHAR ( 100 ) NOT NULL,
                          model VARCHAR ( 100 ) NOT NULL,
                          description TEXT,
                          price NUMERIC (10, 2)  NOT NULL,
                          status int default 1 not null,
                          url TEXT ,
                          created_at TIMESTAMP,
                          updated_at TIMESTAMP
);

INSERT INTO public.bicycles (id, brand, model, description, price,status,url, created_at, updated_at)
VALUES (DEFAULT, 'Ka rang', 'Ez', 'Red fast forever', 15000.00,1,'https://drive.google.com/file/d/1fBYqzYQiXe2xwJ9aK5nKrXcd4VJ41JK6/view?usp=sharing', '2022-07-15 00:15:12.000000',
        '2022-07-15 00:15:18.000000');

create table buyers
(
    id         serial
        constraint buyers_pk
            primary key,
    bicycle_id int          not null
        constraint bicycle_id
            references bicycles,
    name       varchar(255) not null,
    address    TEXT         not null,
    tel        varchar(20),
    slip text,
    created_at TIMESTAMP
);
