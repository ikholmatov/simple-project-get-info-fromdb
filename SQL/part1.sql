CREATE TABLE customers(
    CustomerID uuid PRIMARY KEY,
    FirstName VARCHAR(50) NOT NULL,
    LastName VARCHAR(50) NOT NULL,
    UserName VARCHAR(50) NOT NULL,
    Email VARCHAR(50) NOT NULL,
    Gender VARCHAR(50) NOT NULL,
    BirthDate DATE NOT NULL,
    Password VARCHAR(50) NOT NULL,
    Status BOOLEAN NOT NULL
);

CREATE TABLE phones(
    ID uuid PRIMARY KEY,
    CustomerID uuid,
    Numbers BIGINT [] NOT NULL,
    Code VARCHAR(30) NOT NULL,
    FOREIGN KEY(CustomerID)
    REFERENCES customers(CustomerID)
    ON DELETE CASCADE
);
CREATE TABLE addresses(
    ID uuid NOT NULL PRIMARY KEY,
    CustomerID uuid,
    Country VARCHAR(50) NOT NULL,
    City VARCHAR(50) NOT NULL,
    District VARCHAR(50) NOT NULL ,
    PostalCode BIGINT [] NOT NULL,
    FOREIGN KEY(CustomerID)
    REFERENCES customers(CustomerID)
    ON DELETE CASCADE
);

CREATE TABLE products(
    ID uuid PRIMARY KEY,
    CustomerID uuid,
    Name VARCHAR(50) NOT NULL,
    Cost BIGINT NOT NULL,
    OrderNumber BIGINT NOT NULL,
    Amount BIGINT NOT NULL,
    Currency VARCHAR(50) NOT NULL,
    Rating BIGINT NOT NULL,
    FOREIGN KEY(CustomerID)
    REFERENCES customers(CustomerID)
    ON DELETE CASCADE
);
CREATE TABLE types(
    ID uuid NOT NULL ,
    ProductID uuid NOT NULL ,
    Name TEXT [] NOT NULL,
    FOREIGN KEY(ProductID)
    REFERENCES products(ID)
    ON DELETE CASCADE
);