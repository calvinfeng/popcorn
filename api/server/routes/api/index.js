const router = require('express').Router();

router.use('/recommendations', require('./recommendation'));
// router.use('/auth', require('./auth'));

module.exports = router;