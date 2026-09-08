IF NOT EXISTS (SELECT * FROM sysobjects WHERE name='users' and xtype='U')
BEGIN
    CREATE TABLE users (
        id BIGINT IDENTITY(1,1) PRIMARY KEY,
        email NVARCHAR(255) NOT NULL UNIQUE,
        first_name NVARCHAR(100),
        last_name NVARCHAR(100),
        password_hash NVARCHAR(MAX)
    );
END;
