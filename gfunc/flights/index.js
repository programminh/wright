var exec = require('child_process').execSync

exports.flights = function (event, callback) {
    exec('./flights');

    callback();
};
