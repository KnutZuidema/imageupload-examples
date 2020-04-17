from flask import Flask, request
from werkzeug.datastructures import FileStorage

import config

app = Flask(__name__)
app.config['MAX_CONTENT_LENGTH'] = config.MAX_CONTENT_LENGTH


@app.route('/upload', methods=["POST"])
def upload():
    if config.FORM_IMAGE_KEY not in request.files:
        return {'message': 'missing file'}, 400
    file = request.files[config.FORM_IMAGE_KEY]
    if file.filename == '':
        return {'message': 'missing file'}, 400
    handle_file(file)
    return '', 204


def handle_file(file: FileStorage):
    # do something with the file
    pass


if __name__ == '__main__':
    app.run(config.HOST, config.PORT)
