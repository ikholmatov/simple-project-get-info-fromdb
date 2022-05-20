CREATE TABLE customers(
    Customer_ID uuid PRIMARY KEY,
    FirstName VARCHAR(50) NOT NULL,
    LastName VARCHAR(50) NOT NULL,
    UserName VARCHAR(50) NOT NULL,
    Email_id VARCHAR(50) NOT NULL,
    Gender VARCHAR(50) NOT NULL,
    Birthdate DATE NOT NULL,
    Password VARCHAR(50) NOT NULL
);

CREATE TABLE addresses(
    Address_ID uuid NOT NULL,
    Customer_ID uuid,
    Country VARCHAR(50) NOT NULL,
    City VARCHAR(50) NOT NULL,
    FOREIGN KEY(Customer_ID)
    REFERENCES customers(Customer_ID)
);

CREATE TABLE phones(
    Phone_ID uuid PRIMARY KEY,
    Customer_ID uuid,
    Numbers INT [] NOT NULL,
    Code VARCHAR(30) NOT NULL,
      FOREIGN KEY(Customer_ID)
   REFERENCES customers(Customer_ID)
);

CREATE TABLE products(
    Product_ID uuid PRIMARY KEY,
    Name VARCHAR(50) NOT NULL,
    Amount INT NOT NULL,
    Rating INT
);