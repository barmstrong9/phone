package main
//LEARN HOW TO INSTALL POSTGRESQL
import (
	"bytes"
	_ "github.com/lib/pq"
)
const(
	host = "localhost"
	port = 5432
	user = "Brandon Armstrong"
	password = "Skipper99!"
	dbname = ""
)

func main(){

}

func normalise(phoneN string)string{
	var buf bytes.Buffer
	for _, ch := range phoneN{
		if ch >= '0' && ch <='9'{
			buf.WriteRune(ch)
		}
	}
	return buf.String()
}

