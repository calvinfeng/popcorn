const ev = require('express-validation');
const { createLogger, format, transports } = require('winston');
const { combine, timestamp, label, printf } = format;

// Configure a logger that we will use throughout the application.
const myFormat = printf(info => (
  `${info.timestamp} [${info.label}] ${info.level}: ${info.message}`
));

const log = createLogger({
  level: 'info',
  format: combine(label({ label: 'api'}), timestamp(), myFormat),
  transports: [
    new transports.Console({colorize: true})
  ]
});

function newLogMiddleware() {
  return (req, res, next) => {
    res.locals.log = log;
    next();
  }
}

function validationLogMiddleware(err, req, res, next) {
  // specific for validation errors
  if (err instanceof ev.ValidationError) return res.status(err.status).json(err);

  if (process.env.NODE_ENV !== 'production') {
    return res.status(500).send(err.stack);
  } else {
    return res.status(500);
  }
}

exports.newLogMiddleware = newLogMiddleware;
exports.log = log;
exports.validationLogMiddleware = validationLogMiddleware;