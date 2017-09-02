package mole

import (
	"log"
	"reflect"
)

type field struct {
	name      string
	tag       bool
	typ       reflect.Type
	global    bool
	node      bool
	empty     bool
	omit      bool
	nodeID    uint
	edgeName  string
	edgeValue bool
}

func flattern(typ reflect.Type) error {

	log.Println("DBUG: ", typ.Kind())
	// is a struct
	if typ.Kind() == reflect.Struct {
		//loop over fields in struct
		n := typ.NumField()
		for i := 0; i < n; i++ {
			field := typ.Field(i)
			//  TODO: check pointer
			log.Printf("DBUG: field: %+#v kind: %+#v", field, field.Type.Kind())
			// GET TAG HERE
			// tag, opts, err := getTag(field)
			// if err != nil {
			// 	log.Println(err)
			// }
			// fmt.Println(tag, opts)
			// skip private fields
			private := field.PkgPath != ""
			if private {
				continue
			}
			// need to treat Embded types a bit differently
			embeded := field.Anonymous
			if embeded {
				embededtyp := field.Type
				// recursive if struct
				if embededtyp.Kind() == reflect.Struct {
					log.Printf("\n\nDBUG: recurse\n\n")
					flattern(embededtyp)
				}
				continue

			}
			// recursive if struct
			if field.Type.Kind() == reflect.Struct {
				log.Printf("\n\nDBUG: recurse\n\n")
				stype := field.Type
				flattern(stype)
				continue
			}

			// no idea what to do here but need to check if its a slice of struct
			if field.Type.Kind() == reflect.Slice {
				log.Printf("\n\nDBUG: it's a slice\n\n")
				if field.Type.Elem().Kind() == reflect.Struct {
					flattern(field.Type.Elem())
				}
				//then I got no Idea

			}
		}
	}
	// TODO: check slice
	// TODO: check pointer
	// TODO:  check interface ?
	return nil
}

// func fillFieldbyTag(f field, tag string, opts optionTags) field {
// 	switch tag {
// 	// TODO:: Something I'm sure of it
// 	}
// 	for _, opt := range opts {
// 		switch opt {
// 		case "global":
// 			f.global = true
// 			f.node = true
// 		case "local":
// 			f.global = false
// 			f.node = true
// 		case "-":
// 			f.omit = false
// 		case "label":
// 			f.node = false
// 			f.edgeName = tag
// 		}
// 	}
// 	return f
// }

// func getFieldFromTag(sf reflect.StructField) (field, error) {
// 	tag := sf.Tag.Get(StructTAG)
// 	tagValue, opts := parseTag(tag)
// 	if !isValidTag(tagValue) {
// 		return field{}, errors.New("invalid struct tags")
// 	}
// 	f := field{}
// 	f = fillFieldbyTag(f, tagValue, opts)
// 	return f, nil
// }
