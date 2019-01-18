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

exports.newLogMiddleware = newLogMiddleware;
exports.log = log;