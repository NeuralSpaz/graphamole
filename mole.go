package mole

//Mole struct tags
// three options
// local:
// global:
// label:
// type T struct {
// 	Name       string   `mole:"name, label"`
// 	Genre      []string `mole:"genre, global"`
// 	Rating     float64  `mole:"rating, label"`
// 	Nsfw       bool     `mole:"nsfw, global"`
// 	CoverImage struct {
// 		Original string `mole:"original, label"`
// 		Small    string `mole:"small label"`
// 	} `mole:"cover, local"`
// }

// var AwesomeShow T

// ╔═════════════╗
// ║   (local)   ╠═ <original> ═ T.CoverImage.Original string
// ║    cover    ╠═ <small> ════ T.CoverImage.Small    string
// ╚══════╦══════╝
//        ║
//     <cover>
//        ║
// ╔══════╩══════╗
// ║             ╠══ <name> ════ T.Name   string
// ║ AwesomeShow ╠══ <rating> ══ T.Rating float64               ╔╩╩╩╩╩╩╩╩╗
// ║             ╠══ <nsfw> ════(only if true)══════════════════╣ nsfw   ║
// ║             ╠══ <genre> ══════════════════════════════════╗╚╦╦╦╦╦╦╦╦╝
// ║             ╠══ <genre> ══════════════════╗               ║
// ║             ╠══ <genre> ══╗               ║               ║
// ╚═════════════╝             ║               ║               ║
//  		    			   ║               ║               ║
//  					  ╔════╩═════╗    ╔════╩══════╗   ╔════╩══════╗
//  					  ║ ACTION   ║    ║ADVENTURE  ║   ║ COMEDY    ║
//  					  ║T.Genre[0]║    ║T.Genre[1] ║   ║T.Genre[2] ║
//  					  ╚╦╦╦╦╦╦╦╦╦╦╝    ╚╦╦╦╦╦╦╦╦╦╦╦╝   ╚╦╦╦╦╦╦╦╦╦╦╦╝

// mutation {
// 	set {
// 	   _:class <student> _:x .
// 	   _:class <student> _:y .
// 	   _:class <name> "awesome class"^^<xs:string> .
// 	   _:x <name> "Alice"^^<xs:string> .
// 	   _:x <planet> "Mars"^^<xs:string> .
// 	   _:x <friend> _:y .
// 	   _:y <name> "Bob"^^<xs:string> .
// 	   _:y <age> "18"^^<xs:int> .
// 	}
//    }

// ordering
// if connection
// node predicate node
// if scalar
// node predicate scalar
// if scalar with facet
// node predicate scalar facet([]scalar)

// When filtering by applying a function, Dgraph uses the index to make the search through a potentially large dataset efficient.

// All scalar types can be indexed.

// Types int, float, bool and geo have only a default index each: with tokenizers named int, float, bool and geo.

// Types string and dateTime have a number of indices.
// String Indices

// The indices available for strings are as follows.
// Index name / Tokenizer 	Purpose 	Dgraph functions
// exact 	matching of entire value 	eq, le, ge, gt, lt
// hash 	matching of entire value, useful when the values are large in size 	eq
// term 	matching of terms/words 	eq, allofterms, anyofterms
// fulltext 	matching with language specific stemming and stopwords 	eq, alloftext, anyoftext
// trigram 	regular expressions matching 	regexp
// DateTime Indices

// The indices available for dateTime are as follows.
// Index name / Tokenizer 	Part of date indexed
// year 	index on year (default)
// month 	index on year and month
// day 	index on year, month and day
// hour 	index on year, month, day and hour

// The choices of dateTime index allow selecting the precision of the index. Applications, such as the movies examples in these docs, that require searching over dates but have relatively few nodes per year may prefer the year tokenizer; applications that are dependent on fine grained date searches, such as real-time sensor readings, may prefer the hour index.

// All the dateTime indices are sortable.
var (
	stringIndexExact    = []byte("exact")
	stringIndexHash     = []byte("hash")
	stringIndexTerm     = []byte("term")
	stringIndexFulltext = []byte("fulltext")
	stringIndexTrigram  = []byte("trigram")
)
