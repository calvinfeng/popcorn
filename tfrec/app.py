from flask import Flask, jsonify, render_template
from flask import request
from keras.applications.resnet50 import ResNet50, preprocess_input, decode_predictions
import numpy as np
import cv2


app = Flask(__name__)
model = ResNet50(weights='imagenet')


@app.route('/', methods=['GET'])
def index():
    return render_template('index.html')


@app.route('/api/classify', methods=['POST'])
def classify():
    bytestr = None
    if len(request.files) != 0 and request.files['image'] is not None:
        bytestr = request.files['image'].read()
    elif len(request.data) != 0:
        bytestr = request.data
    else:
        # No data is found, return an error response
        resp = jsonify({'error': 'no image found'})
        resp.status_code = 400
        return resp

    np_data = cv2.imdecode(np.fromstring(bytestr, np.uint8), cv2.IMREAD_COLOR)
    np_data = cv2.resize(np_data, (224, 224))
    x = np_data.reshape(-1, *np_data.shape)
    x = preprocess_input(x)
    preds = model.predict(x)
    preds = decode_predictions(preds, top=3)[0]

    classes = []
    for pred in preds:
        classes.append(pred[1])

    resp = jsonify({ 'classes': classes })
    resp.status_code = 200
    return resp


if __name__ == '__main__':
    app.run(debug=False, threaded=False, port=8080)