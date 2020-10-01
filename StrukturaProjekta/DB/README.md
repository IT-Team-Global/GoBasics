Tukaj imamo implementirano našo komunikacijo z bazo. Tukaj dodamo datoteko [DB.go](DB/DB.go),
v kateri definiramo naše strukture podatkov in interface z funkcijami za bazo (da lahko kadarkoli menjamo tip baze in z tem ne pokvarimo ostalega dela programa)
 
Dodamo še datoteko glede na tip baze. V našem primeru je to [MariaDB.go](DB/Mariadb.go) (MySQL). V tej datoteki implementiramo 
interface iz [DB.go](DB/DB.go). Za vse z bazo VEDNO uporabljamo interface iz [DB.go](DB/DB.go) in nikoli ne uporabljamo direktno objekta za bazo 