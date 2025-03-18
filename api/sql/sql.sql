create database if not exists db_golang;

use db_golang;

drop table if exists usuarios;

create table usuarios (
                          id INT auto_increment primary key,
                          nome VARCHAR(60) not null,
                          nick VARCHAR(20) not null unique,
                          email VARCHAR(60) unique not null,
                          senha varchar(60) not null,
                          criadoEm timestamp default CURRENT_TIMESTAMP()
);

select
    *
from
    usuarios;