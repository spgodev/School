CREATE TABLE IF NOT EXISTS students (
                                        id BIGSERIAL PRIMARY KEY,
                                        name TEXT NOT NULL,
                                        age INT NOT NULL,
                                        gender TEXT NOT NULL CHECK (gender IN ('Male', 'Female')),
                                        height INT NOT NULL
);