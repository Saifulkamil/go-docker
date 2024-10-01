# OVERVIEW
The SPARI Project Test is designed to assess the skills of potential employees at PT Super PARI. This project is built using the Go programming language and developed natively without using any frameworks.

Golang Version: v1.23.1

Require Modules:
- github.com/go-sql-driver/mysql v1.8.1
- github.com/joho/godotenv v1.5.1

# SETUP
To run this project, the project setup must be completed first. The setup process can be done by following these steps:
1. Clone the project from github
```bash
git clone https://github.com/zulfahmidev/SPARI_PROJECT_TEST.git
```
2. Navigate to the project directory:
```bash
cd SPARI_PROJECT_TEST
```
3. Install necessary Go modules:
```bash
go mod tidy
```
4. Duplicate file `.env.example` and change filename to `.env`:
```bash
cp .env.example .env
```
5. Configuration database on `.env` file
```properties
DB_HOST=127.0.0.1 # Your database host
DB_PORT=3306 # Mysql default port
DB_USER=username # Your database username
DB_PASS=password # Your database password
DB_NAME=database_name # Your database name
DB_NAME_TEST=database_name # Your database host for unit test
```
6. Create the database using MySQL:
```bash
mysql -u username -p -e "CREATE DATABASE database_name;"
```
7. Import the database schema from the `database_setup.sql` file:
```bash
mysql -u username -p database_name < database_setup.sql
```
# RUNNING THE PROJECT
- To run the project, use:
```bash
go run main.go
```
- To run tests, use:
```bash
go test ./tests
```
# ENDPOINTS
### `GET /categories` | Get All Categories 
This endpoint retrieves a list of all categories stored in the database.

Example Response:
```json
{
    "message": "Categories loaded successfully",
    "data": [
        {
            "id": 1,
            "name": "example"
        }
    ]
}
```

### `POST /categories` | Create New Category
This endpoint creates a new category in the system.

Example Body:
```json
{
    "name": "example",
}
```
Example Response:
```json
{
    "message": "Category created successfully",
    "data": null
}
```
### `GET /items` | Get All Items 
This endpoint retrieves all items available in the database.

Example Response:
```json
{
    "message": "Items loaded successfully",
    "data": [
        {
            "id": 1,
            "category_id": 1,
            "name": "example_name",
            "description": "example_description",
            "price": 1000,
            "created_at": "2024-10-01 12:15:38"
        }
    ]
}
```
### `POST /items` | Create New Item
This endpoint creates a new item.

Example Body:
```json
{
    "name": "example_name",
    "description": "example_description",
    "price": 1000,
    "category_id": 1
}
```
Example Response:
```json
{
    "message": "Item created successfully",
    "data": null
}
```
### `GET /items/:id` | Get Item by ID
This endpoint retrieves a specific item by its ID.

Example Response:
```json
{
    "message": "Item loaded successfully",
    "data": {
        "id": 1,
        "category_id": 1,
        "name": "example_name",
        "description": "example_description",
        "price": 1000,
        "created_at": "2024-10-01 12:15:38"
    }
}
```
### `PUT /items/:id` | Update Item by ID
This endpoint updates an existing item by its ID

Example Body:
```json
{
    "name": "example2",
    "description": "example_description2",
    "price": 2000,
    "category_id": 2
}
```
Example Response:
```json
{
    "message": "Item updated successfully",
    "data": null
}
```
### `DELETE /items/:id` | Delete Item by ID
This endpoint deletes a specific item by its ID.

Example Response:
```json
{
    "message": "Item deleted successfully",
    "data": null
}
```