IF NOT EXISTS (SELECT * FROM sysobjects WHERE name='suppliers' and xtype='U')
BEGIN
    CREATE TABLE suppliers (
        id BIGINT IDENTITY(1,1) PRIMARY KEY,
        name NVARCHAR(255) NOT NULL
    );
END;
