package mole

import (
	"fmt"
	"log"
	"reflect"
	"sync"
)

type typeInfo struct {
	ref    *fieldInfo
	fields []fieldInfo
}

type fieldInfo struct {
	idx       []int
	tagname   string
	tagopts   optionTags
	flags     fieldFlags
	parent    reflect.Type
	parenName string
}

type fieldFlags int

const (
	fnode fieldFlags = 1 << iota
	fedge
	ffacet
	fOmitEmpty
)

var tinfoMap sync.Map // map[reflect.Type]*typeInfo

func getTypeInfo(typ reflect.Type) (*typeInfo, error) {
	if ti, ok := tinfoMap.Load(typ); ok {
		return ti.(*typeInfo), nil
	}

	tinfo := &typeInfo{}
	// fmt.Println("HEADER: ", typ)
	if typ.Kind() == reflect.Struct {
		n := typ.NumField()
		// tinfo.ref = &fieldInfo{flags: fnode}

		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if (f.PkgPath != "" && !f.Anonymous) || f.Tag.Get(StructTagString) == "-" {
				continue // Unexported field, not to be inculded in type map
			}

			t := f.Type
			// if pointer get elem
			if t.Kind() == reflect.Ptr {
				t = t.Elem()
			}

			finfo, err := structFieldInfo(t, &f)
			if err != nil {
				return nil, err
			}

			finfo.parent = typ
			finfo.parenName = typ.Name()
			tinfo.fields = append(tinfo.fields, *finfo)

			fmt.Println("HEADER: ", finfo)
			if t.Kind() == reflect.Struct {

				child, err := getTypeInfo(t)
				if err != nil {
					return nil, err
				}
				fmt.Println("CHILDHEADER: ", child)
				for _, chfinfo := range child.fields {
					fmt.Println("HEADER: ", chfinfo)
					// 	finfo.idx = append([]int{i}, finfo.idx...)
					// finfo.parent = typ
					// finfo.parenName = typ.Name()
					// 	child.fields = append(child.fields, finfo)
				}
				continue
			}

			// fmt.Println("HEADER: ", finfo)

		}
	}

	ti, _ := tinfoMap.LoadOrStore(typ, tinfo)
	return ti.(*typeInfo), nil
}

// structFieldInfo builds and returns a fieldInfo for f.
func structFieldInfo(typ reflect.Type, f *reflect.StructField) (*fieldInfo, error) {
	finfo := &fieldInfo{idx: f.Index}

	// TODO use tags to get if facet
	switch typ.Kind() {
	case reflect.Struct:
		finfo.flags = fnode
	default:
		finfo.flags = fedge
	}

	tag, opts, err := getTag(*f)
	switch err {
	case ErrEmptyStructTag:
		finfo.tagname = f.Name
		return finfo, nil
	case ErrEmptyStructOptions:
		finfo.tagname = tag
		return finfo, nil
	case ErrInvaidCharInStructTag:
		log.Printf("Using Struct name as: %v", err)
		finfo.tagname = f.Name
		return finfo, nil
	case nil:
		finfo.tagname = tag
		finfo.tagopts = opts
	default:
		log.Printf("Using Struct name as: %v", err)
		finfo.tagname = f.Name
		return finfo, nil
	}

	return finfo, nil

}

func (finfo *fieldInfo) value(v reflect.Value) reflect.Value {
	for i, x := range finfo.idx {
		if i > 0 {
			t := v.Type()
			if t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct {
				if v.IsNil() {
					v.Set(reflect.New(v.Type().Elem()))
				}
				v = v.Elem()
			}
		}
		v = v.Field(x)
	}
	return v
}
