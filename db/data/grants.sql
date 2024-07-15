/* 
 * Grants are only required for endpoints to be authenticated/authorized with JWT.
 * Endpoints to be authenticated/authorized with HTTP Basic Authentication or public/open do not require this.   
 */

/************************ user module ************************/

/****** account ctrl ******/

/* fetch account */
INSERT INTO grants VALUES (DEFAULT, 200101, 2, NOW(), '2099-01-01 10:00:00'); 
INSERT INTO grants VALUES (DEFAULT, 200101, 3, NOW(), '2099-01-01 10:00:00'); 

/* update account */
INSERT INTO grants VALUES (DEFAULT, 200102, 2, NOW(), '2099-01-01 10:00:00');
INSERT INTO grants VALUES (DEFAULT, 200102, 3, NOW(), '2099-01-01 10:00:00');

/****** role ctrl ******/

/* fetch a roles list */
INSERT INTO grants VALUES (DEFAULT, 200201, 3, NOW(), '2099-01-01 10:00:00'); 

/****** endpoint ctrl ******/

/* fetch an endpoints list */
INSERT INTO grants VALUES (DEFAULT, 200301, 3, NOW(), '2099-01-01 10:00:00'); 

/****** user ctrl ******/

/* fetch an users list */
INSERT INTO grants VALUES (DEFAULT, 200401, 3, NOW(), '2099-01-01 10:00:00'); 

/* find an user */
INSERT INTO grants VALUES (DEFAULT, 200402, 3, NOW(), '2099-01-01 10:00:00'); 

/* create a new user */
INSERT INTO grants VALUES (DEFAULT, 200403, 3, NOW(), '2099-01-01 10:00:00'); 

/* update an existing user */
INSERT INTO grants VALUES (DEFAULT, 200404, 3, NOW(), '2099-01-01 10:00:00'); 

/* delete an user */
INSERT INTO grants VALUES (DEFAULT, 200405, 3, NOW(), '2099-01-01 10:00:00'); 

/****** event ctrl ******/

/* fetch an events list */
INSERT INTO grants VALUES (DEFAULT, 200501, 3, NOW(), '2099-01-01 10:00:00'); 