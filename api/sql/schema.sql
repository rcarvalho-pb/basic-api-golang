create database if not exists devbook;
use devbook;

drop table if exists users;

create table users(
  id int auto_increment primary key,
  name varchar(255) not null,
  nick varchar(255) not null unique,
  email varchar(255) not null unique,
  password varchar(255) not null,
  created_at timestamp default current_timestamp()
) engine = INNODB;