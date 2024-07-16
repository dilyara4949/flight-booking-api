CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE roles (
                       id UUID PRIMARY KEY,
                       name VARCHAR UNIQUE NOT NULL
);

CREATE TABLE users (
                       id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                       email VARCHAR UNIQUE NOT NULL,
                       password VARCHAR NOT NULL ,
                       phone VARCHAR,
                       role VARCHAR,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       FOREIGN KEY (role) REFERENCES roles(name)
);

CREATE TABLE ranks (
                       id UUID PRIMARY KEY,
                       name VARCHAR UNIQUE NOT NULL
);

CREATE TABLE flights (
                         id UUID PRIMARY KEY,
                         start_date TIMESTAMP NOT NULL ,
                         end_date TIMESTAMP NOT NULL ,
                         departure VARCHAR NOT NULL ,
                         destination VARCHAR NOT NULL ,
                         rank VARCHAR,
                         price BIGINT NOT NULL ,
                         total_tickets INT NOT NULL ,
                         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                         FOREIGN KEY (rank) REFERENCES Ranks(name)
);

CREATE TABLE tickets (
                         id UUID PRIMARY KEY,
                         flight_id UUID,
                         user_id UUID,
                         rank VARCHAR,
                         price INT,
                         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                         FOREIGN KEY (flight_id) REFERENCES Flights(id),
                         FOREIGN KEY (user_id) REFERENCES Users(id),
                         FOREIGN KEY (rank) REFERENCES Ranks(name)
);

INSERT INTO roles (id, name) VALUES (uuid_generate_v4(), 'user'), (uuid_generate_v4(), 'admin');
INSERT INTO ranks (id, name) VALUES (uuid_generate_v4(), 'economy'), (uuid_generate_v4(),'business'), (uuid_generate_v4(), 'deluxe');
