create database mydb charset=utf8mb4;
use mydb;
create table person (
    id int primary key auto_increment,
    name varchar(20),
    age int,
    rmb float,
    gender bool,
    birthday date
);

insert into person (name,age,rmb,gender,birthday) values
("tom",25,123,true,19900921),
("jerry",18,3445,true,19880101),
("ann",15,2,false,20010101);