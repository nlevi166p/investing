'use strict';

const MongoClient = require('mongodb').MongoClient;
const fs = require('fs');

let rawdata = fs.readFileSync('mock_data.json');
let instruments = JSON.parse(rawdata);

var url = "mongodb://localhost:27017";

MongoClient.connect(url, function (err, db) {
    if (err) throw err;
    let dbo = db.db("market");
    dbo.createCollection("instruments", function (err, col) {
        if (err) throw err;
        console.log("Collection created!");
        col.insertMany(instruments, function (err, res) {
            if (err) throw err;
            console.log(res);
            db.close();
        });
    });
});