package mole

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func TestNestedType(t *testing.T) {
	type Item struct {
		ID    string  //`mole:"xid"`
		Name  string  //`mole:"name"`
		Value float64 //`mole:"value"`
	}

	type NestedEmbeded struct {
		NestEmbStr  string //mole:"name"`
		NestEmbInt  int    //`mole:"count"`
		NestEmbBool bool   //`mole:"enabled"`
		Items       Item   //`mole:"item"`
		// Items       []Item `mole:"items,edge"`
		Anon struct {
			AnonInt  int
			AnonStr  string
			AnonBool bool
		} `mole:"anon"`
	}

	// NOT A REAL TEST
	n := NestedEmbeded{NestEmbBool: true}
	ntyo := reflect.TypeOf(n)
	_, err := getTypeInfo(ntyo)
	if err != nil {
		log.Println(err)
	}

	tinfoMap.Range(func(key, data interface{}) bool {
		fmt.Printf("%+v        %+v\n\n", key, data)

		return true
	})

}

// type Data struct {
// 	ID             interface{} `mole:"xid"`
// 	Name           string      `mole:"name,edge"`
// 	FavoriteColour string      `mole:"favoritecolour,edge"`
// 	Strength       int         `mole:"strength,edge"`
// 	Agility        float64     `mole:"agility,edge"`
// 	Dead           bool        `mole:"dead,edge"`
// 	Friends        []struct {
// 		ID    interface{} `mole:"xid"`
// 		Since time.Time   `mole:"since,facet:friend"`
// 	} `mole:"friend,globalNode"`
// 	Items []Item `mole:"item,localNode"`
// 	Quest string `mole:"quest,edge"`
// }

// var rdf = []byte(`
// mutation {
// 	set {
// 		_:node1 <xid> "knightoftheroundtable2" .
// 		_:node1 <name> "Sir Lancelot" .
// 		_:node1 <favoritecolour> "Blue" .
// 		_:node1 <strength> 10 .
// 		_:node1 <agility> 76.3 .
// 		_:node1 <dead> true .
// 		_:node2 <xid> "knightoftheroundtable0" .
// 		_:node1 <firend> node2 (since=timeyadayada) .
// 		_:node3 <xid> "knightoftheroundtable1" .
// 		_:node1 <firend> node2 (since=timeyadayada) .
// 		_:node4 <name> "Duck" .
// 		_:node4 <value> 0.6 .
// 		_:node1 <item> node4
// 		_:node5 <name> "Sword" .
// 		_:node5 <value> 1.0 .
// 		_:node1 <item> node5
// 		_:node6 <xid> "GuinevereLetter" .
// 		_:node6 <name> "Photo Guinevere" .
// 		_:node6 <value> 2.0 .
// 		_:node1 <item> node6 .
// 		_:node1 <quest> "To Seek The Holy Grail" .
// 	}
// }
// `)

// func TestFlattern(t *testing.T) {

// 	s := Data{
// 		ID:             "knightoftheroundtable2",
// 		Name:           "Sir Lancelot",
// 		FavoriteColour: "Blue",
// 		Strength:       10,
// 		Agility:        76.3,
// 		Dead:           true,
// 		Friends: []struct {
// 			ID    interface{} `mole:"xid"`
// 			Since time.Time   `mole:"since,facet:friend"`
// 		}{
// 			{ID: "knightoftheroundtable0", Since: time.Date(516, 12, 20, 0, 0, 0, 0, time.Local)},
// 			{ID: "knightoftheroundtable1", Since: time.Date(516, 12, 20, 0, 0, 0, 0, time.Local)},
// 		},
// 		Items: []Item{
// 			{Name: "Duck", Value: 0.6},
// 			{Name: "Sword", Value: 1.0},
// 			{ID: "GuinevereLetter", Name: "Photo Guinevere", Value: 2.0},
// 		},
// 		Quest: "To Find The Holy Grail",
// 	}

// 	// fmt.Println(typMap)
// 	if err := NewEncoder(os.Stdout).Encode(s); err != nil {
// 		log.Println(err)
// 	}

// }

// mutation {
// 	set {
// 		   _:node1 <xid> "knightoftheroundtable2" .
// 		   _:node1 <name> "Sir Lancelot" .
// 		   _:node1 <favoritecolour> "Blue" .
// 		   _:node1 <strength> "10" .
// 		   _:node1 <agility> "76.3" .
// 		   _:node1 <dead> "true" .
// 		   _:node2 <xid> "knightoftheroundtable0" .
// 		   _:node2 <name> "King Arthur" .
// 		   _:node9 <name> "Coconuts" .
// 		   _:node9 <value> "200.0" .
// 		   _:node2 <item> _:node9 .
// 		   _:node10 <name> "Aquatic Cutlery" .
// 		   _:node10 <value> "5.0" .
// 		   _:node2 <item> _:node10 .
// 		   _:node1 <friend> _:node2 (since=2016-12-20T00:00:00) .
// 		   _:node3 <xid> "knightoftheroundtable1" .
// 		   _:node3 <name> "Sir Robin" .
// 		   _:node7 <name> "Minstral" .
// 		   _:node7 <value> "100.0" .
// 		   _:node3 <item> _:node7 .
// 		   _:node8 <name> "Big Shield" .
// 		   _:node8 <value> "1.0" .
// 		   _:node3 <item> _:node8 .
// 		   _:node1 <friend> _:node3 (since=2016-12-20T00:00:00) .
// 		   _:node4 <name> "Duck" .
// 		   _:node4 <value> "0.6" .
// 		   _:node1 <item> _:node4 .
// 		   _:node5 <name> "Sword" .
// 		   _:node5 <value> "1.0" .
// 		   _:node1 <item> _:node5 .
// 		   _:node6 <xid> "GuinevereLetter" .
// 		   _:node6 <name> "Guinevere Letter" .
// 		   _:node6 <value> "2.0" .
// 		   _:node1 <item> _:node6 .
// 		   _:node1 <quest> "To Seek The Holy Grail" .
// 		   _:kinghts <xid> "knights" .
// 		   _:kinghts <name> "knights" .
// 		   _:node1 <member> _:knights .
// 		   _:node2 <member> _:knights .
// 		   _:node3 <member> _:knights .
// 	   }
//    }

//    mutation{
// 	schema{
// 	  name: string@index(term) .
// 	  xid: string@index(term) .
//       friend: uid @reverse .
// 	}
//   }

//   mutation{
// 	schema{
// 	  name: string@index(term) .
// 	  xid: string@index(trigram) .
//       friend: uid @reverse .
// 	}
//   }

//    {
// 	all(func: eq(xid,"knightoftheroundtable2")){
// 	  expand(_all_)
// 	}
//   }
