use devbook;

insert into users(name, nick, email, password)
values
("Ramon Carvalho", "LoPapelito", "ramon@email.com", "$2a$10$GKn6zTjR25rvJv926ke8EesZdn7HAt.epHNGbUZllKImCZwPFvK5y"),
("Emilly Coeli", "Mimoquinha", "emilly@email.com", "$2a$10$NwFhSan0KKbfhd8pt03yMeguU5UoXYGq.sj5K2rTdoA7NhswnhEg6"),
("Teste Testando", "LoTestelito", "teste@email.com", "$2a$10$6mTL9SySGkbbYhSr4N75bu6XXJRDNQnvxn8j1uoUeSctTPV12/XHe");

insert into followers
values
(1, 2),
(1, 3),
(2, 3);

insert into publications (title, content, author_id) values
("publi do usuario 1", "publication from user 1", 1),
("publi do usuario 2", "publication from user 2", 2),
("publi do usuario 3", "publication from user 3", 3);