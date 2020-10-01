# Inkrementalna migracija baze

## Uporaba 
SQL baze ob nadgradnjah in spremembah ne smemo ročno spreminjat saj je tako velika možnost da na kaj pozabimo, spremenimo
kaj kar nebi smeli ali kaj podobnega. Prav tako moremo ob določenih nadgradnjah včasih nekatere obstoječe podatke na novo 
preračunat, spremenit ali dodat neke nove. Ta proces moremo definirat tako, da je avtomatski in ponovljiv na različnih okoljih
(dev, produkcija..)

Tu je datoteka Migrator.go, ki lahko te stvari izvede.

Najboljše da datoteko [Migrator.go](Migrator.go) dodate v vsak projekt kjer imate SQL bazo podatkov. Tu je struktura migrator,
ki ima funkcijo Migrate. Takoj po povezavi z bazo inicializirajte strukturo migrator z povezavo na bazo in poljem definicij
ter nato pokličite funkcijo Migrate.

Vsaka definicija vsebuje 5 parametrov. 
 - Verzija baze
 - Opis migracije
 - Funkcija, ki se izvede pred skripto
 - Skripta
 - Funkcija, ki se izvede po skripti


## Primer - [Migrator.go](Migrator.go)

Primer polja definicij
   ```go
        var definitions = []migrationDefinition{
        	{
        		Version:     1,
        		Description: "Initial database creation",
        		Script:      createDB,
        	},
        	{
        		Version:     1.1,
        		Description: "Add some field",
        		Script:      "UPDATE TABLE users ADD phone_number TEXT NOT NULL;",
        	},
        	{
        		Version:        2,
        		Description:    "Migracija z post funkcijo",
        		Script:         addBirthdayAndAgeField,
        		PostScriptFunc: calculateAge,
        	},
        }

        func calculateAge(db *sql.DB) (err error) {
        
            //Preračunaj starosti uporabnikov
        
        }
   ```