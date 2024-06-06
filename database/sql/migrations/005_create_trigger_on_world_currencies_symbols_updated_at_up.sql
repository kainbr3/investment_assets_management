CREATE TRIGGER update_world_currencies_symbols_timestamp
BEFORE UPDATE ON world_currencies_symbols
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();