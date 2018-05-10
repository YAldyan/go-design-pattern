package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
	implement interface HTTP.handler
	karena ada method ServeHTTP
*/
type MyServer struct{}

func (m *MyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Decorator!")
}

/*
	Part untuk decorate, tambahkan fungsi
	logging untuk HTTP.handler
*/
type LoggerServer struct {
	// menggunakan http.handle sebagai tipe data/variabel
	Handler   http.Handler
	LogWriter io.Writer
}

func (s *LoggerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(s.LogWriter, "Request URI: %s\n", r.RequestURI)
	fmt.Fprintf(s.LogWriter, "Host: %s\n", r.Host)
	fmt.Fprintf(s.LogWriter, "Content Length: %d\n",r.ContentLength)
	fmt.Fprintf(s.LogWriter, "Method: %s\n",

	r.Method)fmt.Fprintf(s.LogWriter, "--------------------------------\n")
	s.Handler.ServeHTTP(w, r)
}

/*
	Decorate Authentication
*/
type BasicAuthMiddleware struct {
	// a handler to decorate like in the previous middlewares	
	Handler http.Handler

	User string
	Password string
}

func (s *BasicAuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	user, pass, ok := r.BasicAuth()
	if ok {
		if user == s.User && pass == s.Password {
			s.Handler.ServeHTTP(w, r)
		}else {
			fmt.Fprintf(w, "User or password incorrect\n")
		}
	}else {
		fmt.Fprintln(w, "Error trying to retrieve data from Basic auth")
	}
}

func main() {

	/*
		We have decorated MyServer with logging capabilities without 
		actually modifying it
	*/
	http.Handle("/", &LoggerServer{
	LogWriter:os.Stdout,
		Handler:&MyServer{},
	})

	fmt.Println("starting web server at http://localhost:8000/")
	log.Fatal(http.ListenAndServe(":8000", nil))

	fmt.Println("Enter the type number of server you want to launch from the following:")
	fmt.Println("1.- Plain server")
	fmt.Println("2.- Server with logging")
	fmt.Println("3.- Server with logging and authentication")

	var selection int
	fmt.Fscanf(os.Stdin, "%d", &selection)

	switch selection {
	case 1:
		mySuperServer = new(MyServer)
	case 2:
		mySuperServer = &LoggerServer{
		Handler: new(MyServer),
			LogWriter: os.Stdout,
		}
	case 3:
		var user, password string
		fmt.Println("Enter user and password separated by a space")
		fmt.Fscanf(os.Stdin, "%s %s", &user, &password)

		mySuperServer = &LoggerServer{
			Handler: &BasicAuthMiddleware{
				Handler: new(MyServer),
					User: user,
					Password: password,
			},
			LogWriter: os.Stdout,
		}
	default:
		mySuperServer = new(MyServer)
	}
}
