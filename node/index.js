var CyberQ = require('./cyberq');
var Express = require('express');
var FS = require('fs');
var Util = require('util');
var nconf = require('nconf');

////////////////////////////////////////////////////////////////////////////////
/// Configuration
////////////////////////////////////////////////////////////////////////////////

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
nconf.defaults({
	cyberq: {
		host: '127.0.0.1'
	},
	server: {
		port: '3333',
		poll: 2000
	}
});

// Print settings
console.log('cyberq: ' + JSON.stringify(nconf.get('cyberq')));
console.log('server: ' + JSON.stringify(nconf.get('server')));

//
// TODO: Save the default configuration object to disk.
//
// nconf.save(function (err) {
// 	FS.readFile('config.json', (err, data) => {
// 		console.dir(JSON.parse(data.toString()))
// 	});
// });


////////////////////////////////////////////////////////////////////////////////
/// Server
////////////////////////////////////////////////////////////////////////////////

var app = Express();

var cyberqInstance = new CyberQ(nconf.get());
app.get('/', (req, res) => {
	cyberqInstance.current()
			.then((result) => {
				res.send(`hello world<br><pre>${JSON.stringify(result)}</pre>`);
			})
			.catch((err) => {
				res.send('error loading resource');
				console.log(err);
				return Promise.reject(err);
			})
});

var server = app.listen(nconf.get('server:port'), () => {
  var host = server.address().address;
  var port = server.address().port;

  console.log(`Example app listening at http://${host}:${port}`);
});
