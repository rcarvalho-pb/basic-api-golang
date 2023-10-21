create database if not exists devbook;
use devbook;

drop table if exists publications;
drop table if exists followers;
drop table if exists users;

create table users(
  id int auto_increment primary key,
  name varchar(255) not null,
  nick varchar(255) not null unique,
  email varchar(255) not null unique,
  password varchar(255) not null,
  created_at timestamp default current_timestamp()
) engine = INNODB;

create table followers(
  user_id int not null,
  foreign key(user_id)
  references users(id)
  on delete cascade,

  follower_id int not null,
  foreign key(follower_id)
  references users(id)
  on delete cascade,

  primary key(user_id, follower_id)
) engine=INNODB;

create table publications(
  id int auto_increment primary key,
  title varchar(50) not null,
  content varchar(300) not null,

  author_id int not null,
  foreign key(author_id)
  references users(id)
  on delete cascade,

  likes int default 0,
  created_at timestamp default current_timestamp
) engine=INNODB;