package main

import "fmt"

// Women's Travel & Adventure Group
type wtag struct {
        // Members
        members []member

        // Adventures
        adventures []adventure
}

// Member of Women's Travel & Adventure Group
type member struct {
        // First name
        firstName string

        // Last name
        lastName string

        // Age
        age int
}

// Adventure suggested by a Women's Travel & Adventure Group
type adventure struct {
        // Name
        name string

        // Description
        desc string

        // Difficulty
        difficulty int
}

// Constructor for creating a new Women's Travel & Adventure Group
func NewWomenTravelAdventureGroup() wtag {
        return wtag{}
}

// Adds a new member to the group
func (w *wtag) AddMember(firstName, lastName string, age int) {
        m := member{
                firstName : firstName,
                lastName : lastName,
                age : age,
        }

        w.members = append(w.members, m)
}

// Adds a new adventure to the group
func (w *wtag) AddAdventure(name, desc string, difficulty int) {
        a := adventure{
                name : name,
                desc : desc,
                difficulty : difficulty,
        }

        w.adventures = append(w.adventures, a)
}

// Lists the group members
func (w *wtag) ListMembers() {
        fmt.Printf("Members of Women's Travel & Adventure Group:\n")
        for _, m := range w.members {
                fmt.Printf("- %s %s (%d)\n", m.firstName, m.lastName, m.age)
        }
}

// Lists the group adventures
func (w *wtag) ListAdventures() {
        fmt.Printf("Adventures of Women's Travel & Adventure Group:\n")
        for _, a := range w.adventures {
                fmt.Printf("- %s (%d)\n", a.name, a.difficulty)
        }
}

// Randomly choose an adventure for the group
func (w *wtag) ChooseAdventure() *adventure {
        if len(w.adventures) == 0 {
                return nil
        }

        randIndex := rand.Intn(len(w.adventures))
        return &w.adventures[randIndex]
}

// Main program
func main() {
        // Create a new Women's Travel & Adventure Group
        wtag := NewWomenTravelAdventureGroup()

        // Add members
        wtag.AddMember("Jane", "Doe", 23)
        wtag.AddMember("John", "Doe", 25)
        wtag.AddMember("Alice", "Smith", 30)

        // Add adventures
        wtag.AddAdventure("Mountain Hiking", "Hike the highest mountain and take in the amazing views", 4)
        wtag.AddAdventure("Trekking", "Explore the wilderness with a guide and sleep under the stars", 3)
        wtag.AddAdventure("Paragliding", "Take a leap of faith and soar through the sky with a professional instructor", 5)

        // List members
        wtag.ListMembers()

        // List adventures
        wtag.ListAdventures()

        // Choose an adventure
        chosenAdventure := wtag.ChooseAdventure()
        fmt.Printf("Let's go for: %s\n", chosenAdventure.name)
}