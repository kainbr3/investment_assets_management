CREATE TRIGGER update_iban_countries_currencies_timestamp
BEFORE UPDATE ON iban_countries_currencies
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();