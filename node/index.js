var fs = require('fs');
var nconf = require('nconf');

//
// Setup nconf to use (in-order):
//   1. Command-line arguments
//   2. Environment variables
//   3. A file located at 'path/to/config.json'
//
nconf.argv()
	.env()
	.file({ file: 'config.json' });

// Configure cyberq specific values
nconf.set('cyberq:host', '127.0.0.1');

// Configure server specific values
nconf.set('server:port', '3333');
nconf.set('server:poll', '2000');

// Print settings
console.log('cyberq: ' + nconf.get('cyberq'));
console.log('server: ' + nconf.get('server'));

//
// Save the configuration object to disk
//
nconf.save(function (err) {
	fs.readFile('config.json', (err, data) => {
		console.dir(JSON.parse(data.toString()))
	});
});
