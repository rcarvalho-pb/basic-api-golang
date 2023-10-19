use devbook;

insert into users(name, nick, email, password)
values
("Ramon Carvalho", "LoPapelito", "ramon@email.com", 123),
("Emilly Coeli", "Mimoquinha", "emilly@email.com", 123),
("Teste Testando", "LoTestelito", "teste@email.com", 123);

insert into followers
values
(1, 2),
(1, 3),
(2, 3);