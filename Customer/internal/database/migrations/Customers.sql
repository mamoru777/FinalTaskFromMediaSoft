-- +migrate Up
create table offices
(
    id uuid default uuid_generate_v4() primary key,
    name text not null,
    address text not null,
    created_at timestamp not null
);

-- +migrate Down
alter table customers drop constraint fk_offices_customer;

-- +migrate Down
drop table offices;

-- +migrate Up
create table customers
(
    id uuid default uuid_generate_v4() primary key,
    name text not null,
    office_id uuid not null,
    office_name text not null,
    created_at timestamp not null,
    foreign key (office_id) references offices (id) on delete set null
);

-- +migrate Down
drop table customers;