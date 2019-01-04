package main

import (
	"controllers/managerController"

	mgo "gopkg.in/mgo.v2"
)

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
// }

// func main() {
// 	http.HandleFunc("/", handler)
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

type Person struct {
	Name  string
	Phone string
}

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	managerController.NewManagerController(&session)

	// c := session.DB("test").C("people")
	// err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
	// 	&Person{"Cla", "+55 53 8402 8510"})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// result := Person{}
	// err = c.Find(bson.M{"name": "Ale"}).One(&result)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Phone:", result.Phone)
	// c2 := session.DB("test").C("test-map")

	// myVar := `
	// 	{
	// 		"a": "Test",
	// 		"b": 1,
	// 		"c": true,
	// 		"d": null
	// 	}
	// `
	// m := make(map[string]interface{})
	// json.Unmarshal([]byte(myVar), &m)

	// err = c2.Insert(&m)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// v := make(map[string]string)
	// v["first"] = "1"
	// v["second"] = "2"
	// v["third"] = "3"

	// err = c2.Insert(&v)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
