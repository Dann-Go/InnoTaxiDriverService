CREATE TABLE IF NOT EXISTS drivers (
                       ID SERIAL PRIMARY KEY ,
                       Name TEXT,
                       Phone TEXT,
                       Email TEXT,
                       Password_hash TEXT,
                       Rating FLOAT DEFAULT 0.0,
                       Taxi_type TEXT
);

CREATE INDEX IF NOT EXISTS phone_idx ON drivers (phone);
CREATE INDEX IF NOT EXISTS email_idx ON drivers (email);


