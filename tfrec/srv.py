from flask import (
    Flask, jsonify, render_template, request
)

from concurrent import futures

from keras.applications.resnet50 import (
    ResNet50, preprocess_input, decode_predictions
)

from calculator_pb2_grpc import (
    CaculatorServicer, add_CaculatorServicer_to_server
)

from calculator_pb2 import Output

import numpy as np
import cv2
import grpc


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

    resp = jsonify({'classes': classes})
    resp.status_code = 200
    return resp


class Calculator(CaculatorServicer):
    def Add(self, input, context):
        output = Output()
        output.value = input.left + input.right
        return output


if __name__ == '__main__':
    grpc_srv = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    add_CaculatorServicer_to_server(Calculator(), grpc_srv)
    grpc_srv.add_insecure_port('[::]:8081')

    print('starting gRPC server on port 8081 and Flask server on port 8080')
    grpc_srv.start()
    app.run(debug=False, threaded=False, port=8080)
