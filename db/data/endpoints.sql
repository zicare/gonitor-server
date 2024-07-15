/* 
 * Only required for endpoints to be authenticated/authorized with JWT.
 * Endpoints to be authenticated/authorized with HTTP Basic Authentication or public/open do not require this.   
 */

/* u - users */

INSERT INTO endpoints VALUES (200101, 'GET',    '/u/account');
INSERT INTO endpoints VALUES (200102, 'PUT',    '/u/account');

INSERT INTO endpoints VALUES (200201, 'GET',    '/u/roles');

INSERT INTO endpoints VALUES (200301, 'GET',    '/u/endpoints');

INSERT INTO endpoints VALUES (200401, 'GET',    '/u/users');
INSERT INTO endpoints VALUES (200402, 'GET',    '/u/users/:user_id');
INSERT INTO endpoints VALUES (200403, 'POST',   '/u/users');
INSERT INTO endpoints VALUES (200404, 'PUT',    '/u/users/:user_id');
INSERT INTO endpoints VALUES (200405, 'DELETE', '/u/users/:user_id');

INSERT INTO endpoints VALUES (200501, 'GET',    '/u/events');