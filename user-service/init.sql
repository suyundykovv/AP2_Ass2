CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,            
    username VARCHAR(255) NOT NULL,    
    email VARCHAR(255) NOT NULL,  
    password VARCHAR(255) NOT NULL,    
    role VARCHAR(50) NOT NULL,      
    created_at TIMESTAMP NOT NULL     
);

