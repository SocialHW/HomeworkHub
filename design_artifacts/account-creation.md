###Account Creation

Users will be able to create an account by clicking a link in the top right corner of 
the home page with the text "Register". The link will direct the user to /register where 
they can provide a username, and a password. They submit the request to create an 
account by pressing a button labeled "Create Account". If the username is not taken, they
will be redirected to the index page. From there they can continue to login by clicking 
a link in the top right corner with the text "Login". From there they can enter the 
username and password they provided before. They will be redirected to the index page,
where they will now be logged in.

When the user clicks on the Create Account button, they will be submitting an HTTP POST
request, to /register/details with the information they provided. There will be a 
handler function listening on /register/details which will query the user info table
to find an account with a matching username. If there is no account with a matching 
username, the account can be created by inserting a new row into the database. 