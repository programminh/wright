var exec = require('child_process').execSync
/**
 * Cloud Function.
 *
 * @param {object} event The Cloud Functions event.
 * @param {function} The callback function.
 */
exports.wright = function (event, callback) {
    var cmd = './wright';

    exec('./wright');

    callback();
};
