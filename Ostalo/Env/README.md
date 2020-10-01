# getEnv funkcija

## Uporaba 
Nekje v projek običajno dodamo to getEnv funkcijo (najboljše če kar na dno [main.go](../../StrukturaProjekta/main.go)). 
Z njo nato dostopamo do okoljskih spremenjljivk v katerih se običajno shranjuje konfiguracija 
(podatki za dostop do baze ali ostalih stvari)

Prvi parameter je ime okoljske spremenjljivke, drugi pa privzeta vrednost če spremenjljivka ne obstaja


## Primer - [getEnv.go](getEnv.go)

```go
dbCode := getEnv("DB_PASS", "defaultGeslo")
```