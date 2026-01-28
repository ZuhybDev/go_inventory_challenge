Write a function formatPriceTags that takes itemNames, baseCosts, markupRates and returns a string array of formatted price tags.

The function calculates final prices by applying markup rates to base costs, sorts items by price (lowest to highest), and formats each as an aligned price tag.

Logic:

Calculate final price: baseCost \* (1 + markupRate)
Sort items by final price in ascending order
Format each tag with item name left-aligned in 20 characters and price right-aligned with 2 decimal places
Parameters:

itemNames ([]string): Array of repair item names
baseCosts ([]float64): Array of base repair costs in dollars
markupRates ([]float64): Array of markup rates as decimals (0.25 = 25%)
Returns: String array of formatted price tags, sorted by price. Format: ["Bike Tune-up $37.50", "Phone Screen $62.50"]
