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

create table seguidores (
                            usuario_id int not null,
                            seguidor_id int not null,
                            primary key (usuario_id,seguidor_id ),
                            foreign key (usuario_id)
                                references usuarios(id)
                                on delete cascade,
                            foreign key (seguidor_id)
                                references usuarios(id)
                                on delete cascade
)

