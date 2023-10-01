package provider

var id int64

func SetPetID(petID int64) {
	id = petID
}

func GetPetID() int64 {
	return id
}
