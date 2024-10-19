package handlers

const (
	invalidLatLon                   = "invalid latitude or longitude"
	invalidLimit                    = "invalid limit"
	invalidCity                     = "invalid city"
	weatherFetchingError            = "something went wrong fetching the weather"
	geocodingFetchingError          = "something went wrong fetching the geocode"
	reverseGeocodingFetchingError   = "something went wrong fetching the city"
	successFetchingWeather          = "successfully retrieved the weather"
	successFetchingGeocode          = "successfully retrieved the geocode"
	successFetchingReverseGeocoding = "successfully retrieved the city"
	successCreatingUser             = "successfully created the user"
	successUpdatingUser             = "successfully updated the user"
	successDeletingUser             = "successfully deleted the user"
	userCreationError               = "something went wrong creating the user"
	userUpdationError               = "something went wrong updating the user"
	userDeletionError               = "something went wrong deleting the user"
	userNotFound                    = "user not found"
)
