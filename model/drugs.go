package model

type DrugsComposition struct {
	Name    string
	Synonim []Synonim
	Img     string
	Dosages []Dosage
	Effect  []string
	Alcohol bool
}

type Synonim struct {
	Country string
	Name    string
}

type Dosage struct {
	Type            string
	Age             Ages
	WeightCondition int
	MaxConsumption  string
	Cpd             int
}

type Ages struct {
	MinAge int
	MaxAge int
}

type DrugsProduct struct {
	Categories  string
	Name        string
	Desc        string
	Indication  string
	Composition string
	Dosage      string
	UseRule     string
	NoRegist    string
	Effect      string
	Manufacture string
}

type DrugsProductMini struct {
	Id   int
	Name string
}

/*
{
    "name": "Nama Komposisi Obat",
    "latin-name": "Nama Komposisi Obat dalam latin",
    "synonim": [
        {
            "country": "",
            "name": ""
        },
        {
            "country": "",
            "name": ""
        }
    ],
    "weigth": 120.99,
    "structure-img": "/url/sss/asasas",
    "clearance": "",
    "dosage": [
        {
            "type": "children",
            "age": {
                "min": 12,
                "max": 20
            },
            "weight-condition": 10,
            "max-consumption": "60 mg",
            "cpd": 3
        }
    ],
    "effect": [
        "point 1",
        "point 2",
        "point 3",
        "point 4"
    ],
    "dengerous": 0,
    "contain-alcohol": false,
    "contain-pigs": false
}

*/
