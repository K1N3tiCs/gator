-- name: CreateFeedFollow :many
WITH insert_feed_follows AS (
  INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
  VALUES(
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2
  )
  RETURNING *
)
SELECT iff.*, u.name as user_name, f.name as feed_name
FROM insert_feed_follows iff
JOIN feeds f ON f.id = iff.feed_id
JOIN users u ON u.id = iff.user_id;

-- name: GetFeedFollowsForUser :many
SELECT ff.*, u.name AS user_name, f.name AS feed_name FROM feed_follows ff
INNER JOIN feeds f ON f.id = ff.feed_id
INNER JOIN users u ON u.id = ff.user_id
WHERE u.name = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows
WHERE user_id = $1
AND feed_id = $2;
