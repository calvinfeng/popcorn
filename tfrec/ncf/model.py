from keras.models import Model
from keras.layers import Embedding, Flatten, Input, Dot
from keras.constraints import non_neg
from keras.utils import plot_model

import numpy as np


latent_dim = 5
num_movies = 100
num_users = 100


if __name__ == '__main__':
    movie_input = Input(shape=[1], name='movie-input')
    movie_embedding = Embedding(num_movies + 1, latent_dim,
                                name='movie-embedding', embeddings_constraint=non_neg())(movie_input)
    movie_vec = Flatten(name='movie-flatten')(movie_embedding)

    user_input = Input(shape=[1], name='user-input')
    user_embedding = Embedding(num_users + 1, latent_dim,
                               name='user-embedding', embeddings_constraint=non_neg())(user_input)
    user_vec = Flatten(name='user-flatten')(user_embedding)

    dot = Dot(axes=1, name='dot-product')([movie_vec, user_vec])
    model = Model(inputs=[user_input, movie_input], outputs=dot)

    model.compile('adam', 'mean_squared_error')
    plot_model(model, to_file='model.png', show_shapes=True, show_layer_names=True)

    rand_users = np.random.randint(1, 100, size=(10, 1))
    rand_movies = np.random.randint(1, 100, size=(10, 1))

    print(model.predict([rand_users, rand_movies]))
