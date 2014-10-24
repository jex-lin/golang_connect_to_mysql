# Install Mercurial

    sudo apt-get install mercurial meld

# Install mysql-server and golang-mysql driver

    sudo apt-get install mysql-server
    go get github.com/go-sql-driver/mysql

# Create mysql data for testing

    CREATE DATABASE go_test
    CREATE TABLE tt (name char(10));
    Insert tt values ('jex');
    Insert tt values ('bob');
    Insert tt values ('ted');

# Execute `go run connect_to_mysql.go` and then you will see :

    jex
    bob
    ted
