/* 
 * Create db and user 
 * ==============================================
 *
 * 1. run psql as the postgres superuser
 * sudo -i -u postgres psql
 *
 * 2. once in the psql prompt run the cmds below
 *
 */
 
create database gonitor
with encoding = 'UTF8'
lc_collate = 'en_US.UTF-8'
lc_ctype = 'en_US.UTF-8'
template template0
connection limit = 25;

create user gonitor password 'your_secret_pwd';
grant all on database gonitor to gonitor;
\c gonitor;
grant create on schema public to gonitor;

exit;

/* 
 * Initialize db - both structure and data
 * =============================================
 *
 * 1. Upload the whole db folder to the database server
 *
 * 2. At the database server, change the working directory to the uploaded db directory
 * Example
 * cd ~/gonitor/db
 * 
 * 3. Excecute the initialize.sql file
 * Example
 * psql -h localhost -d gonitor -U gonitor -f initialize_db.sql
 *
 */


/* 
 * Load test data
 * =============================================
 *
 * 1. At the database server, change the working directory to the uploaded db directory
 * Example
 * cd /gonitor/db
 * 
 * 2. Excecute the load_test_data.sql file
 * Example
 * psql -h localhost -d gonitor -U gonitor -f load_test_data.sql
 *
 */