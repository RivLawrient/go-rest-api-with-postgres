CREATE TABLE wallet (
	id VARCHAR(100) NOT NULL,
	bank_name VARCHAR(100) NOT NULL,
	description VARCHAR(100),
	balance BIGINT DEFAULT 0 NOT NULL,
	PRIMARY KEY(id)
);

CREATE TABLE income (
	id VARCHAR(100) NOT NULL,
	source varchar(100) NOT NULL,
 	amount bigint DEFAULT 0 NOT NULL,
  	wallet_id varchar(100) NOT NULL,
  	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  	PRIMARY KEY (id),
  	CONSTRAINT fk_income_wallet FOREIGN KEY (wallet_id) REFERENCES wallet (id)
);

CREATE TABLE expense (
	id VARCHAR(100) NOT NULL,
	item VARCHAR(100) NOT NULL,
	quantity INT DEFAULT 1 NOT NULL,
	price BIGINT DEFAULT 0 NOT NULL,
	total_price BIGINT GENERATED ALWAYS AS (quantity * price) STORED,
	wallet_id VARCHAR(100) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
	PRIMARY KEY (id),
	CONSTRAINT fk_expense_wallet FOREIGN KEY (wallet_id) REFERENCES wallet (id)
);