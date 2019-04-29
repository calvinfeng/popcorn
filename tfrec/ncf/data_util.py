import pandas as pd

from sklearn.model_selection import train_test_split


def load_train_test_dataframe(file_location: str, test_size=0.2):
    df = pd.read_csv(file_location, header=0, names=['user_id', 'movie_id', 'rating', 'timestamp'])

    # Map movie ID to [1, num_movies]
    movie_id_to_new_id = dict()
    movie_id = 1
    for index, row in df.iterrows():
        if movie_id_to_new_id.get(row['movie_id']) is None:
            movie_id_to_new_id[row['movie_id']] = movie_id
            df.at[index, 'movie_id'] = movie_id
            movie_id += 1
        else:
            df.at[index, 'movie_id'] = movie_id_to_new_id.get(row['movie_id'])

    train, test = train_test_split(df, test_size=test_size)
    return (train, test)


if __name__ == '__main__':
    train, test = load_train_test_dataframe('../datasets/100k/ratings.csv')
    print(train.head())
    print(test.head())
