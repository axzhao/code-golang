create database test;

create table if not exists me (
    id bigserial NOT NULL,
    country varchar(128)
);

insert into me (country) values ('a'), ('b'), ('c'), ('d');

alter table me add column arabic_name varchar(128);

update me as t set
    arabic_name = tmp.arabic_name
    from (values ('a', '1a'),('b', '1b'),('c', '1c'),('d', '1d')) as tmp (country, arabic_name)
where tmp.country = t.country;

select * from me;

drop database test;
