create table if not exists food(
    id serial primary key,
    name varchar(100) not null,
    time int not null
);