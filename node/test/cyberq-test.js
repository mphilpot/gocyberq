"use strict";

var Http = require('http');
var Url = require('url');
var FS = require('fs');
var Path = require('path');
var CyberQ = require('../cyberq');
var Util = require('util');

var PORT = 10001;

describe('CyberQ', () => {
	var server;

	var cyberq;

	before((done) => {
		server = Http.createServer((request, response) => {
			var srvUrl = Url.parse('http://' + request.url);

			var knownRequests = {
				'/status.xml': 'cyberq_status.xml',
				'/all.xml': 'cyberq_all.xml',
				'/config.xml': 'cyberq_config.xml',
			}

			var doc = knownRequests[srvUrl.path];
			if (!doc) throw new Error('Unknown request!!!!');

			var file = Path.resolve(__dirname, '..', '..', 'docs', doc);
			var responseStr = FS.readFileSync(file);

			if (!responseStr) throw new Error('Could not read the file!!!!');

			response.end(responseStr);
		});
		server.listen(PORT, done);
	});

	after((done) => {
		server.close(done);
		server = null;
	});

	beforeEach(() => {
		cyberq = new CyberQ({cyberq: {host: `127.0.0.1:${PORT}`}})
	});

	it('should load a full configuration', () => {
		return cyberq.current().then((result) => {
			console.log(Util.inspect(result, {depth: null, colors: true}));
		})
	});
});
