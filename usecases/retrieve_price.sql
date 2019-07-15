SELECT
  p.price
FROM
  kudaki_event.prices p
  JOIN kudaki_user.users u ON p.creator_user_uuid = u.uuid
WHERE
  p.duration_unit = "DAY"
  AND (
    u.role = "ADMIN"
    OR u.role = "KUDAKI_TEAM"
  );