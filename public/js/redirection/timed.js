let durationInSeconds;
let redirectionURL;

function setDuration(duration) {
    durationInSeconds = duration;
}

function setRedirectionURL(url) {
    redirectionURL = url;
}

function updateProgressBar() {
    const progressBar = document.getElementById("progressBar");
    const max = parseInt(progressBar.getAttribute("max"), 10);
    const increment = (max / (durationInSeconds * 1000 / 60));

    let currentValue = 0;

    function update() {
        if (currentValue >= max) {
            window.location.replace(redirectionURL);
            return;
        }

        currentValue += increment;
        progressBar.value = currentValue;

        requestAnimationFrame(update);
    }

    update();
}

function startProgress(duration, url) {
    setDuration(duration);
    setRedirectionURL(url);
    updateProgressBar();
}