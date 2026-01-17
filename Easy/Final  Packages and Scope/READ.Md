Complete your inventory management system by creating a comprehensive main program that integrates all the functionality you've built in previous lessons. This final challenge brings together inventory initialization, stock checking, item addition, stock updates, and report generation into a complete command-line interface with menu-driven navigation.

You will receive three inputs:

A string containing initial inventory data in the format "product1:price1:quantity1,product2:price2:quantity2,product3:price3:quantity3" (e.g., "Laptop:999.99:5,Mouse:25.50:15,Keyboard:75.00:8")
A string containing a sequence of menu operations in the format "operation1,operation2,operation3" where operations can be "check", "add", "update", "report", or "exit" (e.g., "check,add,update,report,exit")
A string containing operation parameters in the format "param1|param2|param3" where each parameter corresponds to the operation at the same position (e.g., "Mouse|Monitor:299.99:10|Laptop:5|full,10")
Your task is to:

Use the same Product struct from previous lessons with Price (float64) and Quantity (int) fields
Parse the first input and initialize the inventory map with the existing product data
Display the system startup message:
"=== INVENTORY MANAGEMENT SYSTEM ==="
"System initialized with [number_of_products] products"
"Starting interactive session..."
Parse the second input to get the sequence of operations to perform
Parse the third input to get the parameters for each operation
For each operation in sequence, display the operation header and perform the corresponding action:
For "check": Display "--- STOCK CHECK ---" and check stock for the specified product
For "add": Display "--- ADD ITEM ---" and add the specified product
For "update": Display "--- UPDATE STOCK ---" and update stock for the specified product
For "report": Display "--- GENERATE REPORT ---" and generate the specified report
For "exit": Display "--- SYSTEM EXIT ---" and show exit information
Implement the checkStock function that returns quantity and error as in previous lessons
Implement the addNewItem function that adds products and returns error as in previous lessons
Implement the updateStock function that updates quantities and returns error as in previous lessons
Implement the generateReport function that creates reports as in previous lessons
For the "check" operation:
Use the parameter as the product name to check
Display: "Checking stock for: [product_name]"
If found: "Stock level: [quantity] units"
If not found: "Product not found in inventory"
For the "add" operation:
Parse the parameter in format "product_name:price:quantity"
Display: "Adding new product: [product_name]"
If successful: "Product added successfully"
If failed: "Failed to add product: [error_message]"
For the "update" operation:
Parse the parameter in format "product_name:change"
Display: "Updating stock for: [product_name]"
If successful and change is positive: "Added [change] units. New stock: [new_quantity]"
If successful and change is negative: "Removed [absolute_change] units. New stock: [new_quantity]"
If failed: "Update failed: [error_message]"
For the "report" operation:
Parse the parameter in format "report_type,threshold"
Display: "Generating [report_type] report with threshold [threshold]"
Generate and display the full report as in previous lessons
For the "exit" operation:
Display final inventory statistics:
"Final inventory status:"
"Total products: [total_products]"
"Total items: [total_items]"
"Total value: $[total_value]"
Display: "Session completed successfully"
Display: "Thank you for using the Inventory Management System"
After each operation (except exit), display: "Operation completed. Continuing to next operation..."
Use the strings package to split input strings, the strconv package to convert strings to numbers, the errors package for error handling, the fmt package for formatted output, and the sort package for alphabetical sorting. Format all prices and values to 2 decimal places. This challenge demonstrates how to create a complete interactive system that integrates multiple functional components into a cohesive user experience with proper error handling and state management.
