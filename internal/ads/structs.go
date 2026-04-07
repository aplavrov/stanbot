package ads

import "time"

type Response struct {
	Result []Property `json:"result"`
}

type Property struct {
	UniqueID          string   `json:"uniqueID"`
	PropID            int      `json:"propId"`
	StatusID          int      `json:"statusId"`
	CityID            int      `json:"cityId"`
	Location          string   `json:"location"`
	Street            string   `json:"street"`
	Floor             string   `json:"floor"`
	Size              int      `json:"size"`
	Structure         string   `json:"structure"`
	Municipality      string   `json:"municipality"`
	Polygons          []string `json:"polygons"`
	PtID              int      `json:"ptId"`
	Price             float64  `json:"price"`
	CoverPhoto        string   `json:"coverPhoto"`
	RentOrSale        string   `json:"rentOrSale"`
	CaseID            int      `json:"caseId"`
	CaseType          string   `json:"caseType"`
	UnderConstruction bool     `json:"underConstruction"`
	Filed             int      `json:"filed"`
	Furnished         int      `json:"furnished"`
	Ceiling           int      `json:"ceiling"`

	FurnishingArray    []string `json:"furnishingArray"`
	BldgOptsArray      []string `json:"bldgOptsArray"`
	HeatingArray       []int    `json:"heatingArray"`
	ParkingArray       []int    `json:"parkingArray"`
	YearOfConstruction int      `json:"yearOfConstruction"`
	JoineryArray       []int    `json:"joineryArray"`
	PetsArray          []int    `json:"petsArray"`
	OtherArray         []string `json:"otherArray"`

	AvailableFrom  time.Time `json:"availableFrom"`
	FirstPublished time.Time `json:"firstPublished"`

	PricePerSize float64 `json:"pricePerSize"`

	ShowPdv  *bool    `json:"showPdv"`  // null → pointer
	OldPrice *float64 `json:"oldPrice"` // null → pointer

	IsSalonac          bool `json:"isSalonac"`
	IsNotLastFloor     bool `json:"isNotLastFloor"`
	IsNoElevatorButLow bool `json:"isNoElevatorButLow"`

	BedroomsArray []string `json:"bedroomsArray"`
	BathroomArray []string `json:"bathroomArray"`
	//RenovationArray     []string `json:"renovationArray"`
	MinLeaseArray       []int    `json:"minLeaseArray"`
	NewDevelopment      bool     `json:"newDevelopment"`
	DistanceCenterArray []string `json:"distanceCenterArray"`

	IsFeatured       bool `json:"isFeatured"`
	IsLux            bool `json:"isLux"`
	NewImagePipeline bool `json:"newImagePipeline"`
}
