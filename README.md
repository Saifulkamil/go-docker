# OVERVIEW
Lorem ipsum dolor sit amet consectetur, adipisicing elit. Minus unde assumenda enim natus accusamus tempore, sed asperiores repellat beatae quis, tempora numquam. Odit non minus rem dicta repellat at vitae aut aspernatur similique! Impedit eos temporibus nam nihil tempora officiis, expedita necessitatibus consequuntur quaerat tenetur? Quisquam officiis nihil quasi dolore.

Golang Version: v1.23.1

Require Modules:
- github.com/go-sql-driver/mysql v1.8.1
- github.com/joho/godotenv v1.5.1

# SETUP
Lorem ipsum dolor sit amet consectetur, adipisicing elit. Minus unde assumenda enim natus accusamus tempore, sed asperiores repellat beatae quis, tempora numquam. 
1. Clone project from github
```
git clone git clone https://github.com/zulfahmidev/SPARI_PROJECT_TEST.git
```
2. Open project directory
```
cd SPARI_PROJECT_TEST
```
3. Install modules
```
go mod tidy
```
3. Duplicate file .env.example and change filename to .env
```
cp .env.example .env
```
4. Configuration database on .env file
```
DB_HOST=127.0.0.1 # Your database host
DB_PORT=3306 # Mysql default port
DB_USER=username # Your database username
DB_PASS=password # Your database password
DB_NAME=database_name # Your database name
DB_NAME_TEST=database_name # Your database host for unit test
```
5. Create database. This example create database using mysql in terminal
```
mysql -u username -p -e "CREATE DATABASE database_name;"
```
6. Import database from file database_setup.sql
```
mysql -u username -p database_name < database_setup.sql
```
6. Import database from file database_setup.sql
```
mysql -u username -p database_name < database_setup.sql
```
# RUN PROJECT
- Command to run project
```
go run main.go
```
- Command to run test
```
go test ./tests
```
# ENDPOINTS
