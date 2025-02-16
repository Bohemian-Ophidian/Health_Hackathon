from flask import Flask, request, redirect, url_for
import os
from map_service import find_hospitals_osm, generate_map

app = Flask(__name__)

@app.route('/') 
def index():
    return '''
        <form action="/submit" method="post">
            Enter Postal Code: <input type="text" name="postal_code">
            <input type="submit">
        </form>
    '''

@app.route('/submit', methods=['POST'])
def submit():
    postal_code = request.form['postal_code']
    hospitals = find_hospitals_osm(postal_code)
    if isinstance(hospitals, list):
        generate_map(hospitals)
    return redirect(url_for('map'))

@app.route('/map')
def map():
    return redirect("http://127.0.0.1:5500/map_service/map.html")

if __name__ == '__main__':
    app.run(debug=True)
