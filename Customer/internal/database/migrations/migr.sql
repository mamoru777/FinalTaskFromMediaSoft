-- +migrate Up
create table customer
(
    ID int Primary key,
    Name char(15) not null,
    Password char(20) not null,
    Name char(15) null,
    Last_Name char(20) null,
    Sur_Name char(20) null,
    Email char(20) null,
    Avatar bytea null
);

-- +migare Down
drop table customer;