from flask import Flask
import logging
import sys

app = Flask(__name__)

# Configure Logging
app.logger.addHandler(logging.StreamHandler(sys.stdout))
app.logger.setLevel(logging.DEBUG)

@app.route('/')
def hello_world():
    return 'Hello brozZz!!'
