package mole

import (
	"reflect"
	"testing"
)

func TestParseTag(t *testing.T) {
	tests := []struct {
		name     string
		tag      string
		wantTag  string
		wantOpts []string
	}{
		{name: "label", tag: "mole,label", wantTag: "mole", wantOpts: []string{"label"}},
		{name: "empty options", tag: "mole", wantTag: "mole", wantOpts: []string{}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tag, opts := parseTag(test.tag)
			if tag != test.wantTag {
				t.Errorf("wanted tag %s but got %s", test.wantTag, tag)
			}
			if len(opts) != len(test.wantOpts) {
				t.Errorf("wanted %d tag options but got %d tag options", len(test.wantOpts), len(opts))
			}
			for k := range opts {
				if opts[k] != test.wantOpts[k] {
					t.Errorf("expected tag option %s but got %s", test.wantOpts[k], opts[k])
				}
			}
		})
	}
}

func TestIsValidTag(t *testing.T) {
	tests := []struct {
		name string
		tag  string
		want bool
	}{
		{name: "empty", tag: "", want: false},
		{name: "hasSpace", tag: " mole ", want: false},
		{name: "mole", tag: "mole", want: true},
		{name: "invalidchar", tag: "mole!", want: false},
		{name: "comma", tag: "mole,something", want: true},
		{name: "notletterNumber", tag: "💖", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ok := isValidTag(test.tag)
			if ok != test.want {
				t.Errorf("wanted tag %v but got %v for %s", test.want, ok, test.tag)

			}
		})
	}
}

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
		{"withoption", "name,hi", "name", optionTags{"hi"}, nil},
		{"withinvalidoption", "name,hi!", "", optionTags{}, ErrInvaidCharInStructTag},
		{"withindirection", "name,<-hi", "name", optionTags{"<-hi"}, nil},
		{"withoutdirection", "name,hi->", "name", optionTags{"hi->"}, nil},
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
