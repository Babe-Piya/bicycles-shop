CREATE TABLE bicycles (
                          id serial PRIMARY KEY,
                          brand VARCHAR ( 100 ) NOT NULL,
                          model VARCHAR ( 100 ) NOT NULL,
                          description TEXT,
                          price NUMERIC (10, 2)  NOT NULL,
                          created_at TIMESTAMP,
                          updated_at TIMESTAMP
);

INSERT INTO public.bicycles (id, brand, model, description, price, created_at, updated_at)
VALUES (DEFAULT, 'Ka rang', 'Ez', 'Red fast forever', 15000.00, '2022-07-15 00:15:12.000000',
        '2022-07-15 00:15:18.000000');