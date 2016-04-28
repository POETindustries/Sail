/*
 * init setup for user management.
 */
(function(){
	var toggleBtn = document.getElementById("user-group-toggle")
	var userTable = document.getElementById("sail-backend-users-table")
	var groupTable = document.getElementById("sail-backend-groups-table")
	toggleBtn.onclick = function(){
		if (toggleBtn.checked) {
			groupTable.style.display = "table";
			userTable.style.display = "none";
		} else {
			groupTable.style.display = "none";
			userTable.style.display = "table";
		}
	};
	initSelectAll(userTable);
	initSelectAll(groupTable);
})();

/*
 * helper function for initializing the 'select-all' functionality
 * for the user and group tables.
 */
function initSelectAll(table) {
	var boxes = table.getElementsByTagName("input")
	boxes[0].onclick = function(){
		for (var j = 1; j < boxes.length; j++) {
			boxes[j].checked = boxes[0].checked;
		}
	};
}

/*
 * hides all table entries that do not match the keyword.
 */
function filter(keyword) {
	console.log(keyword);
}
