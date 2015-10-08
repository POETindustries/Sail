(function() {
    var loc = window.location.pathname;
    var entries = document.getElementById("main_menu").children;
    for (var i = 0; i < entries.length; i++) {
        if (entries[i].pathname == loc) {
            entries[i].id = "active";
            break;
        }
    }
})();
