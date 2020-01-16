package agows

//Route Include route informaiton
type Route struct {
	Pattern    string
	Controller ControllerInterface
	Method     int
	MethodName string
	//HandlerFunc http.HandlerFunc
	//Middlewares []Middleware
}

//Routes A type of map[string]*Route
type Routes map[string]*Route
