CREATE TRIGGER update_iban_currencies_timestamp
BEFORE UPDATE ON iban_currencies
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();