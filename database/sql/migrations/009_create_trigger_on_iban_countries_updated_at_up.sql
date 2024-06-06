CREATE TRIGGER update_iban_countries_timestamp
BEFORE UPDATE ON iban_countries
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();