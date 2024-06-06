CREATE TABLE "world_currencies_symbols" 
("id" uuid DEFAULT uuid_generate_v4(),
"created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
"updated_at" timestamp DEFAULT CURRENT_TIMESTAMP,
"is_active" boolean DEFAULT true,
"name" varchar(150) NOT NULL DEFAULT null,
"code" varchar(3) NOT NULL DEFAULT null,
"symbol" varchar(4) NOT NULL DEFAULT null,
PRIMARY KEY ("id"),
CONSTRAINT "unique_world_currencies_symbols_code" UNIQUE ("code"));