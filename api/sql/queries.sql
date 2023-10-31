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

-- name: UnfollowUser :execresult
delete from followers where user_id = ? and follower_id = ?;

-- name: GetAllUserFollow :many
select u.id, u.name, u.nick, u.email, u.created_at from users u
inner join followers f on u.id = f.follower_id where f.user_id = ?;

-- name: GetAllUserFollowed :many
select u.id, u.name, u.nick, u.email, u.created_at from users u
inner join followers f on u.id = f.user_id where f.follower_id = ?;

-- name: UpdatePassword :exec
update users set password = ? where id = ?;

-- name: CreatePublication :execresult
insert into publications (title, content, author_id, likes) values (?, ?, ?, ?);

-- name: FindPublicationById :one
select p.*, u.nick from publications p inner join users u where p.id = ?;

-- name: FindPublications :many
select distinct p.*, u.nick from publications p inner join users u on u.id = p.author_id inner join followers f on p.author_id = f.user_id where u.id = ? or f.follower_id = ? order by 1 desc;

-- name: UpdatePublication :exec
update publications set title = ?, content = ? where id = ?;

-- name: DeletePublicationById :exec
delete from publications where id = ?;

-- name: GetUserPublications :many
select p.*, u.nick from publications p inner join users u on u.id = p.author_id where author_id = ?;

-- name: LikePublication :exec
update publications set likes = likes + 1 where id = ?;

-- name: DislikePublication :exec
update publications set likes =
case 
  when likes > 0 then likes - 1
  else likes
end
where id = ?;