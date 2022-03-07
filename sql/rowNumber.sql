create database test;

create table if not exists me (
    id bigserial NOT NULL,
    country varchar(128)
);

insert into me (country) values ('a'), ('b'), ('c'), ('a');

select * from me;

select * FROM me WHERE id IN (SELECT id FROM (SELECT id, ROW_NUMBER() OVER w AS rnum FROM me WINDOW w AS (PARTITION BY country ORDER BY id desc)) t WHERE t.rnum > 1);

drop database test;