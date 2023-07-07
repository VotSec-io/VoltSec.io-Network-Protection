
document.addEventListener("DOMContentLoaded", function() {
    var form = document.querySelector("form");

    form.addEventListener("submit", function(event) {
        event.preventDefault();

        var confirmMessage = "Are you sure you want to save the settings?";
        if (confirm(confirmMessage)) {
            // Perform form submission or other desired actions
            form.submit();
        }
    });
});
