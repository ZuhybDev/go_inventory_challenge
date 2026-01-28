Write a function evaluateFirstAidKit that takes supplies, householdSize, currentDate and returns a boolean.

The function determines if a first aid kit meets safety requirements by checking essential items, quantities, and expiration dates.

Logic:

Check if all essential items are present: "bandages", "antiseptic", "pain_reliever"
Verify quantities meet household requirements (minimum 2 per person for each essential item)
Ensure no essential items are expired (expiration date >= current date)
Return true only if all three conditions are met
Parameters:

supplies ([]string): Array where each element is "name:category:expiration:quantity" (expiration format: "YYYY-MM-DD")
householdSize (int): Number of people in household
currentDate (string): Current date in "YYYY-MM-DD" format
Returns: Boolean indicating if the kit passes all safety checks.
