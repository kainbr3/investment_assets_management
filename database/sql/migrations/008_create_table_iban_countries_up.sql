CREATE TABLE "iban_countries" 
("id" uuid DEFAULT uuid_generate_v4(),
"created_at" timestamp DEFAULT current_timestamp,
"updated_at" timestamp DEFAULT current_timestamp,
"is_active" boolean DEFAULT false,
"name" varchar(150) NOT NULL DEFAULT null,
"alpha2_code" varchar(2) NOT NULL DEFAULT null,
"alpha3_code" varchar(3) NOT NULL DEFAULT null,
"numeric_code" int NOT NULL DEFAULT 0,
PRIMARY KEY ("id"),
CONSTRAINT "unique_iban_countries_alpha2_code" UNIQUE ("alpha2_code"),
CONSTRAINT "unique_iban_countries_alpha3_code" UNIQUE ("alpha3_code"),
CONSTRAINT "unique_iban_countries_numeric_code" UNIQUE ("numeric_code"));