package lowrank

// trainTestRatio is a split ratio between number of ratings for training and number of ratings for
// testing.
var trainTestRatio = 0.10

// SetTrainTestRatio sets trainTestRatio.
func SetTrainTestRatio(n float64) {
	trainTestRatio = n
}

// minRatingsPerUser is a threshold to filter out user that did not rate enough movies.
var minRatingsPerUser = 300

// SetMinRatingsPerUser sets minRatingsPerUser
func SetMinRatingsPerUser(n int) {
	minRatingsPerUser = n
}
