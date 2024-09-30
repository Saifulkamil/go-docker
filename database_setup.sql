CREATE TABLE categories (
    id INTEGER PRIMARY KEY AUTOINCRMENT,
    name VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE items (
    id INTEGER PRIMARY KEY AUTOINCRMENT,
    category_id INTEGER,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price REAL NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (category_id) REFERENCES categories(id)
);

CREATE idx_category_name ON categories(name);
CREATE idx_item_name ON items(name);
CREATE idx_item_category_id ON items(category_id);