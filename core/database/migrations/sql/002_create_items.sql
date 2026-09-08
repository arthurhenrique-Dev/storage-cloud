IF NOT EXISTS (SELECT * FROM sysobjects WHERE name='items' and xtype='U')
BEGIN
    CREATE TABLE items (
        id BIGINT IDENTITY(1,1) PRIMARY KEY,
        name NVARCHAR(255) NOT NULL,
        description NVARCHAR(MAX),
        price REAL NOT NULL,
        supplier_id BIGINT
    );
END;
