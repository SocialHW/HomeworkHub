###Database Structure

SQLite uses a .db file to represent the database. We should have some .db file located in the db directory. This file 
will contain all of the tables we use to store information related to the site.
___

There should be a Posts table with the following columns:

post_id primary key, auto inc INTEGER | username TEXT | title TEXT | file_path TEXT 
___

There should be a Comments table with the following columns:

post_id INTEGER | username TEXT | comment TEXT
___

There should be a user info table with the following columns:

ID primary key, auto inc INTEGER | Username TEXT | password TEXT 
___