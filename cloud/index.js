var exec = require('child_process').execSync
/**
 * Cloud Function.
 *
 * @param {object} event The Cloud Functions event.
 * @param {function} The callback function.
 */
exports.wright = function (event, callback) {
    var cmd = './wright';

    exec('./wright', function(error, stdout, stderr) {
	console.log('stdout: ', stdout);
	console.log('stderr: ', stderr);
	if (error !== null) {
            console.log('exec error: ', error);
	}
	callback();
    });
};
