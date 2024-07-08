
/**************************** Below are few postgres related tips *****************************/

/*
 * 1. run psql as the postgres superuser
 * sudo -i -u postgres psql
 *
 * 2. from the psql prompt you can run the cmds below
 *
\c <database_name> (connect to database)
\l (list databases)
\dn (list schemas)
\dt (list tables)
\dv (list views)
\du (list roles)
 *
 */

/*
 * login psql as the user gonitor
 *
 * 1. login psql as the user gonitor
 * psql -h localhost -d gonitor -U gonitor
 *
 */

/*
 * Import data from csv
 *
 * 1. Save the CSV file to /tmp in local server, make sure the encoding is correct
 *
 * 2. login as the postgres linux user
 * sudo -i -u postgres
 * 
 * 3. run copy cmd through psql
 * psql -h 127.0.0.1 -U gonitor -d gonitor -c "\COPY zips(zip,settlement,settlement_type,municipality,state,city) FROM 'C:\tmp\zips.csv' with (FORMAT CSV, DELIMITER ',');"
 * 
 */

/*
 * Drop everything in the database owned by one specific user
 *
 * 1. login psql as the user gonitor
 * psql -h localhost -d gonitor -U gonitor
 *
 * 2. run drop query below

drop owned by gonitor cascade;

 *
 */


/*
 * Retrieve grants for specific user
 *
 * 1. login as the postgres user
 * psql -h localhost -d gonitor -U gonitor
 *
 * 2. run select query below
 *

SELECT table_catalog, table_schema, table_name, privilege_type
FROM   information_schema.table_privileges 
WHERE  grantee = 'gonitor';

 */