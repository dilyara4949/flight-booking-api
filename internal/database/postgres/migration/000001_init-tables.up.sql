CREATE TABLE roles (
                      name VARCHAR PRIMARY KEY
);

CREATE TABLE users (
                      id UUID PRIMARY KEY,
                      email VARCHAR UNIQUE NOT NULL,
                      password VARCHAR NOT NULL,
                      phone VARCHAR,
                      role VARCHAR,
                      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                      updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                      FOREIGN KEY (role) REFERENCES Roles(name)
);

CREATE TABLE ranks (
                      name VARCHAR PRIMARY KEY
);

CREATE TABLE flights (
                        id UUID PRIMARY KEY,
                        start_date TIMESTAMP,
                        end_date TIMESTAMP,
                        departure VARCHAR,
                        destination VARCHAR,
                        rank VARCHAR,
                        price BIGINT,
                        total_tickets INT,
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


insert into roles (name) values ('user'), ('admin');
insert into ranks (name) values ('economy'), ('business'), ('deluxe')
