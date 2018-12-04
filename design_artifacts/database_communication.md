###Database Communication

We should have a file at the db directory of the project that contains all of the functions needed to interact with the 
database.
* There should be a RegisterUser function that will take the email and password as strings, and executes the queries to 
insert the user into the database if it is a valid registration.
* There should be a LoginUser function that will take a username and password as strings, query the database to for the 
password associated with the username entered. If the username exists and the password entered is equal to the password 
in the database, the user is redirected to the home page.  
* There should be a GetPosts function that takes a predicate function (of Post -> Boolean) and returns []Post, 
containing the Posts in the database that meet the predicate. Future implementations may use a comparator function to 
rank the results, such as closest search result match.
* There should be a CreatePost function, that takes a Post struct and inserts all of its fields into. 
