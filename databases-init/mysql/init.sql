SET FOREIGN_KEY_CHECKS = 0;

CREATE TABLE clients (id varchar(255) PRIMARY KEY, name varchar(255), email varchar(255), created_at date);
CREATE TABLE accounts (id varchar(255) PRIMARY KEY, client_id varchar(255), balance NUMERIC(15,2), created_at date, FOREIGN KEY(client_id) REFERENCES clients(id));
CREATE TABLE transactions (id varchar(255) PRIMARY KEY, account_id_from varchar(255), account_id_to varchar(255), amount NUMERIC(15,2), created_at date, FOREIGN KEY(account_id_from) REFERENCES accounts(id), FOREIGN KEY(account_id_to) REFERENCES accounts(id));
	
INSERT INTO clients (id, name, email, created_at) 
  VALUES (
    'f23e0c58-23b3-48c3-8d6b-fe30c0666665',
    'John Doen',
    'j@j.com',
    CURDATE()
  );

INSERT INTO clients (id, name, email, created_at) 
  VALUES (
    '23e16a8d-4d42-4543-998b-cd7594a7e8b7',
    'Jane Doe',
    'jj@j.com',
    CURDATE()
  );

INSERT INTO accounts (id, client_id, balance, created_at)
  VALUES (
    '66607797-85a1-4492-9e78-47897d5c7f30',
    '23e16a8d-4d42-4543-998b-cd7594a7e8b7',
    2000,
    CURDATE()
  );

INSERT INTO accounts (id, client_id, balance, created_at)
  VALUES (
    'a4c59202-c51f-42c7-937e-ec0b4d438e82',
    'f23e0c58-23b3-48c3-8d6b-fe30c0666665',
    3000,
    CURDATE()
  );