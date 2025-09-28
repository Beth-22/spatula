CREATE OR REPLACE FUNCTION calculate_avg_rating()
RETURNS TRIGGER AS $$
BEGIN
    -- Update recipe's average rating when ratings change
    UPDATE recipes 
    SET average_rating = (
        SELECT COALESCE(AVG(rating), 0) 
        FROM ratings 
        WHERE recipe_id = NEW.recipe_id
    )
    WHERE id = NEW.recipe_id;
    
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create trigger for ratings table
CREATE TRIGGER update_recipe_rating
    AFTER INSERT OR UPDATE OR DELETE ON ratings
    FOR EACH ROW
    EXECUTE FUNCTION calculate_avg_rating();