CREATE TABLE Addresses(
    Address_ID SERIAL NOT NULL,
    Customer_ID INT,
    Country VARCHAR(50) NOT NULL,
    City VARCHAR(50) NOT NULL,
    FOREIGN KEY(Customer_ID)
    REFERENCES Customers(Customer_ID)
);
CREATE TABLE Customers(
    Customer_ID SERIAL PRIMARY KEY,
    FirstName VARCHAR(50) NOT NULL,
    LastName VARCHAR(50) NOT NULL,
    UserName VARCHAR(50) NOT NULL,
    Email_id VARCHAR(50) NOT NULL,
    Gender VARCHAR(50) NOT NULL,
    Birthdate DATE NOT NULL,
    Password VARCHAR(50) NOT NULL
);


CREATE TABLE Phones(
    Phone_ID INT ,
    Customer_ID INT,
    Numbers INT [] NOT NULL,
    Code VARCHAR(30) NOT NULL,
      FOREIGN KEY(Customer_ID)
   REFERENCES Customers(Customer_ID)
);

CREATE TABLE Order_c(
    Product_ID SERIAL NOT NULL,
    Name VARCHAR(50) NOT NULL,
    Amount INT NOT NULL,
    Rating INT
);