package imdb

import (
	"fmt"
	"testing"
)

func TestTitle(t *testing.T) {
	_, err := NewTitle(client, "wrong")
	if err != ErrInvalidID {
		t.Errorf("NewTitle(wrong) = %v; want ErrInvalidId", err)
	}

	for _, tt := range []struct {
		ID   string
		want Title
	}{
		{
			ID: "tt0073845",
			want: Title{
				ID:       "tt0073845",
				URL:      "https://www.imdb.com/title/tt0073845",
				Name:     "L'uomo che sfidò l'organizzazione",
				Year:     1975,
				Rating:   "5.2",
				Duration: "87m",
				Directors: []Name{
					Name{ID: "nm0340894", URL: "https://www.imdb.com/name/nm0340894", FullName: "Sergio Grieco"},
				},
				Writers: []Name{
					Name{ID: "nm0340894", URL: "https://www.imdb.com/name/nm0340894", FullName: "Sergio Grieco"},
				},
				Actors: []Name{
					Name{ID: "nm0001886", URL: "https://www.imdb.com/name/nm0001886", FullName: "Howard Ross"},
					Name{ID: "nm0775810", URL: "https://www.imdb.com/name/nm0775810", FullName: "Karin Schubert"},
					Name{ID: "nm0237835", URL: "https://www.imdb.com/name/nm0237835", FullName: "Jean-Claude Dreyfus"},
					Name{ID: "nm0674183", URL: "https://www.imdb.com/name/nm0674183", FullName: "Nadine Perles"},
					Name{ID: "nm0197623", URL: "https://www.imdb.com/name/nm0197623", FullName: "Alberto Dalbés"},
					Name{ID: "nm0130952", URL: "https://www.imdb.com/name/nm0130952", FullName: "José Calvo"},
					Name{ID: "nm0000963", URL: "https://www.imdb.com/name/nm0000963", FullName: "Stephen Boyd"},
					Name{ID: "nm0157826", URL: "https://www.imdb.com/name/nm0157826", FullName: "José Luis Chinchilla"},
					Name{ID: "nm0054079", URL: "https://www.imdb.com/name/nm0054079", FullName: "Pietro Torrisi"},
					Name{ID: "nm0877114", URL: "https://www.imdb.com/name/nm0877114", FullName: "Luciana Turina"},
					Name{ID: "nm0328179", URL: "https://www.imdb.com/name/nm0328179", FullName: "Gaspar 'Indio' González"},
				},
				Genres:        []string{"Crime", "Drama"},
				Languages:     []string{"Italian"},
				Nationalities: []string{"Italy", "France", "Spain"},
				Poster:        Media{ID: "rm143003904", TitleID: "tt0073845", URL: "https://www.imdb.com/title/tt0073845/mediaviewer/rm143003904", ContentURL: "https://ia.media-imdb.com/images/M/MV5BOWVmYzRiZTAtOTRkMS00MDQxLTljNTItM2E2NjhmZGFmZDNkXkEyXkFqcGdeQXVyNTI0NjExNjA@._V1_UX182_CR0,0,182,268_AL_.jpg"},
				AKA:           []string{"Antimetopos me tin mafia", "El hombre que desafió a la organización", "L'homme qui défia l'organisation", "One Man Against the Organization"},
			},
		},
		{
			ID: "tt0437804",
			want: Title{
				ID:       "tt0437803", // ID redirect.
				URL:      "https://www.imdb.com/title/tt0437803",
				Name:     "Alien Siege",
				Type:     "TV Movie",
				Year:     2005,
				Rating:   "3.6",
				Duration: "90m",
				Directors: []Name{
					Name{ID: "nm0821100", URL: "https://www.imdb.com/name/nm0821100", FullName: "Robert Stadd"},
				},
				Writers: []Name{
					Name{ID: "nm1305705", URL: "https://www.imdb.com/name/nm1305705", FullName: "Bill Lundy"},
					Name{ID: "nm0757507", URL: "https://www.imdb.com/name/nm0757507", FullName: "Paul Salamoff"},
				},
				Actors: []Name{
					Name{ID: "nm0424635", URL: "https://www.imdb.com/name/nm0424635", FullName: "Brad Johnson"},
					Name{ID: "nm1658541", URL: "https://www.imdb.com/name/nm1658541", FullName: "Erin Ross"},
					Name{ID: "nm0250169", URL: "https://www.imdb.com/name/nm0250169", FullName: "Lilas Lane"},
					Name{ID: "nm0003405", URL: "https://www.imdb.com/name/nm0003405", FullName: "Nathan Anderson"},
					Name{ID: "nm1108894", URL: "https://www.imdb.com/name/nm1108894", FullName: "Michael Cory Davis"},
					Name{ID: "nm1855195", URL: "https://www.imdb.com/name/nm1855195", FullName: "Gregor Paslawsky"},
					Name{ID: "nm0048844", URL: "https://www.imdb.com/name/nm0048844", FullName: "Ray Baker"},
					Name{ID: "nm0001835", URL: "https://www.imdb.com/name/nm0001835", FullName: "Carl Weathers"},
					Name{ID: "nm1197031", URL: "https://www.imdb.com/name/nm1197031", FullName: "George Zlatarev"},
					Name{ID: "nm1791963", URL: "https://www.imdb.com/name/nm1791963", FullName: "Vladimir Nikolov"},
					Name{ID: "nm1284482", URL: "https://www.imdb.com/name/nm1284482", FullName: "Atanas Srebrev"},
					Name{ID: "nm1838387", URL: "https://www.imdb.com/name/nm1838387", FullName: "Carl Ed"},
					Name{ID: "nm1838452", URL: "https://www.imdb.com/name/nm1838452", FullName: "Ivan Ivanov"},
					Name{ID: "nm1854456", URL: "https://www.imdb.com/name/nm1854456", FullName: "Zarkov Binev"},
					Name{ID: "nm1208093", URL: "https://www.imdb.com/name/nm1208093", FullName: "Julian Vergov"},
				},
				Genres:        []string{"Sci-Fi"},
				Languages:     []string{"English"},
				Nationalities: []string{"USA"},
				Description:   "Earth is attacked by the Kulkus, a hostile breed infected by a lethal virus and needing human blood to develop an antidote. Earth's governments negotiate peace terms with the Kulku ambassador, giving eight million humans shared between the nations to the invaders and in return they would spare the planet. When Heather Chase, the daughter of the scientist Stephen Chase, is one of the selected, her...",
				Poster:        Media{ID: "rm3287190272", TitleID: "tt0437803", URL: "https://www.imdb.com/title/tt0437803/mediaviewer/rm3287190272", ContentURL: "https://ia.media-imdb.com/images/M/MV5BMTk1MTA4NDMwMF5BMl5BanBnXkFtZTcwNTM2MTIzMg@@._V1_UY268_CR4,0,182,268_AL_.jpg"},
				AKA:           []string{"A Föld ostroma", "Alien Blood", "Alien Siege", "Alien Siege - Tod aus dem All", "Etat de siège", "O Perigo Alienígena", "Obca krew"},
			},
		},
		{
			ID: "tt1179034",
			want: Title{
				ID:       "tt1179034",
				URL:      "https://www.imdb.com/title/tt1179034",
				Name:     "From Paris with Love",
				Year:     2010,
				Rating:   "6.5",
				Duration: "92m",
				Directors: []Name{
					Name{ID: "nm0603628", URL: "https://www.imdb.com/name/nm0603628", FullName: "Pierre Morel"},
				},
				Writers: []Name{
					Name{ID: "nm0367867", URL: "https://www.imdb.com/name/nm0367867", FullName: "Adi Hasak"},
					Name{ID: "nm0000108", URL: "https://www.imdb.com/name/nm0000108", FullName: "Luc Besson"},
				},
				Actors: []Name{
					Name{ID: "nm0000237", URL: "https://www.imdb.com/name/nm0000237", FullName: "John Travolta"},
					Name{ID: "nm0001667", URL: "https://www.imdb.com/name/nm0001667", FullName: "Jonathan Rhys Meyers"},
					Name{ID: "nm0810738", URL: "https://www.imdb.com/name/nm0810738", FullName: "Kasia Smutniak"},
					Name{ID: "nm0243983", URL: "https://www.imdb.com/name/nm0243983", FullName: "Richard Durden"},
					Name{ID: "nm0082897", URL: "https://www.imdb.com/name/nm0082897", FullName: "Bing Yin"},
					Name{ID: "nm2362244", URL: "https://www.imdb.com/name/nm2362244", FullName: "Amber Rose Revah"},
					Name{ID: "nm1101713", URL: "https://www.imdb.com/name/nm1101713", FullName: "Eric Godon"},
					Name{ID: "nm2394847", URL: "https://www.imdb.com/name/nm2394847", FullName: "François Bredon"},
					Name{ID: "nm1282803", URL: "https://www.imdb.com/name/nm1282803", FullName: "Chems Dahmani"},
					Name{ID: "nm1410527", URL: "https://www.imdb.com/name/nm1410527", FullName: "Sami Darr"},
					Name{ID: "nm3783445", URL: "https://www.imdb.com/name/nm3783445", FullName: "Julien Hagnery"},
					Name{ID: "nm0830599", URL: "https://www.imdb.com/name/nm0830599", FullName: "Mostéfa Stiti"},
					Name{ID: "nm3266168", URL: "https://www.imdb.com/name/nm3266168", FullName: "Rebecca Dayan"},
					Name{ID: "nm2457148", URL: "https://www.imdb.com/name/nm2457148", FullName: "Michaël Vander-Meiren"},
					Name{ID: "nm1359275", URL: "https://www.imdb.com/name/nm1359275", FullName: "Didier Constant"},
				},
				Genres:        []string{"Action", "Crime", "Thriller"},
				Languages:     []string{"English", "French", "Mandarin", "German"},
				Nationalities: []string{"France"},
				Description:   "In Paris, a young employee in the office of the US Ambassador hooks up with an American spy looking to stop a terrorist attack in the city.",
				Poster:        Media{ID: "rm30481408", TitleID: "tt1179034", URL: "https://www.imdb.com/title/tt1179034/mediaviewer/rm30481408", ContentURL: "https://ia.media-imdb.com/images/M/MV5BODAwMDFjNjktMWY2Mi00MmVhLWI0MjYtNzg4OTI0NzA5YzBjXkEyXkFqcGdeQXVyNTIzOTk5ODM@._V1_UX182_CR0,0,182,268_AL_.jpg"},
				AKA:           []string{"Apo to Parisi me agapi", "Armastusega Pariisist", "Bez soucitu", "Bons baisers de Paris", "De Paris com Amor", "Desde París con amor", "Din Paris, cu dragoste", "Dupla Implacável", "From Paris with Love", "Iz Pariza s ljubavlju", "Iz Pariza z ljubeznijo", "Linkejimai is Paryziaus", "MiParis be'ahava", "Paris'ten sevgilerle", "París en la Mira", "París en la mira", "Pozdrowienia z Paryza", "Párizsból szeretettel", "Sangre y amor en París", "Από το Παρίσι με αγάπη", "З Парижу з любов'ю", "Из Парижа с любовью", "От Париж с любов"},
			},
		},
		{
			ID: "tt0133093",
			want: Title{
				ID:       "tt0133093",
				URL:      "https://www.imdb.com/title/tt0133093",
				Name:     "The Matrix",
				Year:     1999,
				Rating:   "8.7",
				Duration: "136m",
				Directors: []Name{
					Name{ID: "nm0905154", URL: "https://www.imdb.com/name/nm0905154", FullName: "Lana Wachowski"},
					Name{ID: "nm0905152", URL: "https://www.imdb.com/name/nm0905152", FullName: "Lilly Wachowski"},
				},
				Writers: []Name{
					Name{ID: "nm0905152", URL: "https://www.imdb.com/name/nm0905152", FullName: "Lilly Wachowski"},
					Name{ID: "nm0905154", URL: "https://www.imdb.com/name/nm0905154", FullName: "Lana Wachowski"},
				},
				Actors: []Name{
					Name{ID: "nm0000206", URL: "https://www.imdb.com/name/nm0000206", FullName: "Keanu Reeves"},
					Name{ID: "nm0000401", URL: "https://www.imdb.com/name/nm0000401", FullName: "Laurence Fishburne"},
					Name{ID: "nm0005251", URL: "https://www.imdb.com/name/nm0005251", FullName: "Carrie-Anne Moss"},
					Name{ID: "nm0915989", URL: "https://www.imdb.com/name/nm0915989", FullName: "Hugo Weaving"},
					Name{ID: "nm0287825", URL: "https://www.imdb.com/name/nm0287825", FullName: "Gloria Foster"},
					Name{ID: "nm0001592", URL: "https://www.imdb.com/name/nm0001592", FullName: "Joe Pantoliano"},
					Name{ID: "nm0159059", URL: "https://www.imdb.com/name/nm0159059", FullName: "Marcus Chong"},
					Name{ID: "nm0032810", URL: "https://www.imdb.com/name/nm0032810", FullName: "Julian Arahanga"},
					Name{ID: "nm0233391", URL: "https://www.imdb.com/name/nm0233391", FullName: "Matt Doran"},
					Name{ID: "nm0565883", URL: "https://www.imdb.com/name/nm0565883", FullName: "Belinda McClory"},
					Name{ID: "nm0662562", URL: "https://www.imdb.com/name/nm0662562", FullName: "Anthony Ray Parker"},
					Name{ID: "nm0323822", URL: "https://www.imdb.com/name/nm0323822", FullName: "Paul Goddard"},
					Name{ID: "nm0853079", URL: "https://www.imdb.com/name/nm0853079", FullName: "Robert Taylor"},
					Name{ID: "nm0040058", URL: "https://www.imdb.com/name/nm0040058", FullName: "David Aston"},
					Name{ID: "nm0336802", URL: "https://www.imdb.com/name/nm0336802", FullName: "Marc Aden Gray"},
				},
				Genres:        []string{"Action", "Sci-Fi"},
				Languages:     []string{"English"},
				Nationalities: []string{"USA"},
				Description:   "A computer hacker learns from mysterious rebels about the true nature of his reality and his role in the war against its controllers.",
				Poster:        Media{ID: "rm525547776", TitleID: "tt0133093", URL: "https://www.imdb.com/title/tt0133093/mediaviewer/rm525547776", ContentURL: "https://ia.media-imdb.com/images/M/MV5BNzQzOTk3OTAtNDQ0Zi00ZTVkLWI0MTEtMDllZjNkYzNjNTc4L2ltYWdlXkEyXkFqcGdeQXVyNjU0OTQ0OTY@._V1_UX182_CR0,0,182,268_AL_.jpg"},
				AKA:           []string{"Fylkið", "La matrice", "La matriz", "Maatriks", "Matorikkusu", "Matrica", "Matriks", "Matrikss", "Matrix", "Mátrix", "The Matrix", "Матрикс", "Матрица", "Матрицата", "Матриця"},
			},
		},
		{
			ID: "tt0291830",
			want: Title{
				ID:       "tt0291830",
				URL:      "https://www.imdb.com/title/tt0291830",
				Name:     "Corps à coeur",
				Year:     1981,
				Duration: "26m",
				Directors: []Name{
					Name{ID: "nm1029036", URL: "https://www.imdb.com/name/nm1029036", FullName: "Jean-François Després"},
					Name{ID: "nm1003289", URL: "https://www.imdb.com/name/nm1003289", FullName: "Alain Godon"},
				},
				Genres:        []string{"Documentary", "Short"},
				Nationalities: []string{"Canada"},
			},
		},
		{
			ID: "tt1965639",
			want: Title{
				ID:            "tt1965639",
				URL:           "https://www.imdb.com/title/tt1965639",
				Name:          "El clima y las cuatro estaciones",
				Type:          "TV Series",
				Year:          1994,
				Genres:        []string{"Documentary"},
				Languages:     []string{"Spanish"},
				Nationalities: []string{"Spain"},
			},
		},
		{
			ID: "tt0086677",
			want: Title{
				ID:       "tt0086677",
				URL:      "https://www.imdb.com/title/tt0086677",
				Name:     "Brothers",
				Type:     "TV Series",
				Year:     1984,
				Rating:   "8.0",
				Duration: "22m",
				Writers: []Name{
					Name{ID: "nm0515953", URL: "https://www.imdb.com/name/nm0515953", FullName: "David Lloyd"},
					Name{ID: "nm0031246", URL: "https://www.imdb.com/name/nm0031246", FullName: "Greg Antonacci"},
				},
				Actors: []Name{
					Name{ID: "nm0907107", URL: "https://www.imdb.com/name/nm0907107", FullName: "Robert Walden"},
					Name{ID: "nm0716618", URL: "https://www.imdb.com/name/nm0716618", FullName: "Paul Regina"},
					Name{ID: "nm0535933", URL: "https://www.imdb.com/name/nm0535933", FullName: "Brandon Maggart"},
					Name{ID: "nm0604699", URL: "https://www.imdb.com/name/nm0604699", FullName: "Hallie Todd"},
					Name{ID: "nm0533383", URL: "https://www.imdb.com/name/nm0533383", FullName: "Philip Charles MacKenzie"},
					Name{ID: "nm0726939", URL: "https://www.imdb.com/name/nm0726939", FullName: "Robin Riker"},
					Name{ID: "nm0664289", URL: "https://www.imdb.com/name/nm0664289", FullName: "Mary Ann Pascal"},
				},
				Genres:        []string{"Comedy"},
				Languages:     []string{"English"},
				Nationalities: []string{"USA"},
				Description:   "Two conservative men support their younger brother when he comes out as gay, and help him navigate being openly homosexual in 1980s Philadelphia.",
				Poster:        Media{ID: "rm1328617472", TitleID: "tt0086677", URL: "https://www.imdb.com/title/tt0086677/mediaviewer/rm1328617472", ContentURL: "https://ia.media-imdb.com/images/M/MV5BNDM3N2IxY2UtMDc4MS00ZDM5LWJjODUtMjQ5ODg2YWIxMTE4XkEyXkFqcGdeQXVyMDYxMTUwNg@@._V1_UY268_CR87,0,182,268_AL_.jpg"},
				AKA:           []string{"Braća", "Brothers", "Unter Brüdern"},
			},
		},
		{
			ID: "tt1371159",
			want: Title{
				ID:     "tt1371159",
				URL:    "https://www.imdb.com/title/tt1371159",
				Name:   "Iron Man 2",
				Type:   "Video Game",
				Year:   2010,
				Rating: "6.2",
				Directors: []Name{
					Name{ID: "nm4157448", URL: "https://www.imdb.com/name/nm4157448", FullName: "Michael McCormick"},
					Name{ID: "nm4157551", URL: "https://www.imdb.com/name/nm4157551", FullName: "Robert Taylor"},
				},
				Writers: []Name{
					Name{ID: "nm3881781", URL: "https://www.imdb.com/name/nm3881781", FullName: "Matt Fraction"},
				},
				Actors: []Name{
					Name{ID: "nm0000332", URL: "https://www.imdb.com/name/nm0000332", FullName: "Don Cheadle"},
					Name{ID: "nm0519680", URL: "https://www.imdb.com/name/nm0519680", FullName: "Eric Loomis"},
					Name{ID: "nm0000168", URL: "https://www.imdb.com/name/nm0000168", FullName: "Samuel L. Jackson"},
					Name{ID: "nm0482851", URL: "https://www.imdb.com/name/nm0482851", FullName: "Phil LaMarr"},
					Name{ID: "nm0072829", URL: "https://www.imdb.com/name/nm0072829", FullName: "John Eric Bentley"},
					Name{ID: "nm0005243", URL: "https://www.imdb.com/name/nm0005243", FullName: "Meredith Monroe"},
					Name{ID: "nm0133047", URL: "https://www.imdb.com/name/nm0133047", FullName: "Catherine Campion"},
					Name{ID: "nm0224929", URL: "https://www.imdb.com/name/nm0224929", FullName: "Dimitri Diatchenko"},
					Name{ID: "nm0149677", URL: "https://www.imdb.com/name/nm0149677", FullName: "Andrew Chaikin"},
					Name{ID: "nm0101752", URL: "https://www.imdb.com/name/nm0101752", FullName: "Doug Boyd"},
					Name{ID: "nm0946382", URL: "https://www.imdb.com/name/nm0946382", FullName: "Cedric Yarbrough"},
					Name{ID: "nm0217246", URL: "https://www.imdb.com/name/nm0217246", FullName: "Denny Delk"},
					Name{ID: "nm0413996", URL: "https://www.imdb.com/name/nm0413996", FullName: "Roger Jackson"},
					Name{ID: "nm2406242", URL: "https://www.imdb.com/name/nm2406242", FullName: "Adam Harrington"},
					Name{ID: "nm3894439", URL: "https://www.imdb.com/name/nm3894439", FullName: "Ariel Goldberg"},
				},
				Genres:        []string{"Action", "Adventure", "Sci-Fi"},
				Languages:     []string{"English"},
				Nationalities: []string{"USA"},
				Description:   "Video game based upon the film of the same name.",
				Poster:        Media{ID: "rm1884363776", TitleID: "tt1371159", URL: "https://www.imdb.com/title/tt1371159/mediaviewer/rm1884363776", ContentURL: "https://ia.media-imdb.com/images/M/MV5BMWNkYWIyZjctMGFmZi00NGYyLWIzY2QtNGEwNDQ2MDgwOWZiXkEyXkFqcGdeQXVyNTk5Nzg0MDE@._V1_UY268_CR17,0,182,268_AL_.jpg"},
			},
		},
		{
			ID: "tt3038708",
			want: Title{
				ID:     "tt3038708",
				URL:    "https://www.imdb.com/title/tt3038708",
				Name:   "Iron Sky: The Coming Race",
				Year:   2018,
				Rating: "8.2",
				Directors: []Name{
					Name{ID: "nm1993322", URL: "https://www.imdb.com/name/nm1993322", FullName: "Timo Vuorensola"},
				},
				Writers: []Name{
					Name{ID: "nm0120272", URL: "https://www.imdb.com/name/nm0120272", FullName: "Edward George Bulwer-Lytton"},
					Name{ID: "nm0615914", URL: "https://www.imdb.com/name/nm0615914", FullName: "Dalan Musson"},
				},
				Actors: []Name{
					Name{ID: "nm4468181", URL: "https://www.imdb.com/name/nm4468181", FullName: "Lara Rossi"},
					Name{ID: "nm0122095", URL: "https://www.imdb.com/name/nm0122095", FullName: "Vladimir Burlakov"},
					Name{ID: "nm7610067", URL: "https://www.imdb.com/name/nm7610067", FullName: "Kit Dale"},
					Name{ID: "nm0338381", URL: "https://www.imdb.com/name/nm0338381", FullName: "Tom Green"},
					Name{ID: "nm1087430", URL: "https://www.imdb.com/name/nm1087430", FullName: "Julia Dietze"},
					Name{ID: "nm3893266", URL: "https://www.imdb.com/name/nm3893266", FullName: "Edward Judge"},
					Name{ID: "nm2953208", URL: "https://www.imdb.com/name/nm2953208", FullName: "Martin Swabey"},
					Name{ID: "nm3038143", URL: "https://www.imdb.com/name/nm3038143", FullName: "Emily Atack"},
					Name{ID: "nm0001424", URL: "https://www.imdb.com/name/nm0001424", FullName: "Udo Kier"},
					Name{ID: "nm1712811", URL: "https://www.imdb.com/name/nm1712811", FullName: "John Flanders"},
					Name{ID: "nm1094184", URL: "https://www.imdb.com/name/nm1094184", FullName: "Stephanie Paul"},
					Name{ID: "nm3049862", URL: "https://www.imdb.com/name/nm3049862", FullName: "James Quinn"},
					Name{ID: "nm1385437", URL: "https://www.imdb.com/name/nm1385437", FullName: "Jukka Hilden"},
					Name{ID: "nm3136958", URL: "https://www.imdb.com/name/nm3136958", FullName: "Lloyd Li"},
					Name{ID: "nm1295552", URL: "https://www.imdb.com/name/nm1295552", FullName: "Amanda Wolzak"},
				},
				Genres:        []string{"Action", "Adventure", "Comedy", "Sci-Fi"},
				Languages:     []string{"English"},
				Nationalities: []string{"Finland", "Germany", "Belgium"},
				Description:   "A follow-up to the film Iron Sky (2012) in which Nazis plan to take over the world after lying dormant in a secret military base on the moon.",
				Poster:        Media{ID: "rm68886272", TitleID: "tt3038708", URL: "https://www.imdb.com/title/tt3038708/mediaviewer/rm68886272", ContentURL: "https://ia.media-imdb.com/images/M/MV5BZmViNzE1N2QtOTU3NS00NmUzLTg4MWEtODljNDBkYjMyOGZhXkEyXkFqcGdeQXVyNjIwMTM4Ng@@._V1_UY268_CR3,0,182,268_AL_.jpg"},
				AKA:           []string{"Iron Sky 2"},
			},
		},
		{
			ID: "tt0423866",
			want: Title{
				ID:       "tt0423866",
				URL:      "https://www.imdb.com/title/tt0423866",
				Name:     "3-Iron",
				OrigName: "Bin-jip",
				Year:     2004,
				Rating:   "8.1",
				Duration: "88m",
				Directors: []Name{
					Name{ID: "nm1104118", URL: "https://www.imdb.com/name/nm1104118", FullName: "Ki-duk Kim"},
				},
				Writers: []Name{
					Name{ID: "nm1104118", URL: "https://www.imdb.com/name/nm1104118", FullName: "Ki-duk Kim"},
				},
				Actors: []Name{
					{ID: "nm1165901", URL: "https://www.imdb.com/name/nm1165901", FullName: "Seung-Yun Lee"},
					{ID: "nm1030819", URL: "https://www.imdb.com/name/nm1030819", FullName: "Hee Jae"},
					{ID: "nm1891528", URL: "https://www.imdb.com/name/nm1891528", FullName: "Hyuk-ho Kwon"},
					{ID: "nm1873389", URL: "https://www.imdb.com/name/nm1873389", FullName: "Jeong-ho Choi"},
					{ID: "nm1450477", URL: "https://www.imdb.com/name/nm1450477", FullName: "Ju-seok Lee"},
					{ID: "nm1064338", URL: "https://www.imdb.com/name/nm1064338", FullName: "Mi-suk Lee"},
					{ID: "nm1875737", URL: "https://www.imdb.com/name/nm1875737", FullName: "Sung-hyuk Moon"},
					{ID: "nm1280733", URL: "https://www.imdb.com/name/nm1280733", FullName: "Ji-a Park"},
					{ID: "nm1876659", URL: "https://www.imdb.com/name/nm1876659", FullName: "Jae-yong Jang"},
					{ID: "nm1872776", URL: "https://www.imdb.com/name/nm1872776", FullName: "Dah-hae Lee"},
					{ID: "nm1873110", URL: "https://www.imdb.com/name/nm1873110", FullName: "Han Kim"},
					{ID: "nm1833990", URL: "https://www.imdb.com/name/nm1833990", FullName: "Se-jin Park"},
					{ID: "nm1873533", URL: "https://www.imdb.com/name/nm1873533", FullName: "Dong-jin Park"},
					{ID: "nm1048008", URL: "https://www.imdb.com/name/nm1048008", FullName: "Jong-su Lee"},
					{ID: "nm1879243", URL: "https://www.imdb.com/name/nm1879243", FullName: "Ui-soo Lee"},
				},
				Genres:        []string{"Crime", "Drama", "Romance"},
				Languages:     []string{"Korean"},
				Nationalities: []string{"South Korea", "Japan"},
				Description:   "A transient young man breaks into empty homes to partake of the vacationing residents' lives for a few days.",
				Poster:        Media{ID: "rm880057600", TitleID: "tt0423866", URL: "https://www.imdb.com/title/tt0423866/mediaviewer/rm880057600", ContentURL: "https://ia.media-imdb.com/images/M/MV5BMTM1ODIwNzM5OV5BMl5BanBnXkFtZTcwNjk5MDkyMQ@@._V1_UX182_CR0,0,182,268_AL_.jpg"},
				AKA:           []string{"3-Iron", "3-iron", "Bin jib", "Bin-Jip - Der Schattenmann", "Bin-Jip - Leere Häuser", "Bin-jip", "Bin-jip - tomme huse", "Bos ev", "Bosh Evler", "Casa Vazia", "El espíritu de la pasión", "Empty Houses", "Ferro 3", "Ferro 3 - La casa vuota", "Hierro 3", "Hierro-3", "Järn 3:an", "Järntrean", "Khanehaye Khali", "Lehargish B'Bayit", "Locataires", "Lopakodó lelkek", "Menaj in trei", "Olomonahoi mazi", "Provalnik", "Pusty dom", "Rautakolmonen", "Tomme hus", "Tusti namai", "Utsusemi", "Ολομόναχοι Μαζί", "Пустой дом", "Стик Nо 3"},
			},
		},
	} {
		got, err := NewTitle(client, tt.ID)
		if err != nil {
			t.Errorf("NewTitle(%s) error: %v", tt.ID, err)
		} else {
			if err := diffStruct(tt.want, *got); err != nil {
				t.Errorf("NewTitle(%s): %v", tt.ID, err)
			}
		}
	}
}

func TestSingleTitle(t *testing.T) {
	g, err := NewTitle(client, "tt0058453")
	if err != nil {
		t.Fatalf("ERROR: %v", err)
	}
	fmt.Printf("DATA: %+v", g)
}
