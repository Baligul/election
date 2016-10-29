To run server in background
==========================================
./election & <enter>
<enter>

To make service foreground
----------------------------
fg ./election <enter>


==========  start PostgreSQL Server===========================

sudo su root
su - postgres
cd bin
/usr/local/pgsql/bin/pg_ctl -D /usr/local/pgsql/data -l logfile start
