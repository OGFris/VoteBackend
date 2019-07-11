# VoteBackend
A backend server prototype for the Algerian's voting system.

## Installation

 **Requirements:**
    - Mariadb v10.x or higher server.
    - Create a .env following the .env.example patterns to configure the server.
    
 **Setting up the Mysql server:**
 
 ```sh
 # Step 1: Enter the MariaDB Cli
 > mysql -u root
 
 # Step 2: Create the databases in command prompt
 MariaDB [(none)]> create database votes default character set utf8 default collate utf8_general_ci;
 ```
 
## Docs

 - *GET:/candidates*: this route returns a json array of every candidate stored in the database.
 
 - *POST:/candidate*: this route create a new candidate in the database.