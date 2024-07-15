/*
 * view_grants
 */
CREATE VIEW view_grants AS
SELECT 
  g.grant_id,
  g.endpoint_id,
  g.role_id,
  g.start_date,
  g.end_date,
  e.route,
  e.method
FROM 
  grants g
JOIN
  endpoints e ON e.endpoint_id = g.endpoint_id
JOIN
  roles r ON r.role_id = g.role_id;