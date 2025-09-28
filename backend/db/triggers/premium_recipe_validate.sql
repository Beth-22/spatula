-- Function to validate premium recipes
CREATE OR REPLACE FUNCTION validate_premium_recipe()
RETURNS TRIGGER AS $$
BEGIN
  IF NEW.is_premium = true AND (NEW.price IS NULL OR NEW.price <= 0) THEN
    RAISE EXCEPTION 'Premium recipes must have a price greater than 0';
  END IF;
  
  IF NEW.is_premium = false AND NEW.price IS NOT NULL AND NEW.price > 0 THEN
    RAISE EXCEPTION 'Free recipes cannot have a price';
  END IF;
  
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger on recipes table
CREATE TRIGGER validate_premium_recipe_trigger
  BEFORE INSERT OR UPDATE ON recipes
  FOR EACH ROW EXECUTE FUNCTION validate_premium_recipe();