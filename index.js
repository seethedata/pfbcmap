var express = require('express');
var app = express();
var http = require('http');
var httpServer = http.Server(app);
var Promise=require('bluebird');

var bootcampLocations=[
					{city: "Johannesburg, South Africa" },
					{city: "Los Angeles, CA" },
					{city: "Irvine, CA" },
					{city: "Salt Lake City, UT" },
					{city: "Tampa, FL" },
					{city: "San Diego, CA" },
					{city: "Atlanta, GA" },
					{city: "Phoenix, AZ" },
					{city: "Philadelphia, PA" },
					{city: "Houston, TX" },
					{city: "Dallas, TX" },
					{city: "Frankfurt, Germany" },
					{city: "Kansas City, MO" },
					{city: "Paris, France" },
					{city: "Tampa, FL" },
					{city: "Tulsa, OK" },
					{city: "Chicago, IL" },
					{city: "London, England" },
					{city: "Boston, MA" },
					{city: "Seattle, WA" },
					{city: "Charlotte, NC" },
					{city: "Asheville, NC" },
					{city: "Rockville, MD" },
					{city: "Denver, CO" },
					{city: "San Francisco, CA" },
					{city: "Oakland, CA" },
					{city: "Des Moines, IA" }	
				];

var geocodePromises=[];

var googleMapsClient = Promise.promisifyAll(require('@google/maps').createClient({
  key: 'AIzaSyDKmwcRIx2Wg2QYJfhGZxIDhEuO5TXi_h0'
}));



for (var i = 0; i < bootcampLocations.length; i++) {
	geocodePromises.push(googleMapsClient.geocodeAsync({address : bootcampLocations[i].city}));
}

Promise.all(geocodePromises).then(function(result) {
	result.forEach(function geocodeResponse(currentValue, index, array) {
		bootcampLocations[index].position=result[index].json.results[0].geometry.location;
	});


	app.set('views',__dirname + "/views");
	app.set('view engine', 'pug');
	app.use(express.static(__dirname));

	app.get('/', function(req, res){
		res.render('gm.pug', {"locations" : JSON.stringify(bootcampLocations)});
	});
	app.listen(process.env.port || 8080);	
});










