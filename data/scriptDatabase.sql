CREATE DATABASE inventory;
USE inventory;

CREATE TABLE Users(
    Id INT  NOT NULL auto_increment,
    Email VARCHAR(255) NOT NULL,
    Name VARCHAR(255) NOT NULL,
    Password VARCHAR(255) NOT NULL,
    PRIMARY KEY (Id)
);

CREATE TABLE Products(
    Id INT  NOT NULL auto_increment,
    Name VARCHAR(255) NOT NULL,
    Description VARCHAR(255) NOT NULL,
    Price FLOAT NOT NULL,
    CreatedBy INT NOT NULL,
    PRIMARY KEY (Id),
    FOREIGN KEY (CreatedBy) REFERENCES Users(Id)
);

CREATE TABLE Roles (
    Id INT  NOT NULL auto_increment,
    Name VARCHAR(255) NOT NULL,
    PRIMARY KEY (Id)
);

CREATE TABLE User_Roles(
    Id INT  NOT NULL auto_increment,
    UserId INT NOT NULL,
    RoleId INT NOT NULL,
    PRIMARY KEY (Id),
    FOREIGN KEY (UserId) REFERENCES Users(Id),
    FOREIGN KEY (RoleId) REFERENCES Roles(Id)
);

insert into Roles (Id, name) values (1, 'ROOT');
insert into Roles (Id, name) values (2, 'ADMIN');
insert into Roles (Id, name) values (3, 'USER-APP');