from flask import Flask, jsonify, render_template
from flask import request
import numpy as np
import cv2


app = Flask(__name__)


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
    resp = jsonify(np_data.reshape(-1, *np_data.shape).shape)
    resp.status_code = 200
    return resp


if __name__ == '__main__':
    app.run(debug=True, port=8080)