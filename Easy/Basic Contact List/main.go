package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Contact struct {
	Name  string
	Phone string
	Email string
}

func addContact(c []Contact, newCon []Contact) []Contact {

	for _, newContact := range newCon {
		c = append(c, Contact{
			Name:  newContact.Name,
			Phone: newContact.Phone,
			Email: newContact.Email,
		})
	}

	return c

}

func displayContacts(c []Contact) {
	for i, c := range c {
		fmt.Printf("%d. %s - Phone: %s, Email: %s\n", i+1, c.Name, c.Phone, c.Email)
	}

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()x
	contactsInput := scanner.Text()

	//development
	// contactsInput := "Rachel Green:555-1357:rachel@email.com,Monica Geller:555-2468:monica@email.com,Phoebe Buffay:555-3691:phoebe@email.com,Ross Geller:555-4820:ross@email.com,Chandler Bing:555-5173:chandler@email.com,Joey Tribbiani:555-6284:joey@email.com"
	// newContactInput := "Gunther Central:555-7395:gunther@email.com"

	scanner.Scan()
	newContactInput := scanner.Text()

	contactEntries := strings.Split(contactsInput, ",")

	// contact entry

	contactData := make([]Contact, 0)

	for _, c := range contactEntries {
		trimmedContact := strings.Split(strings.TrimSpace(c), ":")

		name := trimmedContact[0]
		phone := trimmedContact[1]
		email := trimmedContact[2]

		contactData = append(contactData, Contact{
			Name:  name,
			Phone: phone,
			Email: email,
		})

	}

	//new contact
	newContact := make([]Contact, 0)

	newC := strings.Split(newContactInput, ":")

	name := newC[0]
	phone := newC[1]
	email := newC[2]

	newContact = append(newContact, Contact{
		Name:  name,
		Phone: phone,
		Email: email,
	})

	// start
	fmt.Println("=== CONTACT MANAGEMENT SYSTEM ===")
	fmt.Println("Initial contact list:")
	displayContacts(contactData)

	totalContacts := 0
	totalSession := 0
	for _, c := range newContact {
		fmt.Printf("Adding new contact: %s (%s) - %s\n", c.Name, c.Phone, c.Email)
		contactList := addContact(contactData, newContact)
		fmt.Println("Updated contact list:")
		displayContacts(contactList)
		totalContacts = len(contactList)
		totalSession++
	}

	fmt.Println("=== CONTACT STATISTICS ===")
	fmt.Printf("Total contacts: %d\n", totalContacts)
	fmt.Printf("Contacts added this session: %d\n", totalSession)
	fmt.Println("Contact management operations completed successfully")

}
