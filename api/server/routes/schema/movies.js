const Joi = require('joi');
const imdbId = Joi.string().required();

const schema = {
  getMovie: {
    params: {
      imdbId: imdbId
    }
  }
};

exports.schema = schema;