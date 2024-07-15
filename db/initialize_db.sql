/* db clean-up */
drop owned by gonitor cascade;

/* db structure */
\i 'structure/tables.sql'
\i 'structure/views.sql'
\i 'structure/triggers.sql'

/* load data */
\i 'data/roles.sql'
\i 'data/agents.sql'
\i 'data/users.sql'
\i 'data/endpoints.sql'
\i 'data/grants.sql'
\i 'data/host_groups.sql'
\i 'data/hosts.sql'
\i 'data/lnk_hosts_groups.sql'
\i 'data/items.sql'
\i 'data/lnk_hosts_items.sql'
\i 'data/events.sql'