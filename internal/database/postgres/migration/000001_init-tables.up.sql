CREATE TABLE roles (
                      id UUID PRIMARY KEY,
                      role VARCHAR UNIQUE NOT NULL
);

CREATE TABLE users (
                      id UUID PRIMARY KEY,
                      email VARCHAR UNIQUE NOT NULL,
                      password VARCHAR NOT NULL,
                      phone VARCHAR,
                      role_id UUID,
                      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                      updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                      FOREIGN KEY (role_id) REFERENCES Roles(id)
);

CREATE TABLE ranks (
                      id UUID PRIMARY KEY,
                      rank VARCHAR UNIQUE NOT NULL
);

CREATE TABLE flights (
                        id UUID PRIMARY KEY,
                        start_date TIMESTAMP,
                        end_date TIMESTAMP,
                        departure VARCHAR,
                        destination VARCHAR,
                        rank_id UUID,
                        price BIGINT,
                        total_tickets INT,
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        FOREIGN KEY (rank_id) REFERENCES Ranks(id)
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
                        FOREIGN KEY (user_id) REFERENCES Users(id)
);
