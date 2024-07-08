/* db clean-up */
drop owned by gonitor cascade;

/* db structure */
\i 'structure/tables.sql'
\i 'structure/views.sql'
\i 'structure/triggers.sql'

/* load data */
\i 'data/roles.sql'
\i 'data/users.sql'