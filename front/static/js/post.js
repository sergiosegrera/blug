"use strict";

document.addEventListener("DOMContentLoaded", setup, false);

function setup() {
    fetch("/api/post/" + 0)
        .then(response => response.json())
        .then(post => {
            document.title = post.title;
            document.querySelector("body").innerHTML = post.html;
        })
}
