const Joi = require('joi');
const imdbId = Joi.string().regex(/^tt[0-9]{7}/).required();
const movieId = Joi.number().required();
const page = Joi.number().integer().positive().required();

const schema = {
  validateImdbId: {
    params: {
      imdbId: imdbId
    }
  },
  validatePage: {
    query: {
      page: page
    }
  }
};

exports.moviesSchema = schema;