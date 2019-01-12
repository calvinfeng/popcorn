const router = require('express').Router();

router.use('/recommendations', require('./recommendation'));
router.use('/authentication', require('./authentication'));

module.exports = router;