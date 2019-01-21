import pandas as pd 
import matplotlib.pyplot as plt
import re
from sklearn.decomposition import PCA
from sklearn.preprocessing import StandardScaler


def load_movie_titles():
  movie_title_by_id = dict()
  for index, row in pd.read_csv('../datasets/100k/movies.csv').iterrows():
    movie_title_by_id[row[0]] = row[1]
  
  return movie_title_by_id


def plot_all(df):
  movie_title_by_id = load_movie_titles()
  ids = final_df['movieId'].values
  x = final_df['pc1'].values
  y = final_df['pc2'].values
  
  fig, ax = plt.subplots()
  ax.scatter(x, y)
  for i, id in enumerate(ids):
    txt = unicode(movie_title_by_id[id], "utf-8")
    ax.annotate(txt, (x[i], y[i]))

  plt.show()


def plot_year_range(df, start_year, end_year):
  movie_title_by_id = load_movie_titles()
  
  titles, x, y = [], [], []
  p = re.compile("\d{4}")
  for index, row in df.iterrows():
    year_match =  p.findall(movie_title_by_id[row['movieId']]) 
    if len(year_match) == 0:
      continue

    year = int(year_match[0])
    if start_year <= year and year <= end_year:
      titles.append(unicode(movie_title_by_id[row['movieId']], "utf-8"))
      x.append(row['pc1'])
      y.append(row['pc2'])

  fig, ax = plt.subplots()
  ax.scatter(x, y)
  for i, title in enumerate(titles):
    ax.annotate(title, (x[i], y[i]))

  plt.show()


def main():

  df = pd.read_csv('../datasets/100k/features.csv')
  print 'Initial data frame'
  print df

  features = df.drop(['movieId'], axis=1).values
  x = StandardScaler().fit_transform(features)
  pca = PCA(n_components=2)

  principal_df = pd.DataFrame(data=pca.fit_transform(x), columns = ['pc1', 'pc2'])
  final_df = pd.concat([df[['movieId']], principal_df], axis = 1)

  print 'Final data frame after PCA'
  print final_df

  plot_year_range(final_df, 1990, 2000)


if __name__ == '__main__':
  main()