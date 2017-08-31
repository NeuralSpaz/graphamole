package mole

import (
	"reflect"
	"testing"
)

func TestReadTag(t *testing.T) {
	tests := []struct {
		name     string
		args     string
		want     string
		wantOpts optionTags
		wantErr  error
	}{
		{"simple", "name", "name", optionTags{}, nil},
		{"invalid", "name!", "", optionTags{}, ErrInvaidCharInStructTag},
		{"with option", "name,hi", "name", optionTags{"hi"}, nil},
		{"with invalid option", "name,hi!", "", optionTags{}, ErrInvaidCharInStructTag},
		{"with in direction", "name,<-hi", "name", optionTags{"<-hi"}, nil},
		{"with out direction", "name,hi->", "name", optionTags{"hi->"}, nil},
		{"not letter not number", "name,💖->", "", optionTags{}, ErrInvaidCharInStructTag},
		{"empty", "", "", optionTags{}, ErrEmptyStructTag},
		{"empty options", "name,", "name", optionTags{}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tag, gotOpts, err := readTag(tt.args)
			if err != tt.wantErr {
				t.Errorf("readTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tag != tt.want {
				t.Errorf("readTag() tag = %v, want %v", tag, tt.want)
			}
			if !reflect.DeepEqual(gotOpts, tt.wantOpts) {
				t.Errorf("readTag() got1 = %v, want %v", gotOpts, tt.wantOpts)
			}
		})
	}
}

func TestGetTag(t *testing.T) {

	t.Run("test struct 1", func(t *testing.T) {
		type Test1 struct {
			Field string `mole:"name,label"`
		}
		t1 := Test1{}
		t1typ := reflect.TypeOf(t1)
		sf, _ := t1typ.FieldByName("Field")
		tag, opts, err := getTag(sf)
		if tag != "name" || opts[0] != "label" || err != nil {
			t.Errorf("failed Get tag for Test1 struct, got %v, %v,%v", tag, opts, err)
		}

	})
	t.Run("test struct 2", func(t *testing.T) {
		type Test2 struct {
			Field string `json:"field" mole:"name,label"`
		}
		t2 := Test2{}
		t2typ := reflect.TypeOf(t2)
		sf, _ := t2typ.FieldByName("Field")
		tag, opts, err := getTag(sf)
		if tag != "name" || opts[0] != "label" || err != nil {
			t.Errorf("failed Get tag for Test1 struct, got %v, %v,%v", tag, opts, err)
		}

	})
	t.Run("test struct 3", func(t *testing.T) {
		type Test3 struct {
			Field string `json:"field" mole:"likes,<-global"`
		}
		t3 := Test3{}
		t3typ := reflect.TypeOf(t3)
		sf, _ := t3typ.FieldByName("Field")
		tag, opts, err := getTag(sf)
		if tag != "likes" || opts[0] != "<-global" || err != nil {
			t.Errorf("failed Get tag for Test1 struct, got %v, %v,%v", tag, opts, err)
		}

	})
}

func BenchmarkGetTag(b *testing.B) {
	type Test1 struct {
		Field string `mole:"name,label"`
	}
	var tag string
	var opts optionTags
	var err error
	t1 := Test1{}
	// t1typ := reflect.TypeOf(t1)
	// sf, _ := t1typ.FieldByName("Field")
	b.Run("teststruct1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {

			t1typ := reflect.TypeOf(t1)
			sf, _ := t1typ.FieldByName("Field")
			tag, opts, err = getTag(sf)
		}
	})

}
