CREATE TABLE bicycles (
                          id serial PRIMARY KEY,
                          brand VARCHAR ( 100 ) NOT NULL,
                          model VARCHAR ( 100 ) NOT NULL,
                          description TEXT,
                          price NUMERIC (10, 2)  NOT NULL,
                          status int default 1 not null,
                          created_at TIMESTAMP,
                          updated_at TIMESTAMP
);

INSERT INTO public.bicycles (id, brand, model, description, price, created_at, updated_at)
VALUES (DEFAULT, 'Ka rang', 'Ez', 'Red fast forever', 15000.00, '2022-07-15 00:15:12.000000',
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
    address    text         not null,
    tel        varchar(20),
    created_at TIMESTAMP
);
