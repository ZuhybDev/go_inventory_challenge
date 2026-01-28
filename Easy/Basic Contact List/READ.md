Create a contact management system that demonstrates how structs, slices, and functions work together to organize and display contact information. This challenge will test your ability to define custom data structures, manage collections of data, and implement functions to manipulate that data.

You will receive two inputs:

A string containing contact information in the format "name1:phone1:email1,name2:phone2:email2,name3:phone3:email3" (e.g., "Alice Johnson:555-0123:alice@email.com,Bob Smith:555-0456:bob@email.com,Carol Davis:555-0789:carol@email.com")
A string containing a new contact to add in the format "name:phone:email" (e.g., "David Wilson:555-0321:david@email.com")
Your task is to:

Define a Contact struct with three string fields: Name, Phone, and Email
Create a function called addContact that takes a slice of Contact structs and a new Contact, then returns a new slice with the contact added
Create a function called displayContacts that takes a slice of Contact structs and prints each contact's information
Parse the first input by splitting on commas to get individual contact entries
For each contact entry, split on colons to extract name, phone, and email
Create Contact structs from the parsed data and store them in a slice
Parse the second input by splitting on colons to create a new contact
Display the application header: "=== CONTACT MANAGEMENT SYSTEM ==="
Display the initial contacts using your displayContacts function with the header: "Initial contact list:"
Display the new contact being added: "Adding new contact: [name] ([phone]) - [email]"
Use your addContact function to add the new contact to the slice
Display the updated contacts using your displayContacts function with the header: "Updated contact list:"
Display contact statistics:
"=== CONTACT STATISTICS ==="
"Total contacts: [number_of_contacts]"
"Contacts added this session: 1"
Display the completion message: "Contact management operations completed successfully"
The displayContacts function should print each contact in the format: "[index]. [name] - Phone: [phone], Email: [email]" where index starts from 1.

Use the strings package to split input strings and the fmt package for formatted output. This challenge demonstrates how to combine structs for data organization, slices for collections, and functions for code organization - fundamental patterns used throughout Go programming.
