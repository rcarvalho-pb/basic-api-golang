-- name: CreateUser :execresult
insert into users(name, nick, email, password) values (?, ?, ?, ?);

-- name: FindUser :many
select * from users where name like ? or nick like ?;

-- name: GetUserById :one
select * from users where id = ?;

-- name: UpdateUserById :exec
update users set name = ?, nick = ?, email = ?, password = ? where id = ?;

-- name: DeleteUserById :exec
delete from users where id = ?;

-- name: GetUserByEmailOrNick :one
select * from users where email like ? or nick like ?;

-- name: FollowUser :execresult
insert ignore into followers (user_id, follower_id) values (?, ?);