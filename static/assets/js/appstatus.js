/* appstatus.js
 * Copyright 2016 UpGuard Inc.
 * Zakk Acreman <zakk.acreman@upguard.com>
 * All rights reserved
 */

function populateDockerStats(data) {
		var table = $('#docker-table');
		var create = document.createElement

		function newTd (value) {
				return $(create('td')).append(value);
		}
		
		data.forEach(function (v, i, data) {
				var id = v["Id"].slice(0,8),
						name = v["Names"][0],
						image = v["Image"],
						status = v["Status"].split(" ")[0],
						time = v["Status"].split(" ").slice(1).join(" ");
				
				table.append($(create('tr'))
										 .append(newTd(id))
										 .append(newTd(name))
										 .append(newTd(image))
										 .append(newTd(status))
										 .append(newTd(time)))
		});
		table.show();
}

function populateFleetUnits(data) {
}
