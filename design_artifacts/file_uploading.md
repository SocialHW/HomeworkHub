###File Uploading

Uploading images to the site represent the core value of the project. Users will have the ability to upload jpg/jpeg, 
png, gif, and pdf files to be stored on the site. There will be a maximum upload size of 8MB in order to restrict users
from overusing the resources of the web server. Each file will be associated with a post, which is a row in the Posts
table containing information such as the user who created the post, the title, and the extension of the file. The file
will be saved in the location "posts/<postid>.<ext>". Since each post will have a unique ID, this will prevent
collisions when storing files on the disk.

In order to create a new post, a user must be logged in and will click a link in the top right corner of the web page.
The user will be redirected to /create-post, where they will be prompted for the information required for a new post.

It should not be allowed to use punctuation such as ;, in order to avoid SQL injection.