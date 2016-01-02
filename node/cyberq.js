"use strict";

var Agent = require('superagent-promise')(require('superagent'), Promise);
var Util = require('util')
var xml2js = require('xml2js');
var _ = require('lodash');

class CyberQ {
	constructor(config) {
		this.host = config.cyberq.host;
	}

	current() {
		return Agent.get(this.getRequestPath('all.xml'))
				.then((response) => {
					return new Promise((resolve, reject) => {
						xml2js.parseString(response.text, (err, result) => {
							if (!err) {
								resolve(result);
							} else {
								reject(err);
							}
						});
					});
				})
				// Log for debugging request reponses.
				// .then((result) => {
				// 	console.log(Util.inspect(result, {depth: null, colors: true}));
				// 	return result;
				// })
				.then((result) => {
					return new CyberQStatus(result);
				});
	}

	getRequestPath(path) {
		return `http://${this.host}/${path}`;
	}
}
/*
{ nutcallstatus:
   { COOK:
      [ { COOK_NAME: [ 'Big Green Egg' ],
          COOK_TEMP: [ '3216' ],
          COOK_SET: [ '4000' ],
          COOK_STATUS: [ '0' ] } ],
     FOOD1:
      [ { FOOD1_NAME: [ 'Chicken Quarters' ],
          FOOD1_TEMP: [ '1482' ],
          FOOD1_SET: [ '1750' ],
          FOOD1_STATUS: [ '0' ] } ],
     FOOD2:
      [ { FOOD2_NAME: [ 'Food2' ],
          FOOD2_TEMP: [ 'OPEN' ],
          FOOD2_SET: [ '1000' ],
          FOOD2_STATUS: [ '4' ] } ],
     FOOD3:
      [ { FOOD3_NAME: [ 'Food3' ],
          FOOD3_TEMP: [ 'OPEN' ],
          FOOD3_SET: [ '1000' ],
          FOOD3_STATUS: [ '4' ] } ],
     OUTPUT_PERCENT: [ '100' ],
     TIMER_CURR: [ '00:00:00' ],
     TIMER_STATUS: [ '0' ],
     DEG_UNITS: [ '1' ],
     COOK_CYCTIME: [ '6' ],
     COOK_PROPBAND: [ '500' ],
     COOK_RAMP: [ '0' ] } }
*/
class CyberQStatus {
	constructor(statusJson) {
		this.pit = new CyberQProbStatus(statusJson.nutcallstatus.COOK);
		this.food1 = new CyberQProbStatus(statusJson.nutcallstatus.FOOD1);
		this.food2 = new CyberQProbStatus(statusJson.nutcallstatus.FOOD2);
		this.food3 = new CyberQProbStatus(statusJson.nutcallstatus.FOOD3);
		this.output = parseInt(statusJson.nutcallstatus.OUTPUT_PERCENT[0]);
		this.timestamp = Date.now();
	}

	pit() { return this.pit; }
	food1() { return this.food1; }
	food2() { return this.food2; }
	food3() { return this.food3; }
	output() { return this.output; }
}

class CyberQProbStatus {
	constructor(statusJson) {
		if (_.isArray(statusJson)) {
			statusJson = statusJson[0];
		}

		for (var key in statusJson) {
			if (/NAME/.test(key)) {
				this.name = statusJson[key][0];
			} else if (/TEMP/.test(key)) {
				this.temp = parseInt(statusJson[key][0]);
			} else if (/SET/.test(key)) {
				this.set = parseInt(statusJson[key][0]);
			} else if (/STATUS/.test(key)) {
				this.status = parseInt(statusJson[key][0]);
			}
		}
	}

	name() { return this.name; }
	temp() { return this.temp; }
	set() { return this.set; }
	status() { return this.status; }
}

module.exports = CyberQ;
