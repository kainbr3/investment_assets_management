CREATE TABLE "iban_countries_currencies" 
("id" uuid DEFAULT uuid_generate_v4(),
"created_at" timestamp DEFAULT current_timestamp,
"updated_at" timestamp DEFAULT current_timestamp,
"is_active" boolean DEFAULT false,
"name" varchar(150) NOT NULL DEFAULT null,
"iban_country_id" uuid,
"iban_currency_id" uuid,
PRIMARY KEY ("id"),
CONSTRAINT "fk_iban_countries_currencies_iban_country" FOREIGN KEY ("iban_country_id") REFERENCES "iban_countries"("id"),
CONSTRAINT "fk_iban_countries_currencies_iban_currency" FOREIGN KEY ("iban_currency_id") REFERENCES "iban_currencies"("id"));