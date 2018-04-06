/*
 * Configure the locations div to be an accordion on page load
 */
$(function () {
    $(".locList").accordion({
        collapsible: true,
    });
})

/*
 * Add a new empty location entry in the locations list.
 */
$(".locAdd").click(function () {
    /* prepend the HTML code in the locations list */
    $(".locList").prepend(`
        <h3></h3>
        <div class="location">
            <p><input type="text" class="date" /></p>
            <p>Lat: <input type="text" class="latitude" />
            Long: <input type="text" class="longitude" /></p>
            <p><textarea class="refs"></textarea></p>
            <p><button class="rmButton" onclick="rmLocation(event);">Remove</button></p>
        </div>`
    );
    /* refresh the date and accordion */
    $(".date").datepicker()
    $(".locList").accordion("refresh");
})

/*
 * Remove a location from the locations list
 */
function rmLocation(event) {
    /* remove the current location */
    var itemToRemove = $(event.target).parent().parent();
    var titleToRemove = itemToRemove.prev();
    itemToRemove.remove();
    titleToRemove.remove();
    /* refresh the accordion */
    $(".locList").accordion("refresh");
}

/* Returns the date of the given input location item */
function getDate(location) {
    return $(location).find(".date").datepicker("getDate");
}

/*
 * Returns the JSON representation of the Person described.
 */
function getPerson() {
    /* definition of the person's characteristics */
    var person = {
        name: $(".personName").val(),
        desc: $(".personDesc").val(),
        locations: []
    };

    /* now get locations */
    $(".locList > div").each(function () {
        person.locations.push({
            time: getDate(this),
            latitude: $(this).find(".latitude").val(),
            longitude: $(this).find(".longitude").val(),
            refs: $(this).find(".refs").val()
        });
    });

    /* sort locations by date */
    person.locations.sort(function (a, b) {
        return new Date(a.time) - new Date(b.time);
    });

    return person;
}

$(".personSave").click(function (event) {
    var person = getPerson();
    console.log(person);

    /* update the accordion headers with the dates */
    $(".locList > h3").each(function () {
        var location = $(this).next();
        var date = getDate(location);
        $(this).text(date);
    })
});