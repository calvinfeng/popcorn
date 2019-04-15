# Model Server

I want to serve Keras models with Python instead of Go because this will speed up development
quite a bit. At least I don't have to learn the TF API in Go, there are too few tutorials on
this.

## Create `venv`

This project uses Python 3. 

    python -m venv venv
    source venv/bin/activate

And then install requirements

    pip install -r requirements.txt

Run the simple Flask application server

    python3 app.py