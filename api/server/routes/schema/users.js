const Joi = require('joi');
const email = Joi.string().email({ minDomainAtoms: 2 }).required();
const movieId = Joi.number().required();
const rating = Joi.number().positive().precision(1).min(0).max(5);

const schema = {
  validateEmail: {
    params: {
      email: email
    }
  },
  validateRatings: {
    body: {
      movieId: movieId,
      rating: rating
    }
  }
};

exports.usersSchema = schema;