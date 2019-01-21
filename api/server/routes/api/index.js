const router = require('express').Router();

router.use('/recommend', require('./recommend'));
router.use('/movies', require('./movies'));

module.exports = router;