# HomeworkHub

**Social media platform for students to upload and view homework assignments**

*Homework hub is a platform for students to find old homeworks from previous semesters of a class, to enhance their 
studying ability*

## Installing

Download the source using the command:

```bash
git clone https://github.com/SocialHW/HomeworkHub.git
``` 

Running this project no longer requires any external dependencies to be running, such as MySQL.

## Building and Running

To run the project as a Go script, simply run this command from the root directory of the project:

```bash
go run main.go
```


## Contributing

We are using [git-flow](https://github.com/nvie/gitflow) 
[(installation instructions)](https://github.com/nvie/gitflow/wiki/Installation) to aid managing our branches. To make 
sure that you have git-flow initialized in your repository use the command:

```bash
git flow init
```


To create a new feature, use the command:

```bash
git flow feature start <feature-name>
```

This will create a new branch named `feature/<feature-name>`, and checkout that branch so you can immediately start work
on it. When the feature is finished and you would like to merge, make sure you have commited and pushed all of your
changes, and head to GitHub where you can make a new pull request for that feature. Make sure to merge into the develop
branch.