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
