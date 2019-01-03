const Http = require('http');
const Url = require('url');
const FS = require('fs');
const Path = require('path')
const Util = require('util');

const PORT = 10001;
const server = Http.createServer((request, response) => {
	try {
		const srvUrl = Url.parse('http://' + request.url);

		const knownRequests = {
			'/status.xml': 'cyberq_status.xml',
			'/all.xml': 'cyberq_all.xml',
			'/config.xml': 'cyberq_config.xml',
		}

		const doc = knownRequests[srvUrl.path];
		if (!doc) throw new Error('Unknown request!!!!');

		const file = Path.resolve(__dirname, '..', 'docs', doc);
		const responseStr = FS.readFileSync(file);

		if (!responseStr) throw new Error('Could not read the file!!!!');

		response.end(responseStr);
	} catch (e) {
		response.statusCode = 500;
		response.end(e.toString());
	}
});

console.log(`Starting fake cyberq on port ${PORT}`);
server.listen(PORT);
